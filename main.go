package main

import (
	"net/http"
	"text/template"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"github.com/PuerkitoBio/goquery"
)

var article = ArticleJson{}
var projectTodos = map[string]TodoList{}


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
	tmpl, _ := template.ParseFiles("static/templates/articles.html", "static/templates/header.html", "static/articles/go_client_communication.html")

	content, err := ioutil.ReadFile("static/articles/golang_js_communication")
	if err != nil {
		fmt.Println(err) // todo
	}

	if err := json.Unmarshal(content, &article); err != nil {
		panic(err) // todo
	}

	tmpl.Execute(w, article)
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/contact.html", "static/templates/header.html")

	formSubmitted := false
	if r.Method == "POST" {
		//handleSend(r)
		formSubmitted = true
	}

	tmpl.Execute(w, formSubmitted)
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/templates/about.html", "static/templates/header.html")
	tmpl.Execute(w, nil)
}

//func handleSend(r *http.Request) {
//	fmt.Println("method:", r.Method) //get request method
//	r.ParseForm()
//	SendMail(r.Form["subject"][0], r.Form["name"][0], r.Form["email"][0], r.Form["message"][0])
//}

func main() {

	getTodoList("oak", "https://github.com/ryanmccauley211/Oak/blob/master/README.md")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/projects/", handleProjects)
	http.HandleFunc("/articles/", handleArticles)
	http.HandleFunc("/contact/", handleContact)
	http.HandleFunc("/about/", handleAbout)

	fmt.Println("Listening...")
	http.ListenAndServe(":8080", nil)
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