package main

import (
	"fmt"
	"log"
)

var cmdNode = SubCommand{
	Name: "Node",
	Usage: `
Available node subcommands: (for details, goknife SUB-COMMAND --help)
** NODE COMMANDS **
goknife node bulk delete REGEX (options)
goknife node create NODE (options)
goknife node delete NODE (options)
goknife node edit NODE (options)
goknife node environment set NODE ENVIRONMENT
goknife node from file FILE (options)
goknife node list (options)
goknife node run_list add [NODE] [ENTRY[,ENTRY]] (options)
goknife node run_list remove [NODE] [ENTRY[,ENTRY]] (options)
goknife node run_list set NODE ENTRIES (options)
goknife node show NODE (options)`,
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

// NodeDelete will delete a node
func NodeDelete(name string) {
	err := client.Nodes.Delete(name)
	if err != nil {
		log.Fatalf("%s\n", err)
	} else {
		fmt.Printf("\nDeleted node [%s]", name)
	}

}
