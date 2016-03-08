package main

import (
	"fmt"
	"log"

	"github.com/go-chef/chef"
)

var cmdRole = SubCommand{
	Name: "Roles",
	Usage: `
Available role subcommands: (for details, goknife SUB-COMMAND --help)

** ROLE COMMANDS **
goknife role bulk delete REGEX (options)
goknife role create ROLE (options)
goknife role delete ROLE (options)
goknife role edit ROLE (options)
goknife role env_run_list add [ROLE] [ENVIRONMENT] [ENTRY[,ENTRY]] (options)
goknife role env_run_list clear [ROLE] [ENVIRONMENT]
goknife role env_run_list remove [ROLE] [ENVIRONMENT] [ENTRIES]
goknife role env_run_list replace [ROLE] [ENVIRONMENT] [OLD_ENTRY] [NEW_ENTRY]
goknife role env_run_list set [ROLE] [ENVIRONMENT] [ENTRIES]
goknife role from file FILE [FILE..] (options)
goknife role list (options)
goknife role run_list add [ROLE] [ENTRY[,ENTRY]] (options)
goknife role run_list clear [ROLE]
goknife role run_list remove [ROLE] [ENTRY]
goknife role run_list replace [ROLE] [OLD_ENTRY] [NEW_ENTRY]
goknife role run_list set [ROLE] [ENTRIES]
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
	fmt.Println("#TODO")
}
