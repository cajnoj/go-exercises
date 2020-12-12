package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	apiKeyFileName = "apikey.secret"
	apiURL         = "http://www.omdbapi.com/"
)

func getAPIKey() string {
	f, err := os.Open(apiKeyFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Expecting api key file %s.\n", apiKeyFileName)
		os.Exit(1)
	}

	var apiKey string
	_, err = fmt.Fscan(f, &apiKey)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to parse api key")
		os.Exit(2)
	}

	return apiKey
}

type movieSearchResult struct {
	Response  string
	Error     string
	Title     string
	PosterURL string `json:"poster"`
}

func main() {
	// Read API key from secret file
	apiKey := getAPIKey()

	// Get search string
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error: Expecting single parameter.")
		os.Exit(3)
	}
	s := os.Args[1]

	// Request title
	q := apiURL + "/?t=" + url.QueryEscape(s) + "&apikey=" + apiKey
	resp, err := http.Get(q)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		panic(fmt.Errorf("Error: search query failed: %s", resp.Status))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		panic(err)
	}
	resp.Body.Close()

	// Convert response body to json
	var movie movieSearchResult
	err = json.Unmarshal(body, &movie)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to parse response: %v\n", err)
		os.Exit(5)
	}

	if movie.Response != "True" {
		fmt.Fprintf(os.Stderr, "Error: Search returned false reponse: %s\n", movie.Error)
		os.Exit(6)
	}

	imageFileName := strings.Replace(movie.Title, " ", "_", -1) + ".jpg"

	// Request poster image
	resp, err = http.Get(movie.PosterURL)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		panic(fmt.Errorf("Error: Poster download failed: %s", resp.Status))
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		resp.Body.Close()
		panic(err)
	}
	resp.Body.Close()

	// Write poster image
	imageFile, err := os.Create(imageFileName)
	_, err = imageFile.Write(body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Done writing poster image file %q\n", imageFileName)
}
