package main

import (
	"database/sql"
	"encoding/json"

	// _ "github.com/mattn/go-sqlite3"

	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

// --------------------------------------------------------------------
func Opendatabase(dbfile string) *sql.DB {
	db, err := sql.Open("sqlite", dbfile)
	if err != nil {
		log.Fatal(err)
	}

	// create the tables
	_, err = db.Exec(createCustomerTable)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(createPrimaryCat)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(createSecondaryCat)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(createEvent)
	if err != nil {
		log.Fatal(err)
	}

	// add or update the categories
	//row := c.db.QueryRow("SELECT id, time, description FROM activities WHERE id=?", id)
	var id, id2 int64
	for _, items := range newDefaultCategories {
		for j, category := range items {
			if j == 0 {
				namecolour := strings.Split(category, "#")
				row := db.QueryRow("SELECT id FROM primary_category WHERE name=?", namecolour[0])
				if row.Scan(&id) == sql.ErrNoRows {
					// need to create
					res, err := db.Exec("INSERT INTO primary_category(name,colour) VALUES(?,?)", namecolour[0], namecolour[1])
					if err != nil {
						log.Fatal("Insert fail: ", err)
					}
					id, _ = res.LastInsertId()
				}
			} else {
				// need to add all the secondary categories
				row := db.QueryRow("SELECT id FROM secondary_category WHERE primary_id=? AND name=?", id, category)
				if row.Scan(&id2) == sql.ErrNoRows {
					// need to create
					_, err := db.Exec("INSERT INTO secondary_category(name,primary_id) VALUES(?,?)", category, id)
					if err != nil {
						log.Fatal("Insert fail: ", err)
					}
				}
			}
		}
	}

	return db
}

// --------------------------------------------------------------------
func getCategoriesJSON(db *sql.DB) (json_data string, err error) {

	type secondary_category struct {
		ID   int64
		Name string
	}

	type primary_category struct {
		ID        int64
		Name      string
		Colour    string
		Secondary []secondary_category
	}

	rows, err := db.Query("SELECT id, name, colour from primary_category")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var categories []primary_category

	// loop thought the rows
	for rows.Next() {
		var pc primary_category
		err := rows.Scan(&pc.ID, &pc.Name, &pc.Colour)
		if err != nil {
			log.Println("Scan of primary category failed, ", err)
		} else {
			rows2, err := db.Query("SELECT id, name from secondary_category WHERE primary_id=? ORDER BY name", pc.ID)
			if err != nil {
				log.Fatal("Could not get secondary_category data")
			}
			defer rows2.Close()
			for rows2.Next() {
				var sc secondary_category
				err := rows2.Scan(&sc.ID, &sc.Name)
				if err != nil {
					log.Println("Scan of secondary category failed, ", err)
				} else {
					pc.Secondary = append(pc.Secondary, sc)
				}
			}
			categories = append(categories, pc)
		}
	}

	b, err := json.Marshal(categories)
	if err != nil {
		fmt.Println("error:", err)
	}

	return string(b), nil
}

// --------------------------------------------------------------------
func getEventsJSON(db *sql.DB, start string, end string) (json_data string, err error) {

	type event struct {
		ID          int64
		Description string
		Start       string
		End         string
		Customer    string
		Colour      string
	}

	rows, err := db.Query(`SELECT
			event.id,
			event.description,
			event.start,
			event.end,
			customer.customername,
			primary_category.colour
		FROM event
		JOIN customer on event.customer_id = customer.id
		JOIN primary_category on event.primary_id=primary_category.id
		WHERE start>=? AND end <=?`,
		start, end)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var events []event
	// loop thought the rows
	for rows.Next() {
		var e event
		err := rows.Scan(&e.ID, &e.Description, &e.Start, &e.End, &e.Customer, &e.Colour)
		if err != nil {
			log.Println("Scan of event, ", err)
		} else {
			events = append(events, e)
		}
	}

	b, err := json.Marshal(events)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b), nil
}

// --------------------------------------------------------------------
func putCustomertoDB(db *sql.DB, customer string) (databaseid int64, err error) {
	row := db.QueryRow("SELECT id from customer WHERE customername = ?", customer)
	if row.Scan(&databaseid) == sql.ErrNoRows {
		// need to create
		log.Println("Creating customer: ", customer)
		res, err := db.Exec("INSERT INTO customer (customername) VALUES (?)", customer)
		if err != nil {
			log.Println("putCustomertoDB:Insert customer fail: ", err)
			return 0, err
		}
		databaseid, _ = res.LastInsertId()
		log.Println("Customer ID: ", databaseid)
	}
	return databaseid, nil
}

// --------------------------------------------------------------------
func putEventDatatoDB(db *sql.DB, event eventData) (databaseid int64, err error) {
	// either add or update event data in the database
	if event.id == 0 {
		// need to add to the database
		customerid, err := putCustomertoDB(db, event.customer)
		if err != nil {
			log.Println("putEventDatatoDB: Could not get customer ID:", err)
			return 0, err
		}
		res, err := db.Exec("INSERT INTO event (description, start, end, customer_id, primary_id, secondary_id) VALUES (?, ?, ?,?,?,?)",
			event.description, event.start, event.end, customerid, event.primaryLogType, event.secondaryLogType)
		if err != nil {
			log.Println("putEventDatatoDB: Could not insert event:", err)
			return 0, err
		}
		databaseid, _ = res.LastInsertId()
	} else {
		// need to update the database
		customerid, err := putCustomertoDB(db, event.customer)
		if err != nil {
			log.Println("putEventDatatoDB: Could not get customer ID:", err)
			return 0, err
		}
		_, err = db.Exec("UPDATE event SET description=?, start=?, end=?, customer_id=?, primary_id=?, secondary_id=? WHERE id=?",
			event.description, event.start, event.end, customerid, event.primaryLogType, event.secondaryLogType, event.id)
		if err != nil {
			log.Println("putEventDatatoDB: Could not insert event:", err)
			return 0, err
		}
		databaseid = event.id
	}
	return databaseid, nil
}

// --------------------------------------------------------------------
func putEventTimeChangetoDB(db *sql.DB, event eventData) (err error) {
	// need to update the times
	log.Println("ID:? Start:? End:?", event.id, event.start, event.end)
	_, err = db.Exec("UPDATE event SET start=?, end=? WHERE id=?",
		event.start, event.end, event.id)
	if err != nil {
		log.Println("putEventTimeChangetoDB: Could not update event:", err)
		return err
	}
	return nil
}
