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

	// Traverse tree
	visit(doc)
}

func visit(n *html.Node) {
	// Skip hidden elements
	if n.DataAtom == atom.Script || n.DataAtom == atom.Style {
		return
	}

	// Print text node
	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	// Recurse child(ren)
	if n.FirstChild != nil {
		visit(n.FirstChild)
	}

	// Recurse sibling(s)
	if n.NextSibling != nil {
		visit(n.NextSibling)
	}
}
