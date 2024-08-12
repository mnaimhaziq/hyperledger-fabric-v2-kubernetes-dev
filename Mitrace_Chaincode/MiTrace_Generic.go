package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// MiTrace_Generic_Chaincode implementation
type MiTrace_Generic_Chaincode struct {
	testMode bool
}

//
// Init implement chaincode initialization [PUBLIC]
// No need init since we cannot predefine a document
func (m *MiTrace_Generic_Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Init MiTrace Generic Chaincode Log")
	_, args := stub.GetFunctionAndParameters()

	if len(args) == 0 {
		fmt.Println("No arguments")
		return shim.Success(nil)
	}

	return shim.Success(nil)
}

// Invoke Functions for MiTrace_Generic
func (m *MiTrace_Generic_Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// Retrieve the requested Smart Contract function and arguments
	var creatorOrg, invoker string
	var creatorCertIssuer []string
	var err error

	fmt.Println("\n--- MiTrace Generic Invoke ---")

	if !m.testMode {
		creatorOrg, creatorCertIssuer, invoker, err = getTxCreatorInfo(stub)
		if err != nil {
			fmt.Errorf("Error extracting creator identity info: %s\n", err.Error())
			return shim.Error(err.Error())
		}
		fmt.Printf("MiTrace Generic Invoke by '%s' from MSP '%s' and OrgName '%s'\n", invoker, creatorOrg, creatorCertIssuer)
	}

	fa, fargs := stub.GetFunctionAndParameters()
	fmt.Printf("case %s : %s", fa, fargs)
	switch function, args := stub.GetFunctionAndParameters(); function {
	// Route to the appropriate handler function to interact with the ledger appropriately
	case "getSanity":
		fmt.Println("\nRunning getSanity func")
		return m.getSanity(stub, creatorOrg, creatorCertIssuer, args)
	case "addSanity":
		fmt.Println("\nRunning addSanity func")
		return m.addSanity(stub, creatorOrg, creatorCertIssuer, args)
	case "updateSanity":
		fmt.Println("\nRunning updateSanity func")
		return m.updateSanity(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "invalidateSanity":
		fmt.Println("\nRunning invalidateSanity func")
		return m.invalidateSanity(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "getAllSanity":
		fmt.Println("\nRunning getAllSanity func")
		return m.getAllSanity(stub, creatorOrg, creatorCertIssuer)
	case "addCompany":
		fmt.Println("\nRunning addCompany func")
		return m.addCompany(stub, creatorOrg, creatorCertIssuer, args)
	case "updateCompany":
		fmt.Println("\nRunning updateCompany func")
		return m.updateCompany(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "invalidateCompany":
		fmt.Println("\nRunning invalidateCompany func")
		return m.invalidateCompany(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "getCompany":
		fmt.Println("\nRunning getCompany func")
		return m.getCompany(stub, creatorOrg, creatorCertIssuer, args)
	case "getAllCompany":
		fmt.Println("\nRunning getAllCompany func")
		return m.getAllCompany(stub, creatorOrg, creatorCertIssuer)
	case "addPermit":
		fmt.Println("\nRunning addPermit func")
		return m.addPermit(stub, creatorOrg, creatorCertIssuer, args)
	case "updatePermit":
		fmt.Println("\nRunning updatePermit func")
		return m.updatePermit(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "invalidatePermit":
		fmt.Println("\nRunning invalidatePermit func")
		return m.invalidatePermit(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "getPermit":
		fmt.Println("\nRunning getPermit func")
		return m.getPermit(stub, creatorOrg, creatorCertIssuer, args)
	case "getAllPermit":
		fmt.Println("\nRunning getAllPermit func")
		return m.getAllPermit(stub, creatorOrg, creatorCertIssuer)
	case "addPrs":
		fmt.Println("\nRunning addPrs func")
		return m.addPrs(stub, creatorOrg, creatorCertIssuer, args)
	case "updatePrs":
		fmt.Println("\nRunning updatePrs func")
		return m.updatePrs(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "invalidatePrs":
		fmt.Println("\nRunning invalidatePrs func")
		return m.invalidatePrs(stub, creatorOrg, creatorCertIssuer, invoker, args)
	case "getPrs":
		fmt.Println("\nRunning getPrs func")
		return m.getPrs(stub, creatorOrg, creatorCertIssuer, args)
	case "getAllPrs":
		fmt.Println("\nRunning getAllPrs func")
		return m.getAllPrs(stub, creatorOrg, creatorCertIssuer)
	case "addSandec":
		fmt.Println("\nRunning addSandec func")
		return m.addSandec(stub, creatorOrg, creatorCertIssuer, args)
	case "getData":
		fmt.Println("\nRunning getData func")
		return m.getData(stub, creatorOrg, creatorCertIssuer, args)
	case "getHistory":
		fmt.Println("\nRunning getHistory func")
		return m.getHistory(stub, creatorOrg, creatorCertIssuer, args)
	default:
		return shim.Error("\nInvalid function name to invoke.")
	}
}

// Main function to call/start Chaincode and change to test mode
func main() {
	mitCC := new(MiTrace_Generic_Chaincode)
	mitCC.testMode = false
	err := shim.Start(mitCC)
	if err != nil {
		fmt.Printf("Error starting MiTrace Generic Chaincode: %s", err)
	}
}
