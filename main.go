// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

const dbfile string = "mytime.db"

var db *sql.DB

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	// need to inject the form based on the required fields from the database
	tmpl := template.Must(template.ParseFiles("timesheet.html"))
	tmpl.Execute(w, nil)
}

func saveEventDatatoDB(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /saveEventDatatoDB request\n")
	if r.Method == "POST" {
		fmt.Println("Receive ajax post data string...")
		err := r.ParseForm()
		if err != nil {
			log.Println("Form Parse Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse the form data!"))
			return
		}
		event := eventData{}
		id := r.Form.Get("eventid")
		if id != "" {
			event.id, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				log.Println("Could not parse ID:", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Could not parse ID"))
				return
			}
		}
		event.start = r.Form.Get("start")
		event.end = r.Form.Get("end")
		event.description = r.Form.Get("description")
		event.customer = r.Form.Get("customer")
		event.primaryLogType, err = strconv.Atoi(r.Form.Get("primaryLogType"))
		event.secondaryLogType, err = strconv.Atoi(r.Form.Get("secondaryLogType"))

		eventid, err := putEventDatatoDB(db, event)
		if err != nil {
			log.Println("Could not write to DB:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not write to database"))
			return
		}
		// for key, value := range r.Form {
		// 	event.
		// 		fmt.Printf("%s - %s\n", key, value)
		// }
		//w.Header().Add("Content-Type", "application/html")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		//w.Write([]byte(tpl.String()))
		w.Write([]byte(strconv.FormatInt(eventid, 10)))
	}
}

func updateEventDatatoDB(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /updateEventDatatoDB request\n")
	if r.Method == "POST" {
		fmt.Println("Receive ajax post data string...")
		err := r.ParseForm()
		if err != nil {
			log.Println("Form Parse Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse the form data!"))
			return
		}
		event := eventData{}
		id := r.Form.Get("id")
		if id != "" {
			event.id, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				log.Println("Could not parse ID:", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Could not parse ID"))
				return
			}
		}
		event.start = r.Form.Get("start")
		event.end = r.Form.Get("end")

		err = putEventTimeChangetoDB(db, event)
		if err != nil {
			log.Println("Could not write update to DB:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not write to database"))
			return
		}
		w.Write([]byte(""))
	}
}

func deleteEventDatafromDB(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /deleteEventDatafromDB request\n")
	if r.Method == "POST" {
		fmt.Println("Receive ajax post data string...")
		err := r.ParseForm()
		if err != nil {
			log.Println("Form Parse Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse the form data!"))
			return
		}
		event := eventData{}
		id := r.Form.Get("id")
		if id != "" {
			event.id, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				log.Println("Could not parse ID:", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - Could not parse ID"))
				return
			}
		}

		err = deleteEventData(db, event)
		if err != nil {
			log.Println("Could not write update to DB:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not write to database"))
			return
		}
		w.Write([]byte(""))
	}
}

func getEventDatafromDB(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /getEventDatafromDB request\n")
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Form Parse Error:", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not parse the form data!"))
			return
		}
		start := r.Form.Get("start")
		end := r.Form.Get("end")
		limit := r.Form.Get("limit")
		maximumreturnedresults := 1000
		if limit != "" {
			maximumreturnedresults, _ = strconv.Atoi(limit)
		}

		jstring, err := getEventsJSON(db, start, end, maximumreturnedresults)
		if err != nil {
			fmt.Println("getEventDatafromDB: Returned JSON string: ", jstring, err)
		}

		w.Write([]byte(jstring))
	}
}

func getCategoriesfromDB(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /getCategoriesfromDB request\n")
	if r.Method == "GET" {
		jstring, err := getCategoriesJSON(db)
		if err != nil {
			fmt.Println("Returned JSON string: ", jstring)
		}
		w.Write([]byte(jstring))
	}
}

func getCustomersfromDB(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /getCustomersfromDB request\n")
	if r.Method == "GET" {
		jstring, err := getCustomersJSON(db)
		if err != nil {
			fmt.Println("Returned JSON string: ", jstring)
		}
		w.Write([]byte(jstring))
	}
}

func loadFile(w http.ResponseWriter, r *http.Request) {
	p := "." + r.URL.Path
	http.ServeFile(w, r, p)
}

func main() {

	db = Opendatabase(dbfile)
	defer db.Close()

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/saveEventDatatoDB", saveEventDatatoDB)
	http.HandleFunc("/updateEventDatatoDB", updateEventDatatoDB)
	http.HandleFunc("/getEventDatafromDB", getEventDatafromDB)
	http.HandleFunc("/getCategoriesfromDB", getCategoriesfromDB)
	http.HandleFunc("/helperfunctions.js", loadFile)
	http.HandleFunc("/getCustomersfromDB", getCustomersfromDB)
	http.HandleFunc("/deleteEventDatafromDB", deleteEventDatafromDB)

	fmt.Println("Server starting on localhost:3333")
	//openURL("http://localhost:3333")
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
