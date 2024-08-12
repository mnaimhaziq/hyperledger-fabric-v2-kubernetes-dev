package main

import (
	"encoding/json"
	"fmt"
	// "strconv"
	// "time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (m *MiTrace_Generic_Chaincode) invalidateCompany(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {

	// Input sanitation
	fmt.Println("--- Start Invalidate Company ---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error(authAll)
	}

	// Check args length
	if len(args) <= 0 {
		fmt.Printf("Argument must not be empty")
		return shim.Error(argsErr + "1")
	}

	aeoParam := args[0]
	aeoPassed := AEOCompany{} // create instance to store trade item
	err := json.Unmarshal([]byte(aeoParam), &aeoPassed)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.DelState(aeoPassed.AEOId)
	if err != nil {
		return shim.Error("Failed to invalidate company: " + err.Error())
	}
	fmt.Printf("\n--- End invalidatecompany:%s ---\n", aeoPassed.AEOId)
	return shim.Success(nil);
}

func (m *MiTrace_Generic_Chaincode) invalidatePermit(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {

	// Input sanitation
	fmt.Println("--- Start Invalidate Permit ---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error(authAll)
	}

	// Check args length
	if len(args) <= 0 {
		fmt.Printf("Argument must not be empty")
		return shim.Error(argsErr + "1")
	}

	pmtParam := args[0]
	pmtPassed := Permit{} // create instance to store trade item
	err := json.Unmarshal([]byte(pmtParam), &pmtPassed)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.DelState(pmtPassed.PermitId)
	if err != nil {
		return shim.Error("Failed to invalidate permit: " + err.Error())
	}
	fmt.Printf("\n--- End invalidatepermit:%s ---\n", pmtPassed.PermitId)
	return shim.Success(nil);
}

func (m *MiTrace_Generic_Chaincode) invalidatePrs(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {
    // Input sanitation
    fmt.Println("--- Start Invalidate PRS ---")

    if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
        return shim.Error("Authentication failed: " + authAll)
    }

    // Check args length
    if len(args) <= 0 {
        return shim.Error("Argument must not be empty: " + argsErr + "1")
    }

    prsParam := args[0]
    prsPassed := PRS{} // create instance to store trade item
    err := json.Unmarshal([]byte(prsParam), &prsPassed)
    if err != nil {
        return shim.Error("Error unmarshalling PRS: " + err.Error())
    }

    // Delete the PRS
    err = stub.DelState(prsPassed.PRSId)
    if err != nil {
        return shim.Error("Failed to invalidate PRS: " + err.Error())
    }

    fmt.Printf("--- End invalidateprs: %s ---\n", prsPassed.PRSId)
    return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) invalidateSanity(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {

	// Input sanitation
	fmt.Println("--- Start Invalidate Sanity ---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error(authAll)
	}

	// Check args length
	if len(args) <= 0 {
		fmt.Printf("Argument must not be empty")
		return shim.Error(argsErr + "1")
	}

	sanityParam := args[0]
	sanityPassed := CheckSanity{} // create instance to store trade item
	err := json.Unmarshal([]byte(sanityParam), &sanityPassed)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.DelState(sanityPassed.CheckID)
	if err != nil {
		return shim.Error("Failed to invalidate sanity data: " + err.Error())
	}
	fmt.Printf("\n--- End invalidateSanity:%s ---\n", sanityPassed.CheckID)
	return shim.Success(nil);
}
