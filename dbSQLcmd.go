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

var newDefaultCategories = [][]string{
	{"Presales#3fe06a", "RFP", "Solution Scope", "Sample Prints", "Sales Demo"},
	{"Project#3fe0d3", "Solution Design", "Meeting", "Training Material", "Pilot"},
	{"Training#b23fe0", "Customer", "Sales", "Technical", "Videos"},
	{"Postsales#e0873f", "Meeting", "Technical Support", "Consultancy"},
	{"Internal#e0d83f", "Meeting", "PD", "Process Development", "Admin", "Showroom", "Sales Support"},
	{"Leave#666664", "Annual Leave", "Sick Leave", "Birthday Leave", "Volunteer Day", "TIL"},
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
