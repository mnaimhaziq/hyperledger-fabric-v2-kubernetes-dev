package main

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// ALL GET/QUERY FUNCTIONS

// get data for any id of any structs
// func (m *MiTrace_Generic_Chaincode) getcompany(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

// 	fmt.Println("--- Start Get Company ---")
	
// 	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
// 		// fmt.Println("Creator: ", creatorOrg)
// 		// fmt.Println("Issuer: ", creatorCertIssuer)
// 		return shim.Error(authAll)
// 	}

// 	if len(args) != 1 {
// 		return shim.Error(argsErr + "1")
// 	}
// 	if len(args[0]) <= 0 {
// 		return shim.Error("The argument value cannot be empty, expecting AEO ID")
// 	}

// 	aeoID := args[0]

// 	queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"COMPANY\",\"aeoId\":\"%s\"} , \"use_index\":[\"_design/indexIdentifierAEOIDDoc\" , \"indexIdentifierAEOID\"]}", aeoID)

// 	queryResults, err := getQueryResultForQueryString(stub, queryString)
// 	if err != nil {
// 		// return shim.Error(strings.Replace(errNotExistMSG, "!", "!\",\"key\":\"" + aeoID, 1))
// 		return shim.Error(errNotExistMSG)
// 	}
// 	return shim.Success(queryResults)
// }

func (m *MiTrace_Generic_Chaincode) getCompany(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
    fmt.Println("--- Start Get Company ---")
    
    if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
        return shim.Error(authAll)
    }

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1 argument")
    }

    id := args[0]
    if len(id) <= 0 {
        return shim.Error("The argument value cannot be empty, expecting an ID")
    }

    var idType string
    if strings.HasPrefix(id, "AEO") {
        idType = "aeoId"
    } else if strings.HasPrefix(id, "SSM") {
        idType = "ssmId"
    } else {
        return shim.Error("Invalid ID provided; ID must start with 'AEO' or 'SSM'")
    }

    queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"COMPANY\",\"%s\":\"%s\"}, \"use_index\":[\"_design/indexIdentifierDoc\",\"indexIdentifier\"]}", idType, id)

    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        return shim.Error(errNotExistMSG)
    }
    return shim.Success(queryResults)
}

// func (m *MiTrace_Generic_Chaincode) getcompany(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

// 	fmt.Println("--- Start Get Company by AEO and SSM ID ---")

// 	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
// 		// fmt.Println("Creator: ", creatorOrg)
// 		// fmt.Println("Issuer: ", creatorCertIssuer)
// 		return shim.Error(authAll)
// 	}

// 	if len(args) != 2 {
// 		return shim.Error("Incorrect number of arguments. Expecting 2 (AEO ID and SSM ID)")
// 	}

// 	aeoID := args[0]
// 	ssmID := args[1]

// 	if len(aeoID) <= 0 || len(ssmID) <= 0 {
// 		return shim.Error("Neither AEO ID nor SSM ID can be empty")
// 	}

// 	queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"COMPANY\",\"aeoId\":\"%s\",\"ssmId\":\"%s\"} , \"use_index\":[\"_design/indexAEOandSSMIDDoc\",\"indexAEOandSSMID\"]}", aeoID, ssmID)

// 	queryResults, err := getQueryResultForQueryString(stub, queryString)
// 	if err != nil {
// 		// return shim.Error(strings.Replace(errNotExistMSG, "!", "!\",\"key\":\"" + aeoID + "," + ssmID, 1))
// 		return shim.Error(errNotExistMSG)
// 	}
// 	return shim.Success(queryResults)
// }

// func (m *MiTrace_Generic_Chaincode) getAEO(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
// 	fmt.Println("--- Start Get AEO ---")
	
// 	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
// 		// fmt.Println("Creator: ", creatorOrg)
// 		// fmt.Println("Issuer: ", creatorCertIssuer)
// 		return shim.Error(authAll)
// 	}

// 	if len(args) != 2 {
// 		return shim.Error(argsErr + "2")
// 	}
	
// 	if len(args[0]) <= 2 {
// 		return shim.Error("The argument value cannot be empty, expecting AEO ID")
// 	}

// 	compId := args[0] // variables (aeoID or ssmID)
// 	value := args[1] // value (0 or 1)

// 	queryString := ""
// 	if (value == "0") {
// 		queryString = fmt.Sprintf("{\"selector\":{\"identifier\":\"AEO\",\"aeoId\":\"%s\"} , \"use_index\":[\"_design/indexIdentifierAEOIDDoc\" , \"indexIdentifierAEOID\"]}", compId)
// 	} else {
// 		queryString = fmt.Sprintf("{\"selector\":{\"identifier\":\"AEO\",\"ssmId\":\"%s\"} , \"use_index\":[\"_design/indexIdentifierSSMIDDoc\" , \"indexIdentifierSSMID\"]}", compId)
// 	}
// 	queryResults, err := getQueryResultForQueryString(stub, queryString)

