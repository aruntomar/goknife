package main

import (
	"fmt"
	"github.com/go-chef/chef"
	"log"
)

// DataBagList will list the databags
func DataBagList() {
	databags, err := client.DataBags.List()
	if err != nil {
		log.Fatalf("%s", err)
	}
	for k := range *databags {
		fmt.Printf("%v\n", k)
	}
}

// DataBagCreate will create the databag
func DataBagCreate(dbname string) {
	_, err := client.DataBags.Create(&chef.DataBag{Name: dbname})
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("Created data_bag[%s]\n", dbname)
}

// DataBagDelete will create the databag
func DataBagDelete(dbname string) {
	_, err := client.DataBags.Delete(dbname)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("Deleted data_bag[%s]\n", dbname)
}

// DataBagShow will show the detailed contents of databag
func DataBagShow(dbname string) {
	items, err := client.DataBags.ListItems(dbname)
	if err != nil {
		log.Fatalf("%s", err)
	}
	for k := range *items {
		fmt.Printf("%v\n", k)
	}
}
