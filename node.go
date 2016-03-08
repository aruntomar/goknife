package main

import (
	"fmt"
	"log"

	"github.com/go-chef/chef"
)

var cmdNode = SubCommand{
	Name: "Node",
	Usage: `

** NODE COMMANDS **
goknife node create NODE (options)
goknife node delete NODE (options)
goknife node list (options)
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
	fmt.Printf("Node Name: %s\n Environment: %s\n Run List: %s\n", node.Name, node.Environment, node.RunList)
}

// NodeDelete will delete a node
func NodeDelete(name string) {
	err := client.Nodes.Delete(name)
	if err != nil {
		fmt.Printf("Error: Node %s not found\n.", name)
		log.Fatalf("%s\n", err)
	} else {
		fmt.Printf("Deleted node [%s]\n", name)
	}
}

// NodeCreate will create a node.
func NodeCreate(name string) {
	newNode := chef.Node{
		Name:                name,
		Environment:         "_default",
		ChefType:            "node",
		JsonClass:           "Chef::Node",
		RunList:             make([]string, 0),
		AutomaticAttributes: make(map[string]interface{}),
		DefaultAttributes:   make(map[string]interface{}),
		OverrideAttributes:  make(map[string]interface{}),
		NormalAttributes:    make(map[string]interface{}),
	}
	_, err := client.Nodes.Post(newNode)
	if err != nil {
		log.Fatalf("%s\n", err)
	} else {
		fmt.Printf("Created node %s.\n", name)
	}
}
