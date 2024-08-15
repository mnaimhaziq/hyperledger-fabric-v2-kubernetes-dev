package main

import (
    "encoding/json"
    "fmt"
    "time"
    "strconv"

    "github.com/hyperledger/fabric-chaincode-go/shim"
    pb "github.com/hyperledger/fabric-protos-go/peer"
)

func (m *MiTrace_Generic_Chaincode) addCompany(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    fmt.Println("--- Start Add Company ---")

    // if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
    //     return shim.Error("Authentication failed")
    // }

    if len(args) == 0 {
        fmt.Println("Argument must not be empty")
        return shim.Error("Arguments must not be empty")
    }

    utc := time.Now().UTC()
    loc, err := time.LoadLocation("Asia/Singapore")
    if err != nil {
        return shim.Error("Failed to load location: " + err.Error())
    }
    currentLocalTime := utc.In(loc).Format(time.RFC1123)

    aeoParam := args[0]
    aeoPassed := AEOCompany{}
    err = json.Unmarshal([]byte(aeoParam), &aeoPassed)

    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("AEO COMPANY UNMARSHALLED 1")
 
    aeoAsBytes, err := stub.GetState(aeoPassed.AEOId)
    if err != nil {
        return shim.Error("Failed to get AEO: " + err.Error())
    } else if aeoAsBytes == nil {

    aeo := AEOCompany{
        ObjectType:          "COMPANY",
        AEOId:               aeoPassed.AEOId,
        SSMId:               aeoPassed.SSMId,
        CompanyName:         aeoPassed.CompanyName,
        CompanyAddress:      aeoPassed.CompanyAddress,
        CompanyTel:          aeoPassed.CompanyTel,
        CompanyFax:          aeoPassed.CompanyFax,
        AEOFileRefNo:        aeoPassed.AEOFileRefNo,
        AEOStatus:           aeoPassed.AEOStatus,
        CreatedBy:           aeoPassed.CreatedBy,
        CreatedDate:         currentLocalTime,
        ModifiedBy:          aeoPassed.ModifiedBy,
        ModifiedDate:        currentLocalTime,
        Remarks:             aeoPassed.Remarks,
        Version:             aeoPassed.Version,
        Reserved:            aeoPassed.Reserved,
        EffectiveDateFrom:   aeoPassed.EffectiveDateFrom,
        EffectiveDateTo:     aeoPassed.EffectiveDateTo,
        AEOCardIssued:       aeoPassed.AEOCardIssued,
        ForwarderCardIssued: aeoPassed.ForwarderCardIssued,
    }

    fmt.Println("AEO COMPANY UNMARSHALLED 2")
 
        aeoJSONasBytes, err := json.Marshal(aeo)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("AEO COMPANY MARSHALLED!")
 
        //Putstate
        err = stub.PutState(aeoPassed.AEOId, aeoJSONasBytes)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("AEOCOMPANY PUTSTATE PASSED!")
 
        fmt.Println("--- End addAEO Successfully ---")
    } else {
        // return shim.Error(strings.Replace(errExistMSG, "!", "!\",\"key\":\"" + aeoPassed.AEOId, 1))
        return shim.Error(errExistMSG)
    }
 
    return shim.Success(nil)
}

//original
// func (m *MiTrace_Generic_Chaincode) addPermit(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

// 	// Input sanitation
// 	fmt.Println("--- Start Add Permit ---")

// 	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
// 		// fmt.Println("Creator: ", creatorOrg)
// 		// fmt.Println("Issuer: ", creatorCertIssuer)
// 		return shim.Error(authAll)
// 	}

// 	// Check args length
// 	if len(args) <= 0 {
// 		fmt.Printf("Argument must not be empty")
// 		return shim.Error(argsErr + "1")
// 	}

// 	// permitID := args[0]
// 	// aeoID := args[1]
// 	// ssmID := args[2]
// 	// permitClass := args[3]
// 	// activityType := args[4]
// 	// permitType := args[5]
// 	// icpID := args[6]
// 	// icpExpiryDate := args[7]
// 	// permitExpiryDate := args[8]
// 	// finalDest := args[9]
// 	// incoTerms := args[10]
// 	// consigNor := args[11]
// 	// consigNee := args[12]
// 	// logisticOperator := args[13]
// 	// itemDetails := args[14]
// 	// smkStatus := args[15]
// 	// permitCreatedBy := args[16]

// 	// var itemAttr []Item
// 	// if itemDetails != "" {
// 	// 	err2 := json.Unmarshal([]byte(itemDetails), &itemAttr)
// 	// 	if err2 != nil {
// 	// 		return shim.Error(err2.Error())
// 	// 	}
// 	// }

