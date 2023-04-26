package main

import (
	"fmt"
	"github.com/hashicorp/nomad/api"
)

func main(){
	address := "http://localhost:4646"
	con := api.DefaultConfig()
	con.Address = address

	client, _ := api.NewClient(con)
	nodeCon := client.Nodes()
	node, _, _ := nodeCon.List(&api.QueryOptions{})
	fmt.Println(node[0].Name)
}