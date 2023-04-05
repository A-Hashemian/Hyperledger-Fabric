
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


// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []SimpleAsset{
		{AssetID: "asset1", Value: 100},
		{AssetID: "asset2", Value: 200},
		{AssetID: "asset3", Value: 300},
		{AssetID: "asset4", Value: 400},
		{AssetID: "asset5", Value: 500},
	}
