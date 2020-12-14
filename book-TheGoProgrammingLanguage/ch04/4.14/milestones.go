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

type milestone struct {
	Number  int
	Title   string
	HTMLURL string `json:"html_url"`
}

var milestones []*milestone

func loadMilestones() {
	const issuesURL = githubURL + "/repos/" + githubOwnerRepo + "/milestones"
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

	err = json.Unmarshal(body, &milestones)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to parse response: %v\n", err)
		os.Exit(5)
	}
}

func milestonesHandler(w http.ResponseWriter, r *http.Request) {
	var milestonesList = template.Must(template.New("milestoneslist").Parse(`
	<h1>Milestones</h1>
	<table>
	<tr style='text-align: left'>
	  <th>#</th>
	  <th>Title</th>
	</tr>
	{{range .}}
	<tr>
	  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

	err := milestonesList.Execute(w, milestones)
	if err != nil {
		log.Println(err)
	}
}
