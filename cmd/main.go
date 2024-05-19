package main

import (
	"encoding/json"
	"os"
	"fmt"

	"github.com/mallett002/go-fake-store/internal/factories"
)

// Store processor calls storeFactoy that builds all the:
	// stores, customers, items and orders
// Store processor exposes functions that
// getStoreWithHighestInventory
// getMostCommonlyPurchasedItem
// etc
// Tests test these functions

func main() {
	franchise := factories.CreateFranchise()

	franchiseJson, _ := json.Marshal(franchise)
    err := os.WriteFile("franchise.json", franchiseJson, 0644)

	if err != nil {
		fmt.Printf("Error: %v", err)
	}

    fmt.Println("Finished creating data! Find it in franchise.json")
}

func GetStoreWithHighestInventory() factories.Store {
	// This will read from the created franchise.json file to get its data
	return factories.Store{}
}