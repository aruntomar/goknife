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
goknife cookbook delete COOKBOOK VERSION
goknife cookbook show COOKBOOK [VERSION]
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
	}
}

// CookbookDelete will delete the cookbook.
func CookbookDelete(name, version string) {
	err := client.Cookbooks.Delete(name, version)
	if err != nil {
		log.Fatalf("Error: Cannot find cookbook named %s to delete", name)
	} else {
		fmt.Printf("Deleted cookbook [%s] [%s]\n", name, version)
	}
}

// CookbookShow will display cookbook details for a specific version
func CookbookShow(name, version string) {
	cb, err := client.Cookbooks.GetVersion(name, version)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	fmt.Printf("attributes: %s\nchef_type: %s\ncookbook_name: %s\n", cb.Attributes, cb.ChefType, cb.Name)
}
