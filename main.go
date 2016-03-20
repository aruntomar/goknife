package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aruntomar/chef"
)

var (
	client *chef.Client
)

// SubCommand struct is to identify different subcommands and run them.
type SubCommand struct {
	// name of the subcommand
	Name string
	// Description
	Description string
	// Usage
	Usage string
}

func init() {
	//get the chef config from env params
	clientName := os.Getenv("CLIENT_NAME")
	clientKey := os.Getenv("CLIENT_KEY")
	baseURL := os.Getenv("CHEF_SERVER_URL")

	// fmt.Printf("\n Client name: %s\n Client Key: %s\n BaseURL: %s\n", clientName, clientKey, baseURL)

	// read the client key
	key, err := ioutil.ReadFile(clientKey)
	if err != nil {
		fmt.Println("Couldn't read key", err)
		os.Exit(1)
	}

	// build a client
	client, err = chef.NewClient(&chef.Config{
		Name:    clientName,
		Key:     string(key),
		BaseURL: baseURL,
		// SkipSSL: true,
	})
	if err != nil {
		fmt.Println("error setting up client", err)
		os.Exit(1)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		usage()
		os.Exit(1)
	}

	switch args[0] {

	case "node":
		listOfArgs := args[1:]
		if len(listOfArgs) == 0 {
			fmt.Println(cmdNode.Usage)
		} else {
			switch args[1] {
			case "list":
				NodeList()
			case "show":
				NodeShow(args[2])
			case "delete":
				if len(listOfArgs) >= 2 {
					var reallyDelete string
					fmt.Printf("Do you really want to delete node %s (Y/N)", args[2])
					fmt.Scanln(&reallyDelete)
					if strings.ToUpper(reallyDelete) == "Y" {
						NodeDelete(args[2])
					}
				} else {
					log.Fatalln("Fatal: You must specify a node name")
				}
			case "create":
				if len(listOfArgs) >= 2 {
					NodeCreate(args[2])
				} else {
					log.Fatalln("Fatal: You must specify a node name")
				}
			}
		}
	case "data":
		switch args[2] {
		case "list":
			DataBagList()
		case "create":
			DataBagCreate(args[3])
		case "delete":
			DataBagDelete(args[3])
		case "show":
			if len(args[3:]) > 1 {
				DataBagGetItem(args[3], args[4])
			} else {
				DataBagShow(args[3])
			}
		}
	case "from":
		for _, fname := range args[5:] {
			jsonData, err := ioutil.ReadFile(fname)
			if err != nil {
				log.Fatalf("%v\n", err)
			}
			var rnd map[string]string
			err = json.Unmarshal(jsonData, &rnd)
			if err != nil {
				log.Fatalf("%v\n", err)
			}
			err = DataBagCreateItem(args[4], rnd)
			if err != nil {
				// DataBagUpdateItem(args[4], rnd["id"], rnd)
			}
		}
	case "client":
		listOfArgs := args[1:]
		// fmt.Println("List of args: ", listOfArgs)
		if len(listOfArgs) == 0 {
			fmt.Println(cmdClient.Usage)
		} else {
			switch args[1] {
			case "list":
				ClientList()
			case "show":
				if len(listOfArgs) >= 2 {
					ClientShow(args[2])
				} else {
					log.Fatalln("Fatal: You must specify a client name")
				}
			case "delete":
				if len(listOfArgs) >= 2 {
					var reallyDelete string
					fmt.Printf("Do you really want to delete client %s (Y/N)", args[2])
					fmt.Scanln(&reallyDelete)
					if strings.ToUpper(reallyDelete) == "Y" {
						ClientDelete(args[2])
					}
				} else {
					log.Fatalln("Fatal: You must specify a client name")
				}
			case "create":
				if len(listOfArgs) >= 2 {
					ClientCreate(args[2], false)
				} else {
					log.Fatalln("Fatal: You must specify a client name")
				}
			}
		}

	case "role":
		listOfArgs := args[1:]
		// fmt.Println("List of args: ", listOfArgs)
		if len(listOfArgs) == 0 {
			fmt.Println(cmdRole.Usage)
		} else {
			switch args[1] {
			case "list":
				RoleList()
			case "show":
				if len(listOfArgs) >= 2 {
					RoleShow(args[2])
				} else {
					log.Fatalln("Fatal: You must specify a role name")
				}
			case "delete":
				if len(listOfArgs) >= 2 {
					var reallyDelete string
					fmt.Printf("Do you really want to delete role %s (Y/N)", args[2])
					fmt.Scanln(&reallyDelete)
					if strings.ToUpper(reallyDelete) == "Y" {
						RoleDelete(args[2])
					}
				} else {
					log.Fatalln("Fatal: You must specify a role name")
				}
			case "create":
				if len(listOfArgs) >= 2 {
					RoleCreate(args[2])
				} else {
					log.Fatalln("Fatal: You must specify a role name")
				}
			}
		}

	case "cookbook":
		listOfArgs := args[1:]
		// fmt.Println("List of args: ", listOfArgs)
		if len(listOfArgs) == 0 {
			fmt.Println(cmdCookbook.Usage)
		} else {
			switch args[1] {
			case "list":
				CookbookList()
			case "show":
				if len(listOfArgs) >= 2 {
					CookbookShow(args[2], args[3])
				} else {
					log.Fatalln("Fatal: You must specify a cookbook name")
				}
			case "delete":
				if len(listOfArgs) >= 2 {
					listOfCb, err := client.Cookbooks.GetAvailableVersions(args[2], "")
					if err != nil {
						log.Fatalf("Error: Cannot find cookbook named %s to delete", args[2])
					}
					for k, v := range listOfCb {
						if len(v.Versions) == 1 {
							CookbookDelete(args[2], v.Versions[0].Version)
						} else {
							var n int
							fmt.Printf("Which version(s) do you want to delete? \n")
							for i, value := range v.Versions {
								fmt.Printf("%d. %s %s\n", i+1, k, value.Version)
							}
							deleteAllVersions := len(v.Versions) + 1
							fmt.Printf("%d. All Versions\n", deleteAllVersions)
							fmt.Scanln(&n)
							if n == deleteAllVersions {
								for _, val := range v.Versions {
									CookbookDelete(args[2], val.Version)
								}
							} else {
								CookbookDelete(args[2], v.Versions[int(n)-1].Version)
							}
						}
					}
				} else {
					log.Fatalln("Fatal: You must provide the name of the cookbook to delete.")
				}
			case "create":
				// 	if len(listOfArgs) >= 2 {
				// 		RoleCreate(args[2])
				// 	} else {
				// 		log.Fatalln("Fatal: You must specify a role name")
				// 	}
			}
		}

	case "search":
		if len(args[1:]) > 1 {
			SearchQuery(args[1], args[2])
		} else {
			ListSearchIndexes()
		}
	default:
		usage()
	}
}

func usage() {
	msg := `
ERROR: You need to pass sub-command (e.g., goknife SUB-COMMAND)

USAGE: knife sub-command (options)

Available subcommands:
	`
	usageMSG := msg + cmdSearch.Usage + cmdDataBag.Usage + cmdClient.Usage + cmdNode.Usage + cmdRole.Usage + cmdCookbook.Usage
	fmt.Println(usageMSG)
}
