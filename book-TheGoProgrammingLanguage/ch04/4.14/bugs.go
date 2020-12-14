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

type bug struct {
	Number  int
	Title   string
	HTMLURL string `json:"html_url"`
}

var bugs []*bug

func loadBugs() {
	const issuesURL = githubURL + "/repos/" + githubOwnerRepo + "/issues"
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

	err = json.Unmarshal(body, &bugs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to parse response: %v\n", err)
		os.Exit(5)
	}
}

func bugsHandler(w http.ResponseWriter, r *http.Request) {
	var bugsList = template.Must(template.New("bugslist").Parse(`
	<h1>Bugs</h1>
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

	err := bugsList.Execute(w, bugs)
	if err != nil {
		log.Println(err)
	}
}
