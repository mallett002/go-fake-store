package main

import (
	"testing"

	"github.com/mallett002/go-fake-store/internal/factories"
	"github.com/stretchr/testify/assert"
)



func TestGetStoreWithHighestInventory(t *testing.T) {
	store := GetStoreWithHighestInventory()
	address := factories.Address{
		City: "salix",
		State: "IA",
		Country: "United States",
		Street: "1620-280th st",
	}

	notExpectedStore := factories.Store{Id: "1", StoreName: "2", Address: address}

	assert.NotEqual(t, notExpectedStore, store)
}