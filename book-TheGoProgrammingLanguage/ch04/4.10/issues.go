// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		var ageCategory string

		if item.CreatedAt.Before(time.Now().AddDate(-1, 0, 0)) {
			ageCategory = "> 1 Year"
		} else if item.CreatedAt.Before(time.Now().AddDate(0, -1, 0)) {
			ageCategory = "> 1 Month"
		} else {
			ageCategory = "< 1 Month"
		}

		fmt.Printf("#%-5d %9.9s %-9s %.55s\n",
			item.Number, item.User.Login, ageCategory, item.Title)
	}
}
