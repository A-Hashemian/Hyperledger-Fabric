
package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimpleAsset is a basic struct for the asset
type SimpleAsset struct {
	AssetID string `json:"assetID"`
	Value   int    `json:"value"`
}


// SmartContract provides functions for managing a SimpleAsset
type SmartContract struct {
	contractapi.Contract
}
