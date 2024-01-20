package main

import (
	"database/sql"
	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
  "fmt"
  "log"
)

const file string = "activities.db"

const create string = `
  CREATE TABLE IF NOT EXISTS activities (
  id INTEGER NOT NULL PRIMARY KEY,
  time DATETIME NOT NULL,
  description TEXT
  );`


func main() {
	db, err := sql.Open("sqlite", file)
  if err != nil {
    log.Fatal(err)
  }
  
  defer db.Close()

  var version string
  err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(version)

  _, err = db.Exec(create)
  if err != nil {
    log.Fatal(err)
  }

}