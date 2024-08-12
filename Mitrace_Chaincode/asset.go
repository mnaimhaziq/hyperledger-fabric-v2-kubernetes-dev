package main

// AEOCompany that holds the AEO Company data
type AEOCompany struct {
	ObjectType 		    string `json:"identifier"`
	AEOId   		    string `json:"aeoId"`
	SSMId 			    string `json:"ssmId"`
	PermitId   				string 		`json:"permitId"`
	CompanyName 	    string `json:"companyName"`
	CompanyAddress      string `json:"companyAddress"`
	CompanyTel          string `json:"companyTel"`
	CompanyFax          string `json:"companyFax"`
	AEOFileRefNo        string `json:"aeoFileRefNo"` 
	AEOStatus 		    string `json:"aeoStatus"`
	///TODO:to be comply with standard modified date and refactor
	ModifiedBy   		string `json:"modifiedBy,omitempty"`
	Remarks    		    string `json:"remarks,omitempty"`
	Version    		    string `json:"version"`
	Reserved    		string `json:"reserved"`
	EffectiveDateFrom   string `json:"effectiveDateFrom"`
	EffectiveDateTo     string `json:"effectiveDateTo"`
	AEOCardIssued       string `json:"aeoCardIssued"`
	ForwarderCardIssued string `json:"forwarderCardIssued"`
	CreatedDate		    string `json:"createdDate"`
	CreatedBy		    string `json:"createdBy"`
	ModifiedDate 		string `json:"modifiedDate,omitempty"`
}

// Permit that holds the Permit data
type Permit struct {
	ObjectType 				string 		`json:"identifier"`
	PermitId   				string 		`json:"permitId"`
	AEOId 					string 		`json:"aeoId"`
	SSMId 					string 		`json:"ssmId"`
	PRSId   				string 		`json:"prsId"`
	PermitClass  			string 		`json:"permitClass"`
	ActivityType 			string 		`json:"activityType"`
	PermitType   			string 		`json:"permitType"`
	ICPId 					string 		`json:"icpId"`
	ICPExpiry 				string 		`json:"icpExpiry"`
	PermitExpiry  			string 		`json:"permitExpiry"` 
	PermitStatus  			string 		`json:"permitStatus"` 
	DestinationCountry 		string 		`json:"destinationCountry"`
	Incoterms   			string 		`json:"incoterms"`
	CompanyName 			string 		`json:"companyName"` 			//ExporterName/Consignor
	ImporterName 			string 		`json:"importerName"` 			//Consignee
	LogisticOperator  		string 		`json:"logisticOperator"`
	ItemArray				[]Item		`json:"itemDetails"`			//(ItemNo, OriginCountry, ItemClassification, HsCode, ItemValueMyr, ItemValueOther, ItemQuantity, ItemBalance)
	SMKStatus 				string 		`json:"smkStatus"`
	CreatedDate  			string 		`json:"createdDate"`
	CreatedBy 				string 		`json:"createdBy"`
	ModifiedDate			string 		`json:"modifiedDate,omitempty"`
	ModifiedBy   			string 		`json:"modifiedBy,omitempty"`
	Remarks    				string 		`json:"remarks,omitempty"`
}

type PRS struct {
	ObjectType 				string 		`json:"identifier"`
	PRSId   				string 		`json:"prsId"`
	AEOId 					string 		`json:"aeoId"`
	SSMId 					string 		`json:"ssmId"`
	PermitId   				string 		`json:"permitId"`
	PRSType   				string 		`json:"prsType"`
	PRSStatus   			string 		`json:"prsStatus"`
	DestinationCountry 		string 		`json:"destinationCountry"`
	CompanyName 			string 		`json:"companyName"` 			//ExporterName/Consignor
	EndUserDetails 			string 		`json:"endUserDetails"`			//NEW
	LogisticOperator  		string 		`json:"logisticOperator"`      	//TBC?
	DeclarationDateTime  	string 		`json:"declarationDateTime"`	//NEW
	ReleaseDateTime  		string 		`json:"releaseDateTime"`		//NEW
	ConsignmentId  			string 		`json:"consignmentId"`			//TBC? same as consigmentNo?
	ContainerNo  			string 		`json:"containerNo"`			//TBC?
	ItemArray				[]Item		`json:"itemDetails"`			//(ItemNo, OriginCountry, Description, ItemClassification, HsCode, ItemValueMYR, ItemValueOtherCurrency, ItemQuantity, UOM, CurrencyCode, ItemBalance)
	SMKNo 					string 		`json:"smkNo"`					//NEW
	SMKReleaseNo			string 		`json:"smkReleaseNo"`			//NEW
	SMKRegisterDate			string 		`json:"smkRegisterDate"`		//NEW
	SMKStatus 				string 		`json:"smkStatus"`
	SMKDateRelease			string 		`json:"smkDateRelease"`			//NEW
	TransportMode			string 		`json:"transportMode"`			//NEW
	ApplicantName			string 		`json:"applicantName"`			//NEW
	ApplicantICNo			string 		`json:"applicantIcNo"`			//NEW
	ApplicantDesignation	string 		`json:"applicantDesignation"`	//NEW
	PortCode				string 		`json:"portCode"`				//NEW
	//more data..
	CreatedDate  			string 		`json:"createdDate"`
	CreatedBy 				string 		`json:"createdBy"`
	ModifiedDate			string 		`json:"modifiedDate,omitempty"`
	ModifiedBy   			string 		`json:"modifiedBy,omitempty"`
	Remarks    				string 		`json:"remarks,omitempty"`
}

type SANDEC struct {
	ObjectType          string  `json:"identifier"`
	SANDECId            string  `json:"sandecId"`
	PRSId               string  `json:"prsId"`
	AEOId               string  `json:"aeoId"`
	SSMId               string  `json:"ssmId"`
	PermitId            string  `json:"permitId"`
	DestinationCountry  string  `json:"destinationCountry"`
	CompanyName         string  `json:"companyName"`
	DeclarationDateTime string  `json:"declarationDateTime"`
	ReleaseDateTime     string  `json:"releaseDateTime"`
	ConsignmentId       string  `json:"consignmentId"`
	ContainerNo         string  `json:"containerNo"`
	ItemArray           []Item  `json:"itemDetails"` // (ItemNo, OriginCountry, Description, ItemClassification, HsCode, ItemValueMYR, ItemQuantity, UOM)
	SMKNo               string  `json:"smkNo"`
	SMKReleaseNo        string  `json:"smkReleaseNo"`
	SMKStatus           string  `json:"smkStatus"`
	SMKDateRelease      string  `json:"smkDateRelease"`
	TransportMode       string  `json:"transportMode"`
	ApplicantName       string  `json:"applicantName"`
	PortCode            string  `json:"portCode"`
	CreatedDate         string  `json:"createdDate"`
	CreatedBy           string  `json:"createdBy"`
	ModifiedDate        string  `json:"modifiedDate,omitempty"`
	ModifiedBy          string  `json:"modifiedBy,omitempty"`
	Remarks             string  `json:"remarks,omitempty"`
}

type Item struct {
	ItemNo 					string `json:"itemNo"`					//Item 1
	OriginCountry   		string `json:"originCountry"`
	Description   			string `json:"description"`				//NEW
	ItemClassification 		string `json:"itemClassification"`
	HsCode 					string `json:"hsCode"`					//ItemID
	ItemValueMYR  			string `json:"itemValueMyr"`
	ItemValueOtherCurrency 	string `json:"itemValueOther"`
	ItemQuantity   			string `json:"itemQuantity"`
	UOM   					string `json:"uom"`						//NEW
	CurrencyCode   			string `json:"currencyCode"`			//NEW
	ItemBalance   			string `json:"itemBalance"`				//OLD but is this = Total Price in the excel sheets?
}

// Sanity Check
type CheckSanity struct {
	ObjectType    	string `json:"identifier"`
	CheckID       	string `json:"checkId"`
	DateTime 		string `json:"dateTime"`
	TestValue   	string `json:"testValue"`
	ModifyTime 		string `json:"modifyTime,omitempty"`
	ModifyBy   		string `json:"modifyBy,omitempty"`
	Remarks    		string `json:"remarks,omitempty"`
}

// User data that corelates with the business location
type User struct {
	ObjectType string `json:"identifier"`
	UserID     string `json:"userID"`
	UserEmail  string `json:"email"`
	UserStatus string `json:"status"`
}

// Query object
type QueryObject struct {
	Key    	string
	TxId	string
	Record 	Item
}
