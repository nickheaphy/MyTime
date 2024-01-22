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
	name TEXT,
	start DATETIME,
	end DATETIME,
	primary_id INTEGER,
	secondary_id INTEGER,
	FOREIGN KEY(primary_id) REFERENCES primary_category(id),
	FOREIGN KEY(secondary_id) REFERENCES secondary_category(id)
	);`

var defaultCategories = map[string][]string{
	"Presales#ff00": {
		"RFP",
		"Solution Scope",
	},
	"Project#aabb": {
		"Solution Design",
		"Meeting",
		"Training Material",
	},
	"Postsales#1111": {
		"Meeting",
		"Technical Help",
	},
	"Internal#1234": {
		"Training Video",
		"Service API",
		"Print Hub",
	},
	"Leave#7666": {
		"Annual Leave",
		"Sick Leave",
		"Birthday Leave",
		"Volunteer Day",
		"TIL",
	},
}
