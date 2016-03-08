package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-chef/chef"
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

	usageMSG := `ERROR: You need to pass a sub-command (e.g., knife SUB-COMMAND)

Usage: knife sub-command (options)
    -s, --server-url URL             Chef Server URL
        --chef-zero-host HOST        Host to start chef-zero on
        --chef-zero-port PORT        Port (or port range) to start chef-zero on.  Port ranges like 1000,1010 or 8889-9999 will try all given ports until one works.
    -k, --key KEY                    API Client Key
        --[no-]color                 Use colored output, defaults to false on Windows, true otherwise
    -c, --config CONFIG              The configuration file to use
        --defaults                   Accept default values for all questions
    -d, --disable-editing            Do not open EDITOR, just accept the data as is
    -e, --editor EDITOR              Set the editor to use for interactive commands
    -E, --environment ENVIRONMENT    Set the Chef environment (except for in searches, where this will be flagrantly ignored)
    -F, --format FORMAT              Which format to use for output
        --[no-]listen                Whether a local mode (-z) server binds to a port
    -z, --local-mode                 Point knife commands at local repository instead of server
    -u, --user USER                  API Client Username
        --print-after                Show the data after a destructive operation
    -V, --verbose                    More verbose output. Use twice for max verbosity
    -v, --version                    Show chef version
    -y, --yes                        Say yes to all prompts for confirmation
    -h, --help                       Show this message

Available subcommands: (for details, knife SUB-COMMAND --help)
`
	fmt.Println(usageMSG)
}
