package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func searchComic(comic *Comic, s string) bool {
	if comic.Month == s ||
		strconv.Itoa(comic.Num) == s ||
		strings.Contains(comic.Link, s) ||
		comic.Year == s ||
		strings.Contains(comic.News, s) ||
		strings.Contains(comic.SafeTitle, s) ||
		strings.Contains(comic.Transcript, s) ||
		strings.Contains(comic.Alt, s) ||
		strings.Contains(comic.Img, s) ||
		strings.Contains(comic.Title, s) ||
		comic.Day == s {
		return true
	}
	return false
}

func main() {
	var err error
	f, err := os.Open("data.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	reader := bufio.NewReader(f)

	var comics []*Comic
	err = json.NewDecoder(reader).Decode(&comics)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	for _, comic := range comics {
		if searchComic(comic, os.Args[1]) {
			fmt.Println(comic.Num)
		}
	}
}
