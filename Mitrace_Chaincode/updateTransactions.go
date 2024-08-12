package main

import (
	"encoding/json"
	"fmt"
	"time"
	// "strings"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"strconv" 
)

// ALL UPDATE DATA FUNCTIONS

func (m *MiTrace_Generic_Chaincode) updateCompany(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {
	// Input sanitation
	fmt.Println("--- Start Update Company Details ---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) && !authenticateAdminUser(creatorOrg, creator) {
		return shim.Error(authAllOrgAdmin)
	}

	// Check args length
	if len(args) <= 0 {
		fmt.Printf("Argument must not be empty")
		return shim.Error(argsErr + "1")
	}

	utc := time.Now().UTC()
	currentLocalTime := ""
	loc, errr := time.LoadLocation("Asia/Singapore")
	if errr != nil {
		return shim.Error(errr.Error())
	} else {
		currentLocalTime = utc.In(loc).Format(time.RFC1123)
	}

	aeoParam := args[0]
	aeoPassed := AEOCompany{} // create instance to store AEO data
	err := json.Unmarshal([]byte(aeoParam), &aeoPassed)
	if err != nil {
		return shim.Error(err.Error())
	}

	aeoAsBytes, err := stub.GetState(aeoPassed.AEOId)
	if err != nil {
		return shim.Error("Failed to get AEO : " + err.Error())
	} else if aeoAsBytes == nil {
		return shim.Error(errNotExistMSG)
	}

	aeo := AEOCompany{}
	err = json.Unmarshal(aeoAsBytes, &aeo)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("AEO COMPANY UNMARSHALLED")
	aeo.ObjectType = "COMPANY"
	
	// Track if any updates occur to determine if version should be incremented
	updatesMade := true

	// Update fields only if they are provided and different from current values
	if aeoPassed.CompanyName != "" && aeoPassed.CompanyName != aeo.CompanyName {
		aeo.CompanyName = aeoPassed.CompanyName
		// updatesMade = true
	}
	if aeoPassed.CompanyAddress != "" && aeoPassed.CompanyAddress != aeo.CompanyAddress {
		aeo.CompanyAddress = aeoPassed.CompanyAddress
		// updatesMade = true
	}
	if aeoPassed.CompanyTel != "" && aeoPassed.CompanyTel != aeo.CompanyTel {
		aeo.CompanyTel = aeoPassed.CompanyTel
		// updatesMade = true
	}
	if aeoPassed.CompanyFax != "" && aeoPassed.CompanyFax != aeo.CompanyFax {
		aeo.CompanyFax = aeoPassed.CompanyFax
		// updatesMade = true
	}
	if aeoPassed.AEOStatus != "" && aeoPassed.AEOStatus != aeo.AEOStatus {
		aeo.AEOStatus = aeoPassed.AEOStatus
		// updatesMade = true
	}
	if aeoPassed.AEOFileRefNo != "" && aeoPassed.AEOFileRefNo != aeo.AEOFileRefNo {
		aeo.AEOFileRefNo = aeoPassed.AEOFileRefNo
		// updatesMade = true
	}
	if aeoPassed.ModifiedBy != "" && aeoPassed.ModifiedBy != aeo.ModifiedBy {
		aeo.ModifiedBy = aeoPassed.ModifiedBy
		// updatesMade = true
	}
	if aeoPassed.Remarks != "" && aeoPassed.Remarks != aeo.Remarks {
		aeo.Remarks = aeoPassed.Remarks
		// updatesMade = true
	}
	if aeoPassed.Reserved != "" && aeoPassed.Reserved != aeo.Reserved {
		aeo.Reserved = aeoPassed.Reserved
		// updatesMade = true
	}
	if aeoPassed.EffectiveDateFrom != "" && aeoPassed.EffectiveDateFrom != aeo.EffectiveDateFrom {
		aeo.EffectiveDateFrom = aeoPassed.EffectiveDateFrom
		// updatesMade = true
	}
	if aeoPassed.EffectiveDateTo != "" && aeoPassed.EffectiveDateTo != aeo.EffectiveDateTo {
		aeo.EffectiveDateTo = aeoPassed.EffectiveDateTo
		// updatesMade = true
	}
	if aeoPassed.AEOCardIssued != "" && aeoPassed.AEOCardIssued != aeo.AEOCardIssued {
		aeo.AEOCardIssued = aeoPassed.AEOCardIssued
		// updatesMade = true
	}
	if aeoPassed.ForwarderCardIssued != "" && aeoPassed.ForwarderCardIssued != aeo.ForwarderCardIssued {
		aeo.ForwarderCardIssued = aeoPassed.ForwarderCardIssued
		// updatesMade = true
	}
	if updatesMade {
		// Increment version
		currentVersion, convErr := strconv.Atoi(aeo.Version)
		if convErr != nil {
			return shim.Error("Current version is invalid: " + convErr.Error())
		}
		aeo.Version = strconv.Itoa(currentVersion + 1)
	}

	aeo.ModifiedDate = currentLocalTime

	aeoJSONasBytes, err := json.Marshal(aeo)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("AEO COMPANY MARSHALLED!")

	// PutState
	err = stub.PutState(aeoPassed.AEOId, aeoJSONasBytes)
	if err != nil {
		return shim.Error("Failed to put AEO data: " + err.Error())
	}
	fmt.Println("AEO COMPANY PUTSTATE PASSED!")

	fmt.Println("--- End update company Successfully ---")

	return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) updatePermit(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {
	// Input sanitation
	fmt.Println("--- Start Update Permit Details ---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) && !authenticateAdminUser(creatorOrg, creator) {
		return shim.Error(authAllOrgAdmin)
	}

	// Check args length
	if len(args) <= 0 {
		fmt.Printf("Argument must not be empty")
		return shim.Error(argsErr + "1")
	}

	utc := time.Now().UTC()
	currentLocalTime := ""
	loc, errr := time.LoadLocation("Asia/Singapore")
	if errr != nil {
		return shim.Error(errr.Error())
	} else {
		currentLocalTime = utc.In(loc).Format(time.RFC1123)
	}

	pmtParam := args[0]
	pmtPassed := Permit{} // create instance to store permit data
	err := json.Unmarshal([]byte(pmtParam), &pmtPassed)
	if err != nil {
		return shim.Error(err.Error())
	}

	pmtAsBytes, err := stub.GetState(pmtPassed.PermitId)
	if err != nil {
		return shim.Error("Failed to get permit : " + err.Error())
	} else if pmtAsBytes == nil {
		return shim.Error(errNotExistMSG)
	}

	permit := Permit{}
	err = json.Unmarshal(pmtAsBytes, &permit)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("PERMIT UNMARSHALLED")
	permit.ObjectType = "PERMIT"
	
	// Track if any updates occur to determine if version should be incremented
	// updatesMade := false

	// Update fields only if they are provided and different from current values

	if pmtPassed.PermitClass != "" && pmtPassed.PermitClass != permit.PermitClass {
		permit.PermitClass = pmtPassed.PermitClass
		// updatesMade = true
	}
	if pmtPassed.ActivityType != "" && pmtPassed.ActivityType != permit.ActivityType {
		permit.ActivityType = pmtPassed.ActivityType
		// updatesMade = true
	}
	if pmtPassed.PermitType != "" && pmtPassed.PermitType != permit.PermitType {
		permit.PermitType = pmtPassed.PermitType
		// updatesMade = true
	}
	if pmtPassed.ICPId != "" && pmtPassed.ICPId != permit.ICPId {
		permit.ICPId = pmtPassed.ICPId
		// updatesMade = true
	}
	if pmtPassed.ICPExpiry != "" && pmtPassed.ICPExpiry != permit.ICPExpiry {
		permit.ICPExpiry = pmtPassed.ICPExpiry
		// updatesMade = true
	}
	if pmtPassed.PermitExpiry != "" && pmtPassed.PermitExpiry != permit.PermitExpiry {
		permit.PermitExpiry = pmtPassed.PermitExpiry
		// updatesMade = true
	}
	if pmtPassed.DestinationCountry != "" && pmtPassed.DestinationCountry != permit.DestinationCountry {
		permit.DestinationCountry = pmtPassed.DestinationCountry
		// updatesMade = true
	}
	if pmtPassed.Incoterms != "" && pmtPassed.Incoterms != permit.Incoterms {
		permit.Incoterms = pmtPassed.Incoterms
		// updatesMade = true
	}
	if pmtPassed.CompanyName != "" && pmtPassed.CompanyName != permit.CompanyName {
		permit.CompanyName = pmtPassed.CompanyName
		// updatesMade = true
	}
	if pmtPassed.ImporterName != "" && pmtPassed.ImporterName != permit.ImporterName {
		permit.ImporterName = pmtPassed.ImporterName
		// updatesMade = true
	}
	if pmtPassed.LogisticOperator != "" && pmtPassed.LogisticOperator != permit.LogisticOperator {
		permit.LogisticOperator = pmtPassed.LogisticOperator
		// updatesMade = true
	}
	if pmtPassed.SMKStatus != "" && pmtPassed.SMKStatus != permit.SMKStatus {
		permit.SMKStatus = pmtPassed.SMKStatus
		// updatesMade = true
	}
	if pmtPassed.CreatedDate != "" && pmtPassed.CreatedDate != permit.CreatedDate {
		permit.CreatedDate = pmtPassed.CreatedDate
		// updatesMade = true
	}
	if pmtPassed.CreatedBy != "" && pmtPassed.CreatedBy != permit.CreatedBy {
		permit.CreatedBy = pmtPassed.CreatedBy
		// updatesMade = true
	}
	if pmtPassed.ModifiedDate != "" && pmtPassed.ModifiedDate != permit.ModifiedDate {
		permit.ModifiedDate = pmtPassed.ModifiedDate
		// updatesMade = true
	}
	if pmtPassed.ModifiedBy != "" && pmtPassed.ModifiedBy != permit.ModifiedBy {
		permit.ModifiedBy = pmtPassed.ModifiedBy
		// updatesMade = true
	}
	if pmtPassed.Remarks != "" && pmtPassed.Remarks != permit.Remarks {
		permit.Remarks = pmtPassed.Remarks
		// updatesMade = true
	}
	permit.ModifiedDate = currentLocalTime

	// if updatesMade {
	// 	// Increment version
	// 	currentVersion, convErr := strconv.Atoi(permit.Version)
	// 	if convErr != nil {
	// 		return shim.Error("Current version is invalid: " + convErr.Error())
	// 	}
	// 	permit.Version = strconv.Itoa(currentVersion + 1)
	// }

	aeoJSONasBytes, err := json.Marshal(permit)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("PERMIT MARSHALLED!")

	// PutState
	err = stub.PutState(pmtPassed.PermitId, aeoJSONasBytes)
	if err != nil {
		return shim.Error("Failed to put permit data: " + err.Error())
	}
	fmt.Println("PERMIT PUTSTATE PASSED!")

	fmt.Println("--- End updatePermit Successfully ---")

	return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) updatePrs(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {
	// Input sanitation
	fmt.Println("--- Start Update PRS Details ---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) && !authenticateAdminUser(creatorOrg, creator) {
		return shim.Error(authAllOrgAdmin)
	}

	// Check args length
	if len(args) <= 0 {
		fmt.Printf("Argument must not be empty")
		return shim.Error(argsErr + "1")
	}

	utc := time.Now().UTC()
	currentLocalTime := ""
	loc, errr := time.LoadLocation("Asia/Singapore")
	if errr != nil {
		return shim.Error(errr.Error())
	} else {
		currentLocalTime = utc.In(loc).Format(time.RFC1123)
	}

	prsParam := args[0]
	prsPassed := PRS{} // create instance to store permit data
	err := json.Unmarshal([]byte(prsParam), &prsPassed)
	if err != nil {
		return shim.Error(err.Error())
	}

	prsAsBytes, err := stub.GetState(prsPassed.PRSId)
	if err != nil {
		return shim.Error("Failed to get prs : " + err.Error())
	} else if prsAsBytes == nil {
		return shim.Error(errNotExistMSG)
	}

	prs := PRS{}
	err = json.Unmarshal(prsAsBytes, &prs)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("PRS UNMARSHALLED")
	prs.ObjectType = "PRS"
	
	// Track if any updates occur to determine if version should be incremented
	// updatesMade := false

	// Update fields only if they are provided and different from current values

	if prsPassed.ModifiedBy != "" && prsPassed.ModifiedBy != prs.ModifiedBy {
		prs.ModifiedBy = prsPassed.ModifiedBy
		// updatesMade = true
	}
	if prsPassed.Remarks != "" && prsPassed.Remarks != prs.Remarks {
		prs.Remarks = prsPassed.Remarks
		// updatesMade = true
	}
	if prsPassed.SMKNo != "" && prsPassed.SMKNo != prs.SMKNo {
		prs.SMKNo = prsPassed.SMKNo
	}
	if prsPassed.SMKStatus != "" && prsPassed.SMKStatus != prs.SMKStatus {
		prs.SMKStatus = prsPassed.SMKStatus
	}
	prs.ModifiedDate = currentLocalTime

	// if updatesMade {
	// 	// Increment version
	// 	currentVersion, convErr := strconv.Atoi(permit.Version)
	// 	if convErr != nil {
	// 		return shim.Error("Current version is invalid: " + convErr.Error())
	// 	}
	// 	permit.Version = strconv.Itoa(currentVersion + 1)
	// }

	aeoJSONasBytes, err := json.Marshal(prs)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("PRS MARSHALLED!")

	// PutState
	err = stub.PutState(prsPassed.PRSId, aeoJSONasBytes)
	if err != nil {
		return shim.Error("Failed to put permit data: " + err.Error())
	}
	fmt.Println("PRS PUTSTATE PASSED!")

	fmt.Println("--- End updatePRS Successfully ---")

	return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) updateSanity(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, creator string, args []string) pb.Response {

	// Input sanitation
	fmt.Println("--- Start Update Permit Details ---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) && !authenticateAdminUser(creatorOrg, creator) {
		return shim.Error(authAllOrgAdmin)
	}

	// Check args length
	if len(args) <= 0 {
		fmt.Printf("Argument must not be empty")
		return shim.Error(argsErr + "1")
	}

	utc := time.Now().UTC()
	currentLocalTime := ""
	loc, errr := time.LoadLocation("Asia/Singapore")
	if errr != nil {
		return shim.Error(errr.Error())
	} else {
		currentLocalTime = utc.In(loc).Format(time.RFC1123)
	}

	lrParam := args[0]
	lrPassed := CheckSanity{} // create instance to store trade item
	err := json.Unmarshal([]byte(lrParam), &lrPassed)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("lrParam="+ lrParam + "\n")

	// comKey := getSanityIDKey("SANITY", lrPassed.CheckID)
	lrecordAsBytes, err := stub.GetState(lrPassed.CheckID)
	// fmt.Printf("KEY="+ comKey + "\n")
	if err != nil {
		return shim.Error("Failed to get Sanity data : " + err.Error())
	} else if lrecordAsBytes == nil {
		return shim.Error(errNotExistMSG)
		// return shim.Error(strings.Replace(errNotExistMSG, "!", "!\"\n\"key\":\"" + lrPassed.CheckID, 1))
	}

	CheckSan := CheckSanity{}
	err = json.Unmarshal(lrecordAsBytes, &CheckSan)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("SANITY UNMARSHALLED")
	CheckSan.ObjectType = "SANITY"
	if lrPassed.DateTime != "" {
		CheckSan.DateTime = lrPassed.DateTime
	}

	if lrPassed.TestValue != "**nil&**" {
		CheckSan.TestValue = lrPassed.TestValue
	}

	CheckSan.ModifyTime = currentLocalTime
	
	if lrPassed.ModifyBy != "" {
		CheckSan.ModifyBy = lrPassed.ModifyBy
	}

	if lrPassed.Remarks != "" {
		CheckSan.Remarks = lrPassed.Remarks
	}

	checkJSONAsBytes, err := json.Marshal(CheckSan)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("SANITY UPDATE MARSHALLED!")

	//Putstate
	err = stub.PutState(lrPassed.CheckID, checkJSONAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("SANITY UPDATE PUTSTATE PASSED!")

	fmt.Println("--- End updateSanity Successfully ---")

	return shim.Success(nil)
}
