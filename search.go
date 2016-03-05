package main

import (
	"fmt"
	"log"
)

// ListSearchIndexes will list the indexes that could be searched.
func ListSearchIndexes() {
	indexes, err := client.Search.Indexes()
	if err != nil {
		log.Fatal("Could not list indexes: ", err)
	}
	for k := range indexes {
		fmt.Println(k)
	}
}
