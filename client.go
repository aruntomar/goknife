package main

import (
	"fmt"
	"log"
)

var cmdClient = SubCommand{
	Name: "Client",
	Usage: `

** CLIENT COMMANDS **
goknife client create CLIENTNAME (options)
goknife client delete CLIENT (options)
goknife client list (options)
goknife client show CLIENT (options)`,
}

// ClientList will list all the clients on the chef server
func ClientList() {
	cls, err := client.Clients.List()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for k := range cls {
		fmt.Println(k)
	}
}

// ClientShow will display the details about a particular client.
func ClientShow(name string) {
	myclient, err := client.Clients.Get(name)
	if err != nil {
		log.Fatalln("Error: The object you are looking for could not be found.")
		log.Fatalf("%s\n", err)
	}
	fmt.Printf(" admin: %t\n chef_type: %s\n name: %s\n validator: %t\n", myclient.Admin, myclient.ChefType, myclient.Name, myclient.Validator)
}

// ClientDelete will delete the provided client
func ClientDelete(name string) {
	// fmt.Println(name)
	err := client.Clients.Delete(name)
	if err != nil {
		log.Fatalf("%s\n", err)
	} else {
		fmt.Printf("Deleted client [%s]", name)
	}
}

// ClientCreate will create a new client.
func ClientCreate(name string, admin bool) {
	result, err := client.Clients.Create(name, admin)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("Created client [%s]\n", name)
	fmt.Println(result.PrivateKey)
}