// 	utc := time.Now().UTC()
// 	currentLocalTime := ""
// 	loc, errr := time.LoadLocation("Asia/Singapore")
// 	if errr != nil {
// 		return shim.Error(errr.Error())
// 	} else {
// 		currentLocalTime = utc.In(loc).Format(time.RFC1123)
// 	}

// 	pmtParam := args[0]
// 	pmtPassed := Permit{}
// 	err := json.Unmarshal([]byte(pmtParam), &pmtPassed)
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}
// 	fmt.Println("PERMIT UNMARSHALLED 1")

//     // Check if AEOId exists
//     fmt.Println("Checking AEO ID existence for:", pmtPassed.AEOId)
//     aeoAsBytes, err := stub.GetState(pmtPassed.AEOId)
//     if err != nil {
//         return shim.Error("Failed to get AEO ID: " + err.Error())
//     } else if aeoAsBytes == nil {
//         return shim.Error("AEO ID does not exist")
//     }
//     fmt.Println("AEO ID exists")

//     // Unmarshal AEOCompany to check SSMId
//     aeoCompany := AEOCompany{}
//     err = json.Unmarshal(aeoAsBytes, &aeoCompany)
//     if err != nil {
//         return shim.Error("Failed to unmarshal AEOCompany: " + err.Error())
//     } 

//     // Check if SSMId matches
//     if aeoCompany.SSMId != pmtPassed.SSMId {
//         return shim.Error("SSM ID does not match with the one stored in AEOCompany")
//     }
//     fmt.Println("SSM ID matches")

// 	pmtAsBytes, err := stub.GetState(pmtPassed.PermitId)
// 	if err != nil {
// 		return shim.Error("Failed to get permit : " + err.Error())
// 	} else if pmtAsBytes == nil {
// 		// Permit doesn't exist, creating new
// 		// pmt := Permit{
// 		// 	ObjectType			: "PERMIT",
// 		// 	PermitId			: permitID,
// 		// 	AEOId				: aeoID,
// 		// 	SSMId				: ssmID,
// 		// 	PermitClass			: permitClass,
// 		// 	ActivityType		: activityType,
// 		// 	PermitType			: permitType,
// 		// 	ICPId				: icpID,
// 		// 	ICPExpiry			: icpExpiryDate,
// 		// 	PermitExpiry		: permitExpiryDate,
// 		// 	DestinationCountry	: finalDest,
// 		// 	Incoterms			: incoTerms,
// 		// 	AEOName				: consigNor,
// 		// 	ImporterName		: consigNee,
// 		// 	LogisticOperator	: logisticOperator,
// 		// 	ItemArray			: itemAttr,
// 		// 	SMKStatus			: smkStatus,
// 		// 	CreatedBy			: permitCreatedBy,	
// 		// 	CreatedDate			: currentLocalTime,		
// 		// }

// 		permit := Permit{}
// 		permit.ObjectType			= "PERMIT"			
// 		permit.PermitId				= pmtPassed.PermitId				
// 		permit.AEOId				= pmtPassed.AEOId				
// 		permit.SSMId				= pmtPassed.SSMId				
// 		permit.PermitClass			= pmtPassed.PermitClass			
// 		permit.ActivityType			= pmtPassed.ActivityType			
// 		permit.PermitType			= pmtPassed.PermitType			
// 		permit.ICPId				= pmtPassed.ICPId				
// 		permit.ICPExpiry			= pmtPassed.ICPExpiry			
// 		permit.PermitExpiry			= pmtPassed.PermitExpiry			
// 		// permit.PermitStatus			= pmtPassed.PermitStatus			
// 		permit.DestinationCountry	= pmtPassed.DestinationCountry	
// 		permit.Incoterms			= pmtPassed.Incoterms			
// 		permit.CompanyName			= pmtPassed.CompanyName				
// 		permit.ImporterName			= pmtPassed.ImporterName			
// 		permit.LogisticOperator		= pmtPassed.LogisticOperator		
// 		permit.ItemArray			= pmtPassed.ItemArray			
// 		permit.SMKStatus			= pmtPassed.SMKStatus			
// 		permit.CreatedBy			= pmtPassed.CreatedBy			
// 		permit.CreatedDate			= currentLocalTime	
// 		permit.Remarks              = pmtPassed.Remarks
			

// 		fmt.Println("PERMIT UNMARSHALLED 2")

// 		pmtJSONasBytes, err := json.Marshal(permit)
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		fmt.Println("PERMIT MARSHALLED!")

// 		//Putstate
// 		err = stub.PutState(pmtPassed.PermitId, pmtJSONasBytes)
// 		if err != nil {
// 			return shim.Error(err.Error())
// 		}
// 		fmt.Println("PERMIT PUTSTATE PASSED!")

// 		fmt.Println("--- End addpermit Successfully ---")
// 	} else {
// 		// return shim.Error(strings.Replace(errExistMSG, "!", "!\"\n\"key\":\"" + pmtPassed.PermitId, 1))
// 		return shim.Error(errExistMSG)
// 	}
// 	return shim.Success(nil)
// }

func (m *MiTrace_Generic_Chaincode) addPermit(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

    // Input sanitation
    fmt.Println("--- Start Add Permit ---")

    if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
        return shim.Error(authAll)
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
    pmtPassed := Permit{}
    err := json.Unmarshal([]byte(pmtParam), &pmtPassed)
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("PERMIT UNMARSHALLED 1")

    // Check if AEOId exists
    fmt.Println("Checking AEO ID existence for:", pmtPassed.AEOId)
    aeoAsBytes, err := stub.GetState(pmtPassed.AEOId)
    if err != nil {
        return shim.Error("Failed to get AEO ID: " + err.Error())
    } else if aeoAsBytes == nil {
        return shim.Error("AEO ID does not exist")
    }
    fmt.Println("AEO ID exists")

    // Unmarshal AEOCompany to check SSMId
    aeoCompany := AEOCompany{}
    err = json.Unmarshal(aeoAsBytes, &aeoCompany)
    if err != nil {
        return shim.Error("Failed to unmarshal AEOCompany: " + err.Error())
    }

    // Check if SSMId matches
    if aeoCompany.SSMId != pmtPassed.SSMId {
        return shim.Error("SSM ID does not match with the one stored in AEOCompany")
    }
    fmt.Println("SSM ID matches")

    pmtAsBytes, err := stub.GetState(pmtPassed.PermitId)
    if err != nil {
        return shim.Error("Failed to get permit: " + err.Error())
    } else if pmtAsBytes == nil {
        // Permit doesn't exist, creating new

        permit := Permit{}
        permit.ObjectType = "PERMIT"
        permit.PermitId = pmtPassed.PermitId
        permit.AEOId = pmtPassed.AEOId
        permit.SSMId = pmtPassed.SSMId
        permit.PermitClass = pmtPassed.PermitClass
        permit.ActivityType = pmtPassed.ActivityType
        permit.PermitType = pmtPassed.PermitType
        permit.ICPId = pmtPassed.ICPId
        permit.ICPExpiry = pmtPassed.ICPExpiry
        permit.PermitExpiry = pmtPassed.PermitExpiry
        permit.DestinationCountry = pmtPassed.DestinationCountry
        permit.Incoterms = pmtPassed.Incoterms
        permit.CompanyName = pmtPassed.CompanyName
        permit.ImporterName = pmtPassed.ImporterName
        permit.LogisticOperator = pmtPassed.LogisticOperator
        permit.ItemArray = pmtPassed.ItemArray
        permit.SMKStatus = pmtPassed.SMKStatus
        permit.CreatedBy = pmtPassed.CreatedBy
        permit.CreatedDate = currentLocalTime
        permit.Remarks = pmtPassed.Remarks

        fmt.Println("PERMIT UNMARSHALLED 2")

        pmtJSONasBytes, err := json.Marshal(permit)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("PERMIT MARSHALLED!")

        // Putstate
        err = stub.PutState(pmtPassed.PermitId, pmtJSONasBytes)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("PERMIT PUTSTATE PASSED!")

        // **Start of added part**
        // Add PermitId to AEOCompany
        if aeoCompany.PermitId == "" {
            aeoCompany.PermitId = pmtPassed.PermitId
        } else {
            aeoCompany.PermitId = aeoCompany.PermitId + "," + pmtPassed.PermitId
        }

        // Marshal updated AEOCompany
        updatedAeoAsBytes, err := json.Marshal(aeoCompany)
        if err != nil {
            return shim.Error("Failed to marshal updated AEOCompany: " + err.Error())
        }

        // Save updated AEOCompany to state
        err = stub.PutState(pmtPassed.AEOId, updatedAeoAsBytes)
        if err != nil {
            return shim.Error("Failed to update AEOCompany record: " + err.Error())
        }
        // **End of added part**

        fmt.Println("--- End addPermit Successfully ---")
    } else {
        return shim.Error(errExistMSG)
    }
    return shim.Success(nil)
}


//original 
// func (m *MiTrace_Generic_Chaincode) addPrs(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
//     fmt.Println("--- Start Add PRS ---")

//     if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
//         return shim.Error("Authentication failed")
//     }

//     if len(args) == 0 {
//         fmt.Println("Argument must not be empty")
//         return shim.Error("Arguments must not be empty")
//     }

//     utc := time.Now().UTC()
//     loc, err := time.LoadLocation("Asia/Singapore")
//     if err != nil {
//         return shim.Error("Failed to load location: " + err.Error())
//     }
//     currentLocalTime := utc.In(loc).Format(time.RFC1123)

//     prsParam := args[0]
//     prsPassed := PRS{}
//     err = json.Unmarshal([]byte(prsParam), &prsPassed)

//     if err != nil {
//         return shim.Error(err.Error())
//     }
//     fmt.Println("PRS UNMARSHALLED 1")

//     // Check if AEOId exists
//     fmt.Println("Checking AEO ID existence for:", prsPassed.AEOId)
//     aeoAsBytes, err := stub.GetState(prsPassed.AEOId)
//     if err != nil {
//         return shim.Error("Failed to get AEO ID: " + err.Error())
//     } else if aeoAsBytes == nil {
//         return shim.Error("AEO ID does not exist")
//     }
//     fmt.Println("AEO ID exists")

//     // Unmarshal AEOCompany to check SSMId
//     aeoCompany := AEOCompany{}
//     err = json.Unmarshal(aeoAsBytes, &aeoCompany)
//     if err != nil {
//         return shim.Error("Failed to unmarshal AEOCompany: " + err.Error())
//     }

//     // Check if SSMId matches
//     if aeoCompany.SSMId != prsPassed.SSMId {
//         return shim.Error("SSM ID does not match with the one stored in AEOCompany")
//     }
//     fmt.Println("SSM ID matches")
 
// 	 // Check if PermitId exists
// 	 permitAsBytes, err := stub.GetState(prsPassed.PermitId)
// 	 if err != nil {
// 		 return shim.Error("Failed to get Permit ID: " + err.Error())
// 	 } else if permitAsBytes == nil {
// 		 return shim.Error("Permit ID does not exist")
// 	 }
 
//     prsAsBytes, err := stub.GetState(prsPassed.PRSId)
//     if err != nil {
//         return shim.Error("Failed to get PRS: " + err.Error())
//     } else if prsAsBytes == nil {

// 		prs := PRS{}
// 		prs.ObjectType						= "PRS"	
//         prs.PRSId 							= prsPassed.PRSId
//         prs.PRSType 						= prsPassed.PRSType
//         prs.AEOId 							= prsPassed.AEOId
//         prs.SSMId 							= prsPassed.SSMId
//         prs.PermitId 						= prsPassed.PermitId
//         prs.DestinationCountry 				= prsPassed.DestinationCountry
//         prs.CompanyName 					= prsPassed.CompanyName
//         prs.EndUserDetails 					= prsPassed.EndUserDetails
//         prs.LogisticOperator 				= prsPassed.LogisticOperator
//         prs.DeclarationDateTime 			= prsPassed.DeclarationDateTime
//         prs.ReleaseDateTime 				= prsPassed.ReleaseDateTime
//         prs.ConsignmentId 					= prsPassed.ConsignmentId
//         prs.ContainerNo 					= prsPassed.ContainerNo
//         prs.ItemArray 						= prsPassed.ItemArray
//         prs.SMKNo 							= prsPassed.SMKNo
//         prs.SMKReleaseNo 					= prsPassed.SMKReleaseNo
//         prs.SMKRegisterDate 				= prsPassed.SMKRegisterDate
//         prs.SMKStatus 						= prsPassed.SMKStatus
//         prs.SMKDateRelease 					= prsPassed.SMKDateRelease
//         prs.TransportMode 					= prsPassed.TransportMode
//         prs.ApplicantName 					= prsPassed.ApplicantName
//         prs.ApplicantICNo 					= prsPassed.ApplicantICNo
//         prs.ApplicantDesignation 			= prsPassed.ApplicantDesignation
//         prs.PortCode 						= prsPassed.PortCode
//         prs.CreatedDate 					= currentLocalTime
//         prs.CreatedBy 						= prsPassed.CreatedBy
// 		// prs.ModifiedDate 				= currentLocalTime
//         // prs.ModifiedBy 					= prsPassed.ModifiedBy
//         prs.Remarks 						= prsPassed.Remarks

//    		fmt.Println("PRS UNMARSHALLED 2")
 
