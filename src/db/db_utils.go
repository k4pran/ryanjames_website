package db

import (
	"fmt"
	"log"
)

func doesDBExist(name string) bool {
	result, err := DB.Query("SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME LIKE '" + name + "'")
	if (err != nil) {
		fmt.Printf("failed to check if database " + name + " exists. %v", err.Error())
	}

	var database string
	for result.Next() {
		result.Scan(&database)
		if (database == name) {
			return true;
		}
	}
	return false;
}

func doesTableExist(name string) bool {
	result, err := DB.Query("SHOW TABLES LIKE '" + name + "'")
	if (err != nil) {
		log.Fatal("failed to check if table " + name + " exists. %v", err.Error())
	}

	var table string
	for result.Next() {
		result.Scan(&table)
		if (table == name) {
			return true;
		}
	}
	return false;
}