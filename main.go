// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

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
		data := r.FormValue("post_data")
		fmt.Println("Receive ajax post data string ", data)
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("%s - %s\n", key, value)
		}
		//w.Header().Add("Content-Type", "application/html")
		//w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		//w.Write([]byte(tpl.String()))
	}
}

func getEventDatafromDB(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /getEventDatafromDB request\n")
	if r.Method == "GET" {
		start := r.FormValue("start")
		end := r.FormValue("end")
		parseTime, err := time.Parse(start, "Wed Jan 17 2024 08:15:00 GMT 1300 (New Zealand Daylight Time)")
		if err == nil {
			fmt.Println("Start ", parseTime)
		} else {
			fmt.Println("Parse Error")
		}
		fmt.Println("Receive ajax get data string ", start, end)
	}
}

func getCategoriesfromDB(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /getCategoriesfromDB request\n")
	if r.Method == "GET" {
		jstring, err := getCategoriesJSON(db)
		if err != nil {
			fmt.Println("Returned JSON string: ", jstring)
		}
		w.Write([]byte(jstring))
	}
}

func main() {

	db = Opendatabase(dbfile)
	defer db.Close()

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/saveEventDatatoDB", saveEventDatatoDB)
	http.HandleFunc("/getEventDatafromDB", getEventDatafromDB)
	http.HandleFunc("/getCategoriesfromDB", getCategoriesfromDB)

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