//         prsJSONasBytes, err := json.Marshal(prs)
//         if err != nil {
//             return shim.Error(err.Error())
//         }
//         fmt.Println("PRS MARSHALLED!")
 
//         //Putstate
//         err = stub.PutState(prsPassed.PRSId, prsJSONasBytes)
//         if err != nil {
//             return shim.Error(err.Error())
//         }
//         fmt.Println("PRS PUTSTATE PASSED!")
 
//         fmt.Println("--- End addprs Successfully ---")
//     } else {
//         return shim.Error(errExistMSG)
//     }
 
//     return shim.Success(nil)
// }

func (m *MiTrace_Generic_Chaincode) addPrs(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
    fmt.Println("--- Start Add PRS ---")

    if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
        return shim.Error("Authentication failed")
    }

    if len(args) == 0 {
        fmt.Println("Argument must not be empty")
        return shim.Error("Arguments must not be empty")
    }

    utc := time.Now().UTC()
    loc, err := time.LoadLocation("Asia/Singapore")
    if err != nil {
        return shim.Error("Failed to load location: " + err.Error())
    }
    currentLocalTime := utc.In(loc).Format(time.RFC1123)

    prsParam := args[0]
    prsPassed := PRS{}
    err = json.Unmarshal([]byte(prsParam), &prsPassed)

    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("PRS UNMARSHALLED 1")

    // Check if AEOId exists
    fmt.Println("Checking AEO ID existence for:", prsPassed.AEOId)
    aeoAsBytes, err := stub.GetState(prsPassed.AEOId)
    if err != nil {
        return shim.Error("Failed to get AEO ID: " + err.Error())
    } else if aeoAsBytes == nil {
        return shim.Error("AEO ID does not exist")
    }
    fmt.Println("AEO ID exists")

    // Unmarshal AEOCompany to check SSMId
    aeoCompany := AEOCompany{}
    err = json.Unmarshal(aeoAsBytes, &aeoCompany)
    if err != nil {
        return shim.Error("Failed to unmarshal AEOCompany: " + err.Error())
    } 

    // Check if SSMId matches
    if aeoCompany.SSMId != prsPassed.SSMId {
        return shim.Error("SSM ID does not match with the one stored in AEOCompany")
    }
    fmt.Println("SSM ID matches")

    // Check if PermitId exists
    permitAsBytes, err := stub.GetState(prsPassed.PermitId)
    if err != nil {
        return shim.Error("Failed to get Permit ID: " + err.Error())
    } else if permitAsBytes == nil {
        return shim.Error("Permit ID does not exist")
    }

    // **Start of added part**
    // Unmarshal Permit to update with PRSId
    permit := Permit{}
    err = json.Unmarshal(permitAsBytes, &permit)
    if err != nil {
        return shim.Error("Failed to unmarshal Permit: " + err.Error())
    }

    // Add PRSId to Permit
    if permit.PRSId == "" {
        permit.PRSId = prsPassed.PRSId
    } else {
        permit.PRSId = permit.PRSId + "," + prsPassed.PRSId
    }

    // Marshal updated Permit
    updatedPermitAsBytes, err := json.Marshal(permit)
    if err != nil {
        return shim.Error("Failed to marshal updated Permit: " + err.Error())
    }

    // Save updated Permit to state
    err = stub.PutState(prsPassed.PermitId, updatedPermitAsBytes)
    if err != nil {
        return shim.Error("Failed to update Permit record: " + err.Error())
    }
    // **End of added part**

    // Check if PRS exists
    prsAsBytes, err := stub.GetState(prsPassed.PRSId)
    if err != nil {
        return shim.Error("Failed to get PRS: " + err.Error())
    } else if prsAsBytes == nil {
        
		prs := PRS{}
		prs.ObjectType						= "PRS"	
        prs.PRSId 							= prsPassed.PRSId
        prs.PRSType 						= prsPassed.PRSType
		prs.PRSStatus 						= "Active"
        prs.AEOId 							= prsPassed.AEOId
        prs.SSMId 							= prsPassed.SSMId
        prs.PermitId 						= prsPassed.PermitId
        prs.DestinationCountry 				= prsPassed.DestinationCountry
        prs.CompanyName 					= prsPassed.CompanyName
        prs.EndUserDetails 					= prsPassed.EndUserDetails
        prs.LogisticOperator 				= prsPassed.LogisticOperator
        prs.DeclarationDateTime 			= prsPassed.DeclarationDateTime
        prs.ReleaseDateTime 				= prsPassed.ReleaseDateTime
        prs.ConsignmentId 					= prsPassed.ConsignmentId
        prs.ContainerNo 					= prsPassed.ContainerNo
        prs.ItemArray 						= prsPassed.ItemArray
        prs.SMKNo 							= prsPassed.SMKNo
        prs.SMKReleaseNo 					= prsPassed.SMKReleaseNo
        prs.SMKRegisterDate 				= prsPassed.SMKRegisterDate
        prs.SMKStatus 						= "Released"
        prs.SMKDateRelease 					= prsPassed.SMKDateRelease
        prs.TransportMode 					= prsPassed.TransportMode
        prs.ApplicantName 					= prsPassed.ApplicantName
        prs.ApplicantICNo 					= prsPassed.ApplicantICNo
        prs.ApplicantDesignation 			= prsPassed.ApplicantDesignation
        prs.PortCode 						= prsPassed.PortCode
        prs.CreatedDate 					= currentLocalTime
        prs.CreatedBy 						= prsPassed.CreatedBy
		// prs.ModifiedDate 				= currentLocalTime
        // prs.ModifiedBy 					= prsPassed.ModifiedBy
        prs.Remarks 						= prsPassed.Remarks

   		fmt.Println("PRS UNMARSHALLED 2")

        // Update itemBalance in Permit
		for i, permitItem := range permit.ItemArray {
			for _, prsItem := range prs.ItemArray {
				if permitItem.ItemNo == prsItem.ItemNo {
					permitItemBalance := 0.0
					if permitItem.ItemBalance != "" {
						permitItemBalance, err = strconv.ParseFloat(permitItem.ItemBalance, 64)
						if err != nil {
							return shim.Error("Failed to parse Permit Item Balance: " + err.Error())
						}
					} else {
						permitItemBalance, err = strconv.ParseFloat(permitItem.ItemQuantity, 64)
						if err != nil {
							return shim.Error("Failed to parse Permit Item Quantity: " + err.Error())
						}
					}
					prsItemQuantity, err := strconv.ParseFloat(prsItem.ItemQuantity, 64)
					if err != nil {
						return shim.Error("Failed to parse PRS Item Quantity: " + err.Error())
					}
					permitItem.ItemBalance = fmt.Sprintf("%.2f", permitItemBalance-prsItemQuantity)
					permit.ItemArray[i] = permitItem
				}
			}
		}

		// Marshal updated Permit
		permitJSONasBytes, err := json.Marshal(permit)
		if err != nil {
			return shim.Error("Failed to marshal updated Permit: " + err.Error())
		}

		// PutState for updated Permit
		err = stub.PutState(prsPassed.PermitId, permitJSONasBytes)
		if err != nil {
			return shim.Error("Failed to update Permit: " + err.Error())
		}
 
        prsJSONasBytes, err := json.Marshal(prs)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("PRS MARSHALLED!")

        // Putstate
        err = stub.PutState(prsPassed.PRSId, prsJSONasBytes)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("PRS PUTSTATE PASSED!")

        fmt.Println("--- End addprs Successfully ---")
    } else {
        return shim.Error(errExistMSG)
    }

    return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) addSandec(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
    fmt.Println("--- Start Add SANDEC ---")

    if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
        return shim.Error("Authentication failed")
    }

    if len(args) == 0 {
        fmt.Println("Argument must not be empty")
        return shim.Error("Arguments must not be empty")
    }

    utc := time.Now().UTC()
    loc, err := time.LoadLocation("Asia/Singapore")
    if err != nil {
        return shim.Error("Failed to load location: " + err.Error())
    }
    currentLocalTime := utc.In(loc).Format(time.RFC1123)

    sanDecParam := args[0]
    sanDecPassed := SANDEC{}
    err = json.Unmarshal([]byte(sanDecParam), &sanDecPassed)

    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("SANDEC UNMARSHALLED 1")

	// Check if PRSId exists
    fmt.Println("Checking PRS existence for:", sanDecPassed.PRSId)
    prsAsBytes, err := stub.GetState(sanDecPassed.PRSId)
    if err != nil {
        return shim.Error("Failed to get PRS: " + err.Error())
    } else if prsAsBytes == nil {
        return shim.Error("PRS does not exist")
    }

    // Check if AEOId exists
    fmt.Println("Checking AEO ID existence for:", sanDecPassed.AEOId)
    aeoAsBytes, err := stub.GetState(sanDecPassed.AEOId)
    if err != nil {
        return shim.Error("Failed to get AEO ID: " + err.Error())
    } else if aeoAsBytes == nil {
        return shim.Error("AEO ID does not exist")
    }
    fmt.Println("AEO ID exists")

    // Unmarshal AEOCompany to check SSMId
    aeoCompany := AEOCompany{}
    err = json.Unmarshal(aeoAsBytes, &aeoCompany)
    if err != nil {
        return shim.Error("Failed to unmarshal AEOCompany: " + err.Error())
    }

    // Check if SSMId matches
    if aeoCompany.SSMId != sanDecPassed.SSMId {
        return shim.Error("SSM ID does not match with the one stored in AEOCompany")
    }
    fmt.Println("SSM ID matches")

    // Check if PermitId exists
    fmt.Println("Checking Permit existence for:", sanDecPassed.PermitId)
    permitAsBytes, err := stub.GetState(sanDecPassed.PermitId)
    if err != nil {
        return shim.Error("Failed to get Permit: " + err.Error())
    } else if permitAsBytes == nil {
        return shim.Error("Permit does not exist")
    }

    // Unmarshal PRS to check status
    prs := PRS{}
    err = json.Unmarshal(prsAsBytes, &prs)
    if err != nil {
        return shim.Error("Failed to unmarshal PRS: " + err.Error())
    }

    // Check if SANDEC already exists
    sanDecAsBytes, err := stub.GetState(sanDecPassed.SANDECId)
    if err != nil {
        return shim.Error("Failed to get SANDEC: " + err.Error())
    } else if sanDecAsBytes == nil {

        	sanDec := SANDEC{}
            sanDec.ObjectType          	= "SANDEC"
            sanDec.SANDECId            	= sanDecPassed.SANDECId
            sanDec.PRSId               	= sanDecPassed.PRSId
            sanDec.AEOId               	= sanDecPassed.AEOId
            sanDec.SSMId               	= sanDecPassed.SSMId
            sanDec.PermitId            	= sanDecPassed.PermitId
            sanDec.DestinationCountry  	= sanDecPassed.DestinationCountry
            sanDec.CompanyName         	= sanDecPassed.CompanyName
            sanDec.DeclarationDateTime 	= sanDecPassed.DeclarationDateTime
            sanDec.ReleaseDateTime     	= sanDecPassed.ReleaseDateTime
            sanDec.ConsignmentId       	= sanDecPassed.ConsignmentId
            sanDec.ContainerNo         	= sanDecPassed.ContainerNo
            sanDec.ItemArray           	= sanDecPassed.ItemArray
            sanDec.SMKNo               	= sanDecPassed.SMKNo
            sanDec.SMKReleaseNo        	= sanDecPassed.SMKReleaseNo
            sanDec.SMKStatus           	= "Released"
            sanDec.SMKDateRelease      	= sanDecPassed.SMKDateRelease
            sanDec.TransportMode       	= sanDecPassed.TransportMode
            sanDec.ApplicantName       	= sanDecPassed.ApplicantName
            sanDec.PortCode            	= sanDecPassed.PortCode
            sanDec.CreatedDate         	= currentLocalTime
            sanDec.CreatedBy           	= sanDecPassed.CreatedBy
            sanDec.Remarks             	= sanDecPassed.Remarks

        fmt.Println("SANDEC UNMARSHALLED 2")

        sanDecJSONasBytes, err := json.Marshal(sanDec)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("SANDEC MARSHALLED!")

        // PutState for SANDEC
        err = stub.PutState(sanDecPassed.SANDECId, sanDecJSONasBytes)
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("SANDEC PUTSTATE PASSED!")

        fmt.Println("--- End addSandec Successfully ---")
    } else {
        return shim.Error("SANDEC already exists")
    }

    return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) addUser(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {

	// Input sanitation
	fmt.Println("--- Start Add User for AEO Project ---")

	if !m.testMode && !authenticateOrg1(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authOrg1)
	}

	// Check args length
	if len(args) != 3 {
		return shim.Error(argsErr + "3")
	}

	// Grab args
	if len(args[0]) <= 0 {
		return shim.Error("User ID Required !")
	}
	if len(args[1]) <= 0 {
		return shim.Error("User Email Required !")
	}
	if len(args[2]) <= 0 {
		return shim.Error("User Status Required !")
	}

	userID := args[0]
	email := args[1]
	status := args[2]

	trParam := "{\"userID\":\"" + userID + "\",\"identifier\":\"USER\"}"
	trPassed := User{} //create instance to client
	err2 := json.Unmarshal([]byte(trParam), &trPassed)
	if err2 != nil {
		return shim.Error(err2.Error())
	}
	fmt.Println("USER UNMARSHALLED")
	tiAsBytes, err := stub.GetState(trPassed.UserID)
	if err != nil {
		return shim.Error("Failed to get user : " + err.Error())
	} else if tiAsBytes == nil {
		// Userdoesn't exist, creating new
		user := User{
			ObjectType: "USER",
			UserID:     userID,
			UserEmail:  email,
			UserStatus: status,
		}

		userJSONasBytes, err := json.Marshal(user)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("USER MARSHALLED!")

		//Putstate
		err = stub.PutState(userID, userJSONasBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("USER PUTSTATE PASSED!")

		fmt.Println("--- end addUser successfully ---")
	} else {
		// return shim.Error(strings.Replace(errExistMSG, "!", "!\"\n\"key\":\"" + userID, 1))
		return shim.Error(errExistMSG)
	}
	return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) addSanity(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
	
	// Input sanitation
	fmt.Println("--- Start Add Sanity for AEO Project---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		// fmt.Println("Creator: ", creatorOrg)
		// fmt.Println("Issuer: ", creatorCertIssuer)
		return shim.Error(authAll)
	}

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

	sanityParam := args[0]
	sanityPassed := CheckSanity{}
	fmt.Printf("sanityParam = " + sanityParam + "\n")

	err := json.Unmarshal([]byte(sanityParam), &sanityPassed)

	if err != nil {
		fmt.Println(err.Error())
		return shim.Error(err.Error())
	}

	fmt.Println("SANITY UNMARSHALLED 1")

	fmt.Printf("SANITY ID "+ sanityPassed.CheckID + "\n")
	comKey := getSanityIDKey("SANITY", sanityPassed.CheckID)
	fmt.Println("SANITY ID HASH " + comKey)
	
	srecordAsBytes, err := stub.GetState(sanityPassed.CheckID)
	if err != nil {
		return shim.Error("Failed to get Id record : " + err.Error())
	} else if srecordAsBytes == nil {
		Sanity := CheckSanity{}
		Sanity.ObjectType 	= "SANITY"
		Sanity.CheckID 		= sanityPassed.CheckID
		Sanity.DateTime 	= currentLocalTime
		Sanity.TestValue 	= sanityPassed.TestValue

		// err = json.Unmarshal(srecordAsBytes, &Sanity)
		// if err != nil {
		// 	return shim.Error(err.Error())
		// }

		fmt.Println("SANITY UNMARSHALLED 2")

		sanityJSONAsBytes, err := json.Marshal(Sanity)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("SANITY MARSHALLED!")

		//Putstate
		err = stub.PutState(sanityPassed.CheckID, sanityJSONAsBytes)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("SANITY PUTSTATE PASSED!")

		fmt.Println("--- end addSanity successfully ---")
	} else {
		// return shim.Error(strings.Replace(errExistMSG, "!", "!\"\n\"key\":\"" + sanityPassed.CheckID, 1))
		return shim.Error(errExistMSG)

	}
	return shim.Success(nil)
}

func (m *MiTrace_Generic_Chaincode) addSanityFlat(stub shim.ChaincodeStubInterface, creatorOrg string, creatorCertIssuer []string, args []string) pb.Response {
	
	fmt.Println("--- Start addSanityFlat for AEO Project---")

	if !m.testMode && !authenticateAllOrg(creatorOrg, creatorCertIssuer) {
		return shim.Error(authAll)
	}

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

	sanityParam := args[0]
	sanityPassed := CheckSanity{}
	err := json.Unmarshal([]byte(sanityParam), &sanityPassed)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("SANITY UNMARSHALLED 1")
	
	fmt.Printf("SANITY ID " + sanityPassed.CheckID + "\n")
	comKey := getSanityIDKey("SANITY", sanityPassed.CheckID)
	fmt.Println("SANITY ID HASH " + comKey)
	
	// CheckSan := CheckSanity{
	// 	ObjectType:    "SANITY",
	// 	CheckID:       	chkId,
	// 	DateTime: 		currentLocalTime,
	// 	TestValue:   	tValue,
	// }

	Sanity := CheckSanity{}
	Sanity.ObjectType 	= "SANITY"
	Sanity.CheckID 		= sanityPassed.CheckID
	Sanity.DateTime 	= currentLocalTime
	Sanity.TestValue 	= sanityPassed.TestValue

	fmt.Println("SANITY UNMARSHALLED 2")

	checkJSONAsBytes, err := json.Marshal(Sanity)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("SANITY MARSHALLED!")

	err = stub.PutState(sanityPassed.CheckID, checkJSONAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("SANITY PUTSTATE PASSED!")

	fmt.Println("--- End addSanityFlat Successfully ---")
	
	return shim.Success(nil)
}
