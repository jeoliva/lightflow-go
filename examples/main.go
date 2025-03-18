package main

import (
	"fmt"

	"github.com/jeoliva/lightflow-go"
)

func main() {
	api_base_url := ""
	api_token := ""
	api := lightflow.NewAPIClient(api_base_url, api_token)

	// List assets
	println("Listing assets")
	assets, err := api.GetAssets()
	if err != nil {
		panic(err)
	}

	for _, asset := range assets.Items {
		fmt.Printf("UUID: %s, Url: %s\n", *asset.UUID, asset.Parameters.Input.UrlPath)
	}
}
