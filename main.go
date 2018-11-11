package main

import (
	"net/http"
	"text/template"
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"os"
)

var article = ArticleJson{}
var projectTodos = map[string]TodoList{}
var email string;
var pass  string;

var contactMap = map[string]interface{}{
	"email": "ryanmccauley211@gmail.com",
	"github": "https://github.com/ryanmccauley211",
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/home.html", "static/templates/header.html")
	tmpl.Execute(w, nil)
}

func handleProjects(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/projects.html", "static/templates/header.html", "static/projects/oak.html")
	tmpl.Execute(w, projectTodos)
}

func handleArticles(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("static/templates/articles.html", "static/templates/header.html", "static/articles/git_part1_introduction.html", "static/articles/go_client_communication.html", "static/articles/a-frame_intro.html", "static/articles/binary_decimal_and_hex.html")

	tmpl.Execute(w, article)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/contact.html", "static/templates/header.html")

	formSubmitted := false
	if r.Method == "POST" {
		handleSend(r)
		formSubmitted = true
	}

	tmpl.Execute(w, formSubmitted)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/about.html", "static/templates/header.html")
	tmpl.Execute(w, nil)
}

func handleSend(r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	r.ParseForm()
	name := r.Form["first_name"][0] + " " + r.Form["last_name"][0];
	SendMail(r.Form["subject"][0], name, r.Form["email"][0], r.Form["message"][0], email, pass)
}

func handleTagRequests(w http.ResponseWriter, r *http.Request) {

}

func main() {

	email = os.Args[1];
	pass  = os.Args[2];

	getTodoList("oak", "https://github.com/ryanmccauley211/Oak/blob/master/README.md")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/projects", handleProjects)
	http.HandleFunc("/articles", handleArticles)
	http.HandleFunc("/contact", handleContact)
	http.HandleFunc("/about", handleAbout)

	http.HandleFunc("/tag", handleTagRequests)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe("localhost:8040", nil))
}

func getTodoList(projectName string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		checkedTasks := []string{}
		uncheckedTasks := []string{}
		delTasks := []string{}
		doc.Find(".contains-task-list li").Each(func(i int, s *goquery.Selection) {
			input := s.Find("input")
			_, checked := input.Attr("checked")
			del := s.Find("del")

			if checked {
				checkedTasks = append(checkedTasks, s.Text())
			} else if del.Size() > 0 {
				delTasks = append(delTasks, del.Text())
			} else {
				uncheckedTasks = append(uncheckedTasks, s.Text())
			}
			projectTodos[projectName] = TodoList{checkedTasks, uncheckedTasks, delTasks, true}
		})
	}
	return nil
}

type ArticleJson struct {
	Title    string   `json:"title"`
	Date     string   `json:"date"`
	Tags     []string `json:"tags"`
	ImageUrl string   `json:"image"`
	Content  []map[string][]string `json:"content"`
}

type TodoList struct {
	Checked   []string
	Unchecked []string
	Deleted   []string
	Loaded      bool
}