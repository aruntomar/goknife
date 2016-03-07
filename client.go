package main

import (
	"fmt"
	"log"
)

var cmdClient = SubCommand{
	Name: "Client",
	Usage: `
Available client subcommands: (for details, knife SUB-COMMAND --help)

** CLIENT COMMANDS **
knife client bulk delete REGEX (options)
knife client create CLIENTNAME (options)
knife client delete CLIENT (options)
knife client edit CLIENT (options)
Usage: /usr/bin/knife (options)
knife client key delete CLIENT KEYNAME (options)
knife client key edit CLIENT KEYNAME (options)
knife client key list CLIENT (options)
knife client key show CLIENT KEYNAME (options)
knife client list (options)
knife client reregister CLIENT (options)
knife client show CLIENT (options)`,
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
		log.Fatalf("%s\n", err)
	}
	fmt.Printf(" admin: %t\n chef_type: %s\n name: %s\n validator: %t\n", myclient.Admin, myclient.ChefType, myclient.Name, myclient.Validator)
}
