package db

import (
	"log"
	"fmt"
)

func DropDB(name string) {
	if (!doesDBExist(name)) {
		fmt.Printf("Database %s does not exist \n", name)
		return;
	}
	_, err := DB.Exec("DROP DATABASE " + name)
	if (err != nil) {
		log.Fatalf("failed to drop database " + name + ". %v", err.Error())
	}
	fmt.Printf("Database %s dropped successfully.\n", name)
}

func DropTable(name string) {
	if (!doesTableExist(name)) {
		fmt.Printf("Table %s does not exist \n", name)
		return;
	}
	_, err := DB.Exec("DROP TABLE " + name)
	if (err != nil) {
		log.Fatalf("failed to drop table " + name + ". %v", err.Error())
	}
	fmt.Printf("Table %s dropped successfully.\n", name)
}