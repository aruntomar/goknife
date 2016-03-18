package main

import (
	"fmt"
	"log"
	// "github.com/go-chef/chef"
	"chef"
)

var cmdDataBag = SubCommand{
	Name: "DataBag",
	Usage: `

** DATA BAG COMMANDS **
goknife data bag create BAG [ITEM] (options)
goknife data bag delete BAG [ITEM] (options)
goknife data bag from file BAG FILE|FOLDER [FILE|FOLDER..] (options)
goknife data bag list (options)
goknife data bag show BAG [ITEM] (options)
	`,
}

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

// DataBagGetItem will display the item details.
func DataBagGetItem(dbname, dbitem string) {
	item, err := client.DataBags.GetItem(dbname, dbitem)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	value := item.(map[string]interface{})
	for k, v := range value {
		fmt.Printf("%s : %s\n", k, v)
	}
	// fmt.Printf("%v\n", item)
}

// DataBagCreateItem creates a data bag item
func DataBagCreateItem(dbname string, dbitem chef.DataBagItem) (err error) {
	err = client.DataBags.CreateItem(dbname, dbitem)
	if err != nil {
		item := dbitem.(map[string]string)
		DataBagUpdateItem(dbname, item["id"], dbitem)
		// log.Fatalf("%s\n", err)
		return err
	}
	fmt.Printf("Updated data_bag_item[%s::%s]\n", dbname, dbitem)
	return nil
}

// DataBagUpdateItem will update the db item if already exists.
func DataBagUpdateItem(dbname string, itemid string, dbitem chef.DataBagItem) {
	err := client.DataBags.UpdateItem(dbname, itemid, dbitem)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("Updated data_bag_item[%s::%s]\n", dbname, dbitem)
}
