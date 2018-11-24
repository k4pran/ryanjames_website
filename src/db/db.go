package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

var DB *sql.DB;

func Connect(name string) {
	var err error;
	DB, err = sql.Open("mysql", "root:pazzwort9131@/sys")
	if (err != nil) {
		log.Fatalf("failed to connect to database. %v", err.Error())
		DB.Close();
	}

	if (!doesDBExist(name)) {
		CreateDB(name)
		DB.Close()
		DB, err = sql.Open("mysql", "root:pazzwort9131@/" + name)
		if (err != nil) {
			log.Fatalf("failed to connect to database. %v", err.Error())
			DB.Close();
		}
	}

	defer DB.Close()

	CreateAllTables()
	fmt.Println("-------------------------------")

	articlesInfo := parseArticles()
	updateArticlesTable(articlesInfo)
	fmt.Println("-------------------------------")

}

func parseArticles() []*ArticleInfo {
	var articlesInfo []*ArticleInfo
	file, _ := ioutil.ReadFile("static/mini_db/articles")
	err := json.Unmarshal([]byte(file), &articlesInfo)
	if (err != nil) {
		log.Fatalf("Unable to parse json tag information. %v", err.Error())
	}
	return articlesInfo
}


func updateArticlesTable(articlesInfo []*ArticleInfo) {
	for _, data := range articlesInfo {

		stmt, err := DB.Prepare("INSERT INTO articles (article_id, title, published) VALUES (?, ?, ?) " +
			"ON DUPLICATE KEY " +
				"UPDATE article_id=?, title=?, published=?")
		if (err != nil) {
			log.Fatal("failed to prepare sql insert into articles.", err.Error())
		}

		_, err = stmt.Exec(data.Ref, data.FoundIn, data.Date,
							  data.Ref, data.FoundIn, data.Date) // for updates on duplicates
		if (err != nil) {
			log.Fatal("failed to execute sql insert into articles.", err.Error())
		}

		if (err != nil) {
			log.Fatal("failed to retrieve row inserted id. ", err.Error())
		}
		fmt.Printf("Article inserted at row %s\n", data.Ref)

		updateTagsTable(data.Ref, data.Tags)
	}
}

func updateTagsTable(articleID string, tags []string) {

	for _, tag := range tags {
		stmt, err := DB.Prepare("INSERT INTO tags (tag, article_id) VALUES (?, ?) " +
			"ON DUPLICATE KEY " +
				"UPDATE tag=?, article_id=?")
		if (err != nil) {
			log.Fatal("failed to prepare sql insert into tags.", err.Error())
		}

		_, err = stmt.Exec(tag, articleID,
			               tag, articleID) // for updates on duplicates
		if (err != nil) {
			log.Fatal("failed to execute sql insert into tags.", err.Error())
		}
		fmt.Printf("Tag inserted at row %s\n", articleID)
	}
	fmt.Println()
}

type ArticleInfo struct {
	Ref      string   `json:"ref"`
	FoundIn  string   `json:"found_in"`
	Date     string   `json:"date"`
	Tags     []string `json:"tags"`
}