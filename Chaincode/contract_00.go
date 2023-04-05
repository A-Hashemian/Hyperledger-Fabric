
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
	
	
	for _, asset := range assets {
		err := ctx.GetStub().PutState(asset.AssetID, []byte(fmt.Sprintf("%d", asset.Value)))
		if err != nil {
			return fmt.Errorf("failed to put asset %s on ledger: %v", asset.AssetID, err)
		}
	}

	return nil
}


// CreateAsset adds a new asset to the ledger
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, assetID string, value int) error {
	exists, err := s.AssetExists(ctx, assetID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", assetID)
	}

	asset := SimpleAsset{
		AssetID: assetID,
		Value:   value,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(assetID, assetJSON)
}


// ReadAsset returns the asset stored in the ledger
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, assetID string) (*SimpleAsset, error) {
	assetJSON, err := ctx.GetStub().GetState(assetID)
	if err != nil {
		return nil, fmt.Errorf("failed to read asset %s from ledger: %v", assetID, err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", assetID)
	}

	var asset SimpleAsset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// AssetExists returns true when asset with given ID exists in the ledger
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, assetID string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(assetID)
	if err != nil {
		return false, fmt.Errorf("failed to read asset %s from ledger: %v", assetID, err)
	}

	return assetJSON != nil, nil
}

// UpdateAsset updates an existing asset in the ledger
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, assetID string, newValue int) error {
	exists, err := s.AssetExists(ctx, assetID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", assetID)
	}

	asset := SimpleAsset{
		AssetID: assetID,
		Value:   newValue,
		
		
		}
assetJSON, err := json.Marshal(asset)
if err != nil {
	return err
}
	
return ctx.GetStub().PutState(assetID, assetJSON)
}
