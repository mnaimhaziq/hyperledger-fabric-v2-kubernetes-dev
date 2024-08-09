package main

import (
	"encoding/json" // Add this import
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)


// KeyValueContract contract for handling writing and reading from the world state
type KeyValueContract struct {
	contractapi.Contract
}

type Asset struct {
	AppraisedValue int    `json:"AppraisedValue"`
	Color          string `json:"Color"`
	ID             string `json:"ID"`
	Owner          string `json:"Owner"`
	Size           int    `json:"Size"`
}

// Create adds a new key with value to the world state
func (sc *KeyValueContract) Create(ctx contractapi.TransactionContextInterface,  id string, color string, size int, owner string, appraisedValue int) error {
	existing, err := ctx.GetStub().GetState(id)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	if existing != nil {
		return fmt.Errorf("Cannot create world state pair with key %s. Already exists", id)
	}

	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}
	err = ctx.GetStub().PutState(id, assetJSON)

	return nil
}

// Update changes the value with key in the world state
func (sc *KeyValueContract) Update(ctx contractapi.TransactionContextInterface,  id string, color string, size int, owner string, appraisedValue int) error {
	existing, err := ctx.GetStub().GetState(id)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	if existing == nil {
		return fmt.Errorf("Cannot update world state pair with key %s. Does not exist", id)
	}

	asset := Asset{
		ID:             id,
		Color:          color,
		Size:           size,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	
	if err != nil {
		return errors.New("Unable to interact with world state")
	}
	err = ctx.GetStub().PutState(id, assetJSON)

	return nil
}

// Read returns the value at key in the world state
func (sc *KeyValueContract) Read(ctx contractapi.TransactionContextInterface, id string) (string, error) {
	assetJSON, err := ctx.GetStub().GetState(id)

	if err != nil {
		return "", errors.New("Unable to interact with world state")
	}

	if assetJSON == nil {
		return "", fmt.Errorf("Cannot read world state pair with key %s. Does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)

	if err != nil {
		return nil, err
	}
	return &asset, nil
}

// DeleteAsset deletes an given asset from the world state.
func (sc *SmartContract) Delete(ctx contractapi.TransactionContextInterface, id string) error {
	existing, err := ctx.GetStub().GetState(id)

	if err != nil {
		return errors.New("Unable to interact with world state")
	}

	if existing == nil {
		return fmt.Errorf("Cannot delete world state pair with key %s. Does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// GetAllAssets returns all assets found in world state
func (sc *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
