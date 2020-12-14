package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type user struct {
	ID        int
	Login     string
	HTMLURL   string `json:"html_url"`
	SiteAdmin bool   `json:"site_admin"`
}

var users []*user

func loadUsers() {
	const issuesURL = githubURL + "/repos/" + githubOwnerRepo + "/assignees"
	var (
		err  error
		resp *http.Response
		body []byte
	)

	resp, err = http.Get(issuesURL)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		panic(fmt.Errorf("Error: search query failed: %s", resp.Status))
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		panic(err)
	}
	resp.Body.Close()

	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to parse response: %v\n", err)
		os.Exit(5)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var usersList = template.Must(template.New("userslist").Parse(`
	<h1>Users</h1>
	<table>
	<tr style='text-align: left'>
	  <th>#</th>
	  <th>Login Name</th>
	  <th>Site Admin</th>
	</tr>
	{{range .}}
	<tr>
	  <td><a href='{{.HTMLURL}}'>{{.ID}}</a></td>
	  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
	  <td>{{.SiteAdmin}}</td>
	</tr>
	{{end}}
	</table>
	`))

	err := usersList.Execute(w, users)
	if err != nil {
		log.Println(err)
	}
}
