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
					log.Println("Scan of secondardy category failed, ", err)
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
