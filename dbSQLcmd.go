package main

const createCustomerTable string = `
	CREATE TABLE IF NOT EXISTS customer (
	id INTEGER NOT NULL PRIMARY KEY,
	customername TEXT
	);`

const createPrimaryCat string = `
	CREATE TABLE IF NOT EXISTS primary_category (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	colour TEXT
	);`

const createSecondaryCat string = `
	CREATE TABLE IF NOT EXISTS secondary_category (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	primary_id INTEGER,
	FOREIGN KEY(primary_id) REFERENCES primary_category(id)
	);`

const createEvent string = `
	CREATE TABLE IF NOT EXISTS event (
	id INTEGER NOT NULL PRIMARY KEY,
	description TEXT,
	start DATETIME,
	end DATETIME,
	customer_id INTEGER,
	primary_id INTEGER,
	secondary_id INTEGER,
	FOREIGN KEY(primary_id) REFERENCES primary_category(id),
	FOREIGN KEY(secondary_id) REFERENCES secondary_category(id),
	FOREIGN KEY(customer_id) REFERENCES customer(id)
	);`

const getCustomers string = `
	SELECT * FROM customer ORDER BY customername
	`

const get30Events string = `
	SELECT * FROM event ORDER BY end LIMIT 30
`

// --space-cadet: #21295cff;
// --yale-blue: #1b3b6fff;
// --lapis-lazuli: #065a82ff;
// --cerulean: #1c7293ff;
// --powder-blue: #9eb3c2ff;
var newDefaultCategories = [][]string{
	{"Presales#124e78", "RFP", "Solution Scope", "Sample Prints", "Sales Demo", "Misc"},
	{"Project#f0f0c9", "Solution Design", "Meeting", "Training Material", "Pilot", "Misc", "Development"},
	{"Training#f2bb05", "Customer", "Sales", "Technical Documentation", "Videos", "Misc"},
	{"Postsales#d74e09", "Meeting", "Technical Support", "Consultancy", "Misc"},
	{"Internal#6e0e0a", "Meeting", "Professional Development", "Process Development", "Admin", "Showroom", "Sales Support", "Technical Support", "Misc"},
	{"Leave#cccccc", "Annual Leave", "Sick Leave", "Birthday Leave", "Volunteer Day", "TIL", "Public Holiday"},
}

type eventData struct {
	id               int64
	start            string
	end              string
	description      string
	customer         string
	primaryLogType   int
	secondaryLogType int
}
