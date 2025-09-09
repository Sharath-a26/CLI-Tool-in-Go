package main

 import (
	"fmt"
	"log"
	"github.com/CLI-Tool-in-Go/internal/pokeapi"
 )
 func callBackMap() error {
	pokeapi_client := pokeapi.NewClient()

	resp, err := pokeapi_client.LocationAreas()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas")
	for _, v := range resp.Results {
		fmt.Printf("- %s\n", v.Name)
		
	}
	return nil
 }