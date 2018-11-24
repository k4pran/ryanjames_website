package db

import (
	"log"
	"fmt"
)

func CreateDB(name string) {
	if (doesDBExist(name)) {
		fmt.Printf("Database %s already exists\n", name)
		return;
	}

	_, err := DB.Exec("CREATE DATABASE " + name)
	if (err != nil) {
		log.Fatalf("failed to create database " + name + ". %v", err.Error())
	}
	fmt.Printf("Database %s created successfully.\n", name)
}

func CreateAllTables() {
	CreateArticlesTable()
	CreateTagsTable()
}

func CreateArticlesTable() {
	if (doesTableExist("articles")) {
		fmt.Println("Table 'articles' already exists")
		return;
	}

	_, err := DB.Exec("CREATE TABLE articles " +
		"(article_id INT NOT NULL, " +
		 "title VARCHAR(255) NOT NULL, " +
		 "published DATE, " +
		 "PRIMARY KEY (article_id))")
	if (err != nil) {
		log.Fatal("failed to create articles table.", err.Error())
	}
	fmt.Printf("Table articles created successfully\n")
}

func CreateTagsTable() {
	if (doesTableExist("tags")) {
		fmt.Println("Table 'tags' already exists")
		return;
	}

	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS tags " +
		"(tag_id INT NOT NULL AUTO_INCREMENT, " +
		"tag VARCHAR(255) NOT NULL, " +
		"article_id INT NOT NULL, " +
		"FOREIGN KEY (article_id) REFERENCES articles(article_id), " +
		"PRIMARY KEY (tag_id))")
	if (err != nil) {
		log.Fatal("failed to create tags table.", err.Error())
	}
	fmt.Printf("Table tags created successfully\n")
}