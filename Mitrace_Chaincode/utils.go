package main

// IMPORTANT UTILITIES! DO NOT DELETE!
import (
	"bytes"
	"crypto/x509"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// ERROR CODES
// const errReq 		= "\"ERR1000 - Required!\"" // Error code for required (Not used in chaincode. Call in Middleware Validation Model)
const errExistMSG 		= "ERR2000 - Exist!" // Error code for exist
const errNotExistMSG 	= "ERR3000 - Does not exist!" // Error code for does not exist

// ERROR JSON MESSAGE
// const errReqMSG			= "{\"error\":" + 	errReq 		+	"}"
// const errExistMSG		= errExist
// const errNotExistMSG 	= errNotExist

// To get the certificate details
func getTxCreatorInfo(stub shim.ChaincodeStubInterface) (string, []string, string, error) {
	var mspid string
	var err error
	var empt []string
	var cert *x509.Certificate
	var orgs []string

	mspid, err = cid.GetMSPID(stub)
	if err != nil {
		fmt.Printf("Error getting MSP identity: %s\n", err.Error())
		return "", empt, "", err
	}

	cert, err = cid.GetX509Certificate(stub)
	if err != nil {
		fmt.Printf("Error getting client certificate: %s\n", err.Error())
		return "", empt, "", err
	}

	if len(cert.Subject.Organization) > 0 {
		orgs = cert.Subject.Organization
	} else if len(cert.Issuer.Organization) > 0 {
		orgs = cert.Issuer.Organization
	}

	// return mspid, cert.GetAttributeValue("o"), nil
	fmt.Println("MSP ID: '%s' | Organization: '%s'", mspid, cert.Issuer.Organization)
	return mspid, orgs, cert.Subject.CommonName, nil
}

// ORG ACCESS CONTROL
// Change according to org MSP ID and Organization Name
func authenticateOrg1(mspID string, certON []string) bool {
	org1Name := []string{"jkdm.gov.my", "moh1.moh.gov.my","Jkdm.gov.my"}
	org1MSP := []string{"jkdmmsp", "moh1msp","JkdmMSP"}
	var o1, o2, fin bool
	for _, oN := range org1Name {
		for _, orgName := range certON {
			// fmt.Println(oN , orgName)
			orgNameLower := strings.ToLower(orgName)
			if oN == orgNameLower {
				o1 = true
			}
		}
	}

	for _, oMSP := range org1MSP {
		// fmt.Println(oMSP , mspID)
		mspIDLower := strings.ToLower(mspID)
		if oMSP == mspIDLower {
			o2 = true
		}
	}

	if o1 && o2 {
		fin = true
	}
	return fin
}

func authenticateOrg2(mspID string, certON []string) bool {
	org2Name := []string{"miti.gov.my", "moh2.moh.gov.my"}
	org2MSP := []string{"mitimsp", "moh2msp","MitiMSP"}
	var o1, o2, fin bool
	for _, oN := range org2Name {
		for _, orgName := range certON {
			// fmt.Println(oN , orgName)
			orgNameLower := strings.ToLower(orgName)
			if oN == orgNameLower {
				o1 = true
			}
		}
	}

	for _, oMSP := range org2MSP {
		// fmt.Println(oMSP , mspID)
		mspIDLower := strings.ToLower(mspID)
		if oMSP == mspIDLower {
			o2 = true
		}
	}

	if o1 && o2 {
		fin = true
	}
	return fin
	// return (mspID == "Org2MSP") && (certCN[0] == "org2.gov.my")
}

func authenticateAllOrg(mspID string, certON []string) bool {
	allOrgName := []string{"jkdm.gov.my", "miti.gov.my", "aeo.gov.my", "oga.gov.my", "Jkdm.gov.my", "Miti.gov.my", "Aeo.gov.my", "Oga.gov.my"}
	allOrgMSP := []string{"jkdmmsp", "mitimsp", "aeomsp", "ogamsp", "JkdmMSP", "MitiMSP", "AeoMSP", "OgaMSP"}
	var o1, o2, fin bool
	for _, oN := range allOrgName {
		for _, orgName := range certON {
			fmt.Println(oN , orgName)
			orgNameLower := strings.ToLower(orgName)
			if oN == orgNameLower {
				o1 = true
			}
		}
	}

	for _, oMSP := range allOrgMSP {
		fmt.Println(oMSP , mspID)
		mspIDLower := strings.ToLower(mspID)
		if oMSP == mspIDLower {
			o2 = true
		}
	}
	
		// Print values of o1 and o2
	fmt.Printf("o1: %v\n", o1)
	fmt.Printf("o2: %v\n", o2)

	if o1 && o2 {
		fin = true
	}
	return fin
}

func authenticateAdminUser(mspID string, originator string) bool {
	originators := []string{"admin@jkdm.gov.my", "admin@miti.gov.my","admin@aeo.gov.my", "admin@oga.gov.my", "admin@moh1.moh.gov.my"}
	org1MSP := []string{"jkdmmsp", "mitimsp", "aeomsp", "ogamsp", "moh1msp"}
	var o1, o2, fin bool
	for _, oN := range originators {
		origin := strings.ToLower(originator)
		if oN == origin {
			o1 = true
		}
	}

	for _, oMSP := range org1MSP {
		// fmt.Println(oMSP , mspID)
		mspIDLower := strings.ToLower(mspID)
		if oMSP == mspIDLower {
			o2 = true
		}
	}

	if o1 && o2 {
		fin = true
	}
	return fin
}

// END ACCESS CONTROL

// func getEPCKey(stub shim.ChaincodeStubInterface, identity string, epcID string) (string, error) {
// 	epcKey, err := stub.CreateCompositeKey("identity~epcID", []string{identity, epcID})
// 	if err != nil {
// 		return errMsg, err
// 	}
// 	return epcKey, nil
// }

// func getTIKey(stub shim.ChaincodeStubInterface, identity string, itemId string) (string, error) {
// 	tiKey, err := stub.CreateCompositeKey("identity~itemId", []string{identity, itemId})
// 	if err != nil {
// 		return errMsg, err
// 	}
// 	return tiKey, nil
// }

// func getEventKey(stub shim.ChaincodeStubInterface, identity string, eventId string) (string, error) {
// 	eventKey, err := stub.CreateCompositeKey("identity~eventId", []string{identity, eventId})
// 	if err != nil {
// 		return errMsg, err
// 	}
// 	return eventKey, nil
// }

// func getBizLocKey(stub shim.ChaincodeStubInterface, identity string, blID string) (string, error) {
// 	bzKey, err := stub.CreateCompositeKey("identity~bizlocId", []string{identity, blID})
// 	if err != nil {
// 		return errMsg, err
// 	}
// 	return bzKey, nil
// }

func getcompanyIDKey(identity string, AEOID string) (string) {
	AEOKey := identity + "~" + AEOID

	hash := sha256.Sum256([]byte(AEOKey))
	return (hex.EncodeToString(hash[:]))
}

func getpermitIDKey(identity string, PMTID string) (string) {
	permitKey := identity + "~" + PMTID

	hash := sha256.Sum256([]byte(permitKey))
	return (hex.EncodeToString(hash[:]))
}

func getSanityIDKey(identity string, SANITYID string) (string) {
	sanityKey := identity + "~" + SANITYID

	hash := sha256.Sum256([]byte(sanityKey))
	return (hex.EncodeToString(hash[:]))
}

func unmarshalcompany(aeoID string) (AEOCompany, error) {
	param := "{\"aeoId\":\"" + aeoID + "\",\"identifier\":\"AEO\"}"
	passed := AEOCompany{}
	err := json.Unmarshal([]byte(param), &passed)
	if err != nil {
		return passed, err
	}
	return passed, nil
}

func unmarshalpermit(pmtID string) (Permit, error) {
	param := "{\"permitId\":\"" + pmtID + "\",\"identifier\":\"PERMIT\"}"
	passed := Permit{}
	err := json.Unmarshal([]byte(param), &passed)
	if err != nil {
		return passed, err
	}
	return passed, nil
}

func unmarshalSanity(sanityID string) (CheckSanity, error) {
	param := "{\"checkId\":\"" + sanityID + "\",\"identifier\":\"SANITY\"}"
	passed := CheckSanity{}
	err := json.Unmarshal([]byte(param), &passed)
	if err != nil {
		return passed, err
	}
	return passed, nil
}

func isJSON(str string) bool {
	var obj map[string]interface{}

	if json.Unmarshal([]byte(str) , &obj) == nil {
		return true;
	}

	return false;
}

func (m *MiTrace_Generic_Chaincode) getData(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

	if !m.testMode && !(authenticateAllOrg(creatorOrg, creatorCertIssuer)) {
		return shim.Error("Caller from neither Org! Access Denied !")
	}

	if len(args) != 1 {
		return shim.Error(argsErr + "1")
	}
	if len(args[0]) <= 0 {
		return shim.Error("The argument value cannot be empty, expecting any ID")
	}	
	
	dataId := args[0]
	dataAsBytes, err := stub.GetState(dataId)
	if err != nil {
		return shim.Error("Failed to get state for " + dataId)
	}
	if len(dataAsBytes) == 0 {
		return shim.Error("No record found for " + dataId)
	}
	fmt.Printf("Query Response: %s\n", string(dataAsBytes))
	return shim.Success(dataAsBytes)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return nil, err
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (m *MiTrace_Generic_Chaincode) getDocumentBySelector(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	// Input sanitation
	// fmt.Println("- start get ticket by selector -")
	// if len(args[0]) <= 0 {
	// 	return shim.Error("1st argument must be a non-empty string")
	// }
	// if len(args[1]) <= 0 {
	// 	return shim.Error("2nd argument must be a non-empty string")
	// }

	selector := args[0]
	key := args[1]

	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"ticket\",\"%s\":\"%s\"}}", selector, key)
	queryString := fmt.Sprintf("{\"selector\":{\"%s\":\"%s\"}}", selector, key)
	fmt.Println("queryString: ", queryString)
	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	buffer, err := constructQueryResponseFromIterator(resultsIterator)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Printf("--- getDocumentBySelector queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

func constructQueryResponseFromIterator(resultsIterator shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		fmt.Printf("next line")
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))

		// buffer.WriteString(", \"TxId\":")
		// buffer.WriteString("\"")
		// buffer.WriteString(string(queryResponse.TxId))
		// buffer.WriteString("\"")
		buffer.WriteString("}")

		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	return &buffer, nil
}

func constructQueryResponseFromIteratorCompositeKey(resultsIterator shim.StateQueryIteratorInterface, stub shim.ChaincodeStubInterface) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		fmt.Printf("next line")

		_, compositeKeyParts, err := stub.SplitCompositeKey(queryResponse.Key)
		if err != nil {
			fmt.Println("Error spliting Composite Key")
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"status\":")
		buffer.WriteString("\"")
		buffer.WriteString(compositeKeyParts[0])
		buffer.WriteString("\"")

		buffer.WriteString(", \"documentId\":")
		buffer.WriteString(string(compositeKeyParts[1]))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")
	return &buffer, nil
}

// func (m *MiTrace_Generic_Chaincode) getHistory(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

// 	// if !m.testMode && !(authenticateOrg2(creatorOrg, creatorCertIssuer) || authenticateOrg1(creatorOrg, creatorCertIssuer)) {
// 	// 	return shim.Error("Caller from neither Org! Access Denied !")
// 	// }
// 	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
// 		// fmt.Println("Creator: ", creatorOrg)
// 		// fmt.Println("Issuer: ", creatorCertIssuer)
// 		return shim.Error(authAll)
// 	}

// 	if len(args) != 1 {
// 		return shim.Error(argsErr + "1")
// 	}

// 	if len(args[0]) <= 0 {
// 		return shim.Error("ID Required !")
// 	}

// 	uID := args[0]

// 	fmt.Printf("- start getHistory: %s\n", uID)

// 	resultsIterator, err := stub.GetHistoryForKey(uID)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	defer resultsIterator.Close()

// 	// buffer is a JSON array containing historic values for the marble
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	bArrayMemberAlreadyWritten := false
// 	for resultsIterator.HasNext() {
// 		response, err := resultsIterator.Next()
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		// Add a comma before array members, suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"TxId\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(response.TxId)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Value\":")
// 		// if it was a delete operation on given key, then we need to set the
// 		//corresponding value null. Else, we will write the response.Value
// 		//as-is (as the Value itself a JSON marble)
// 		if response.IsDelete {
// 			buffer.WriteString("null")
// 		} else {
// 			buffer.WriteString(string(response.Value))
// 		}

// 		buffer.WriteString(", \"Timestamp\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"IsDelete\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(strconv.FormatBool(response.IsDelete))
// 		buffer.WriteString("\"")

// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 	}
// 	buffer.WriteString("]")

// 	fmt.Printf(buffer.String())

// 	return shim.Success(buffer.Bytes())
// }

func (m *MiTrace_Generic_Chaincode) getHistory(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

	// if !m.testMode && !(authenticateOrg2(creatorOrg, creatorCertIssuer) || authenticateOrg1(creatorOrg, creatorCertIssuer)) {
	// 	return shim.Error("Caller from neither Org! Access Denied !")
	// }
	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authAll)
	}

	if len(args) != 1 {
		return shim.Error(argsErr + "1")
	}

	if len(args[0]) <= 0 {
		return shim.Error("ID Required !")
	}

	uID := args[0]

	fmt.Printf("- start getHistory: %s\n", uID)

	resultsIterator, err := stub.GetHistoryForKey(uID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON marble)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf(buffer.String())

	return shim.Success(buffer.Bytes())
}


const authOrg1 = "Caller not a member of Org1. Access denied !"
const authAllOrgAdmin = "Caller not an Admin of any Orgs. Access denied !"
const authOrg2 = "Caller not a member Org2. Access denied !"
const authAll = "Caller from neither Org. Access denied !"
const argsErr = "Incorrect number of arguements. Expecting "
const compKey = "CREATED COMPOSITE KEY!"
const compKeySuccess = "COMPOSITE KEY PUTSTATE SUCCESS!"
const errMsg = "Error: "
