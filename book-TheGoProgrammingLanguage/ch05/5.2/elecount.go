package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	// Counts map
	counts := make(map[string]int)

	// Traverse tree
	visit(counts, doc)

	// Print map
	for k, v := range counts {
		fmt.Println(k, v)
	}
}

func visit(counts map[string]int, n *html.Node) {
	// Increment DataAtom cell
	counts[n.DataAtom.String()]++

	// Recurse child(ren)
	if n.FirstChild != nil {
		visit(counts, n.FirstChild)
	}

	// Recurse sibling(s)
	if n.NextSibling != nil {
		visit(counts, n.NextSibling)
	}
}
