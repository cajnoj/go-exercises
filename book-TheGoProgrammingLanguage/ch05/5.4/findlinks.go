package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

var elementLinkKeys map[atom.Atom]string

func init() {
	elementLinkKeys = map[atom.Atom]string{
		atom.A:      "href",
		atom.Img:    "src",
		atom.Script: "src",
		atom.Link:   "href",
	}
}

func visit(links []string, n *html.Node) []string {
	var ok bool
	var key string
	if key, ok = elementLinkKeys[n.DataAtom]; ok {
		for _, a := range n.Attr {
			if a.Key == key {
				links = append(links, a.Val)
			}
		}
	}

	// Recurse child(ren)
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	// Recurse sibling(s)
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}
