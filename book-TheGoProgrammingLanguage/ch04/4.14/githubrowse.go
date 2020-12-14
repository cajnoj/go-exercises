package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/bugs", bugsHandler)
	http.HandleFunc("/milestones", milestonesHandler)
	http.HandleFunc("/users", usersHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type emptyStruct struct{}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var mainPage = template.Must(template.New("main").Parse(`
	<a href='bugs'>Bugs</a><p/>
	<a href='milestones'>Milestones</a><p/>
	<a href='users'>Users</a>
`))

	err := mainPage.Execute(w, new(emptyStruct))
	if err != nil {
		log.Println(err)
	}
}

const (
	githubURL       = "https://api.github.com"
	githubOwnerRepo = "golang/go"
)

func init() {
	log.Println("Loading bugs from github...")
	loadBugs()

	log.Println("Loading milestones from github...")
	loadMilestones()

	log.Println("Loading users from github...")
	loadUsers()

	log.Println("Finished loading data from github API")

	log.Println("Ready to serve")
}
