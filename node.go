package main

import (
	"fmt"
	"log"
)

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