// 	fmt.Println("RESULTS: ", queryResults, "LEN: ", len(queryResults))
// 	if err != nil {
// 		if (value == "0") {
// 			return shim.Error(strings.Replace(errAEOIDNotExistMSG, " !", " (" + compId + ") !", 1))
// 		} else {
// 			return shim.Error(strings.Replace(errSSMIDNotExistMSG, " !", " (" + compId + ") !", 1))
// 		}
// 	}
// 	if (len(queryResults) == 2) {
// 		if (value == "0") {
// 			return shim.Error(strings.Replace(errAEOIDNotExistMSG, " !", " (" + compId + ") !", 1))
// 		} else {
// 			return shim.Error(strings.Replace(errSSMIDNotExistMSG, " !", " (" + compId + ") !", 1))
// 		}	
// 	}
// 	return shim.Success(queryResults)
// }

func (m *MiTrace_Generic_Chaincode) getAllCompany(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string) pb.Response {

	fmt.Println("--- Start Get All Company ---")
	
	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authAll)
	}

	queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"COMPANY\"} , \"use_index\":[\"_design/indexIdentifierDoc\" , \"indexIdentifier\"]}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (m *MiTrace_Generic_Chaincode) getPermit(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
    fmt.Println("--- Start Get Permit ---")
    
    if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
        return shim.Error(authAll)
    }

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1 argument")
    }

    id := args[0]
    if len(id) <= 0 {
        return shim.Error("The argument value cannot be empty, expecting an ID")
    }

    var idType string
    if strings.HasPrefix(id, "PMT") {
        idType = "permitId"
    } else if strings.HasPrefix(id, "AEO") {
        idType = "aeoId"
    } else {
        return shim.Error("Invalid ID provided; ID must start with 'PMT' or 'AEO'")
    }

    queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"PERMIT\",\"%s\":\"%s\"}, \"use_index\":[\"_design/indexIdentifierDoc\",\"indexIdentifier\"]}", idType, id)

    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        return shim.Error(errNotExistMSG)
    }
    return shim.Success(queryResults)
}

func (m *MiTrace_Generic_Chaincode) getAllPermit(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string) pb.Response {

	fmt.Println("--- Start Get All Permit ---")
	
	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authAll)
	}

	queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"PERMIT\"} , \"use_index\":[\"_design/indexIdentifierDoc\" , \"indexIdentifier\"]}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (m *MiTrace_Generic_Chaincode) getPrs(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
    fmt.Println("--- Start Get PRS ---")
    
    if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
        return shim.Error(authAll)
    }

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1 argument")
    }

    id := args[0]
    if len(id) <= 0 {
        return shim.Error("The argument value cannot be empty, expecting an ID")
    }

    var idType string
    if strings.HasPrefix(id, "PRS") {
        idType = "prsId"
    } else if strings.HasPrefix(id, "PMT") {
        idType = "permitId"
    } else {
        return shim.Error("Invalid ID provided; ID must start with 'PRS' or 'PMT'")
    }

    queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"PRS\",\"%s\":\"%s\"}, \"use_index\":[\"_design/indexIdentifierDoc\",\"indexIdentifier\"]}", idType, id)

    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        return shim.Error(errNotExistMSG)
    }
    return shim.Success(queryResults)
}

func (m *MiTrace_Generic_Chaincode) getAllPrs(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string) pb.Response {

	fmt.Println("--- Start Get All PRS ---")
	
	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authAll)
	}

	queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"PRS\"} , \"use_index\":[\"_design/indexIdentifierDoc\" , \"indexIdentifier\"]}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (m *MiTrace_Generic_Chaincode) getSanity(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authAll)
	}

	if len(args) != 1 {
		return shim.Error(argsErr + "1")
	}
	if len(args[0]) <= 0 {
		return shim.Error("The argument value cannot be empty, expecting Check ID")
	}

	checkID := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"SANITY\",\"checkId\":\"%s\"} , \"use_index\":[\"_design/indexIdentifierCheckIDDoc\" , \"indexIdentifierCheckID\"]}", checkID)

	// fmt.Println("getSanity queryString : [%s]", queryString)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		// return shim.Error(strings.Replace(errNotExistMSG, "!", "!\"\n\"key\":\"" + checkID, 1))
		return shim.Error(errNotExistMSG)
	}
	return shim.Success(queryResults)
}

func (m *MiTrace_Generic_Chaincode) getAllSanity(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string) pb.Response {

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authAll)
	}

	queryString := fmt.Sprintf("{\"selector\":{\"identifier\":\"SANITY\"} , \"use_index\":[\"_design/indexIdentifierDoc\" , \"indexIdentifier\"]}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

//try
