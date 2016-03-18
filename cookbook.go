package main

import (
	"fmt"
	"log"
)

var cmdCookbook = SubCommand{
	Name: "Cookbook",
	Usage: `

** COOKBOOK COMMANDS **
goknife cookbook list
`,
}

// CookbookList will list the cookbooks on the chef server.
func CookbookList() {
	cbList, err := client.Cookbooks.List()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	for k, v := range cbList {
		for _, i := range v.Versions {
			fmt.Printf("%s\t\t\t\t\t%s\n", k, i.Version)
		}
		// fmt.Printf("%s\t%s\n", k, v.Versions)
	}
}
