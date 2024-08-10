package main

import (
	"fmt"
	"net/http"
  "html/template"
	"hbw/views"
  "hbw/db"
  "time"
)

var index *views.View
var contact *views.View
var projects_page *views.View

type Project struct {
	Id int64
	Title string
	Description string
	Created_at time.Time
}

func GetAllProjects(w http.ResponseWriter) {
  projects, err := db.GetProjects()
  if err != nil {
    fmt.Println("Error getting projects from DB", err)
  }
  t, _ := template.ParseFiles("views/projects.html")
  err = t.Execute(w, projects)
  if err != nil {
    fmt.Println("Error executing projects template with projects", err)
  }
}

func main() {
	fmt.Println("Starting Server on port 3000")
	
	index = views.NewView("bootstrap", "views/index.html")
	projects_page = views.NewView("bootstrap", "views/projects.html")
	contact = views.NewView("bootstrap", "views/contact.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/createProject", create_projects_handler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Page")
	index.Render(w, nil)
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Projects Page")
  all_projects, err := db.GetProjects()
  if err != nil {
    fmt.Println("Error Getting All Projects from DB", err)
  }
	projects_page.Render(w, all_projects)
}

func create_projects_handler(w http.ResponseWriter, r *http.Request) {
  title := r.FormValue("title")
  description := r.FormValue("description")
  err := db.CreateProject(title, description)
  if err != nil {
    fmt.Println("create_projects_handler failed to add project to db", err)
  }
  http.Redirect(w, r, "/projects", http.StatusSeeOther)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Contact Page")
	contact.Render(w, nil)
}
