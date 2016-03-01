package main

import (
	"fmt"
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
