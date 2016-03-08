package main

import (
	"fmt"
	"log"
	"sort"
)

var cmdSearch = SubCommand{
	Name: "Search",
	Usage: `
** SEARCH COMMANDS **
goknife search INDEX QUERY (options)
`,
}

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

// SearchQuery will search the indexes for the provided query.
func SearchQuery(index, statement string) {
	q, err := client.Search.NewQuery(index, statement)
	if err != nil {
		log.Fatalln("Problem with building search query: ", err)
	}
	res, err := q.Do(client)
	if err != nil {
		log.Fatalln("Error running query: ", err)
	}
	fmt.Printf("%d results found\n", res.Total)
	// fmt.Printf("%v\n", res.Rows)
	for i := range res.Rows {
		results := res.Rows[i].(map[string]interface{})
		keys := []string{}
		for key := range results {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for i := range keys {
			key := keys[i]
			value := results[key]
			fmt.Printf("%s \t %s \n", key, value)
		}
		fmt.Printf("\n")
	}
}
