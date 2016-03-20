package main

import (
	"fmt"
	"github.com/aruntomar/chef"
	"log"
)

var cmdRole = SubCommand{
	Name: "Roles",
	Usage: `

** ROLE COMMANDS **
goknife role create ROLE (options)
goknife role delete ROLE (options)
goknife role list (options)
goknife role show ROLE (options)
  `,
}

// RoleCreate will create a role
func RoleCreate(name string) {
	newRole := chef.Role{
		Name:               name,
		ChefType:           "role",
		Description:        "",
		RunList:            make([]string, 0),
		DefaultAttributes:  make(map[string]interface{}),
		OverrideAttributes: make(map[string]interface{}),
		JsonClass:          "Chef::Role",
	}
	result, err := client.Roles.Create(&newRole)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf("%v\n", result)
}

// RoleList will list all the roles
func RoleList() {
	roles, err := client.Roles.List()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	for k := range *roles {
		fmt.Println(k)
	}
}

// RoleShow will display role details.
func RoleShow(name string) {
	role, err := client.Roles.Get(name)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	fmt.Printf(" chef_type:\t %s \n default_attributes: \t %s\n description: \t %s\n name: \t %s\n run_list: \t %s\n override_attributes: \t %s\n json_class: \t %s\n",
		role.ChefType, role.DefaultAttributes, role.Description, role.Name, role.RunList, role.OverrideAttributes, role.JsonClass)
}

// RoleDelete will delete a role
func RoleDelete(name string) {
	err := client.Roles.Delete(name)
	if err != nil {
		fmt.Printf("Role %s could not be found.", name)
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Role %s deleted.\n", name)
	}
}
