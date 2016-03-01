package main

import (
	"flag"
	"fmt"
	"github.com/go-chef/chef"
	"io/ioutil"
	"log"
	"os"
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
		switch args[1] {
		case "list":
			NodeList()
		case "show":
			NodeShow(args[2])
		}
	default:
		usage()
	}
}

// NodeList will fetch and display node list.
func NodeList() {
	allNodes, err := client.Nodes.List()
	if err != nil {
		log.Fatalf("%s", err)
	}
	for k := range allNodes {
		fmt.Println(k)
	}
}

// NodeShow will disply the details of the given node.
func NodeShow(name string) {
	node, err := client.Nodes.Get(name)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("Node Name: %s\n Environment: %s\n Run List: %s\n ", node.Name, node.Environment, node.RunList)
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
