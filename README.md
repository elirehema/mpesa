# mpesa
M-PESA Open API for Golang 

## NOTE
This dependency was not created or maintained by M-PESA Team. It is part of Open Source projects and you can clone and modify as you per your requirements or contribute for improvement. 

## Getting Started 
This dependency is a quick start to work with [M-PESA Open API's](https://openapiportal.m-pesa.com/api-documentation).

## Usage 
To start use the depenency get mpesa dependency by

```golang
go get github.com/elirehema/mpesa
```
Create a `.env` file in the root of your project and add the following properties

```.env
ENVIROMENT=DEV //specify if it is production or not DEV | PROD
COUNTRY={{YOUR COUNTRY}} e.g vodacomTZN
API_PUBLIC_KEY={{API PUBLIC KEY}}
API_KEY={{API KEY}}
```

## CUSTOMER TO BUSINESS

The C2B API call is used as a standard customer-to-business transaction. Funds from the customer’s mobile money wallet will be deducted and be transferred to the mobile money wallet of the business. To authenticate and authorize this transaction, M-Pesa Payments Gateway will initiate a USSD Push message to the customer to gather and verify the mobile money PIN number. This number is not stored and is used only to authorize the transaction.

#### Example
```go

import (
	"github.com/elirehema/mpesa"
	"github.com/elirehema/mpesa/datas"
)

func main() {
	customerToBusinessTransaction()
}

func makeCustomerToBusinessTransaction() {
	transaction := datas.RequestCustomerToBusiness{
		Amount:              "10",
		Country:             "TZN",
		Currency:            "TZS",
		Msisdn:              "000000000001",
		ServiceProviderCode: "000000",
		ConversationId:      "asv02e5958774f7ba228d83d0d689761",
	
```

## BUSINESS TO CUSTOMER

The B2C API Call is used as a standard business-to-customer funds disbursement. Funds from the business account’s wallet will be deducted and paid to the mobile money wallet of the customer. Use cases for the B2C includes:
-  Salary payments 
-  Funds transfers from business
- Charity pay-out

#### Example

```go
import (
	"github.com/elirehema/mpesa"
	"github.com/elirehema/mpesa/datas"
)

func main() {
	makeBusinessToCustomerTransaction()
}

func makeBusinessToCustomerTransaction() {
	transaction := datas.RequestBusinessToCustomer{
		Amount:              "10",
		Country:             "TZN",
		Currency:            "TZS",
		Msisdn:              "000000000001",
		ServiceProviderCode: "000000",
		ConversationId:      "asv02e5958774f7ba228d83d0d68976a",
		TransactionRef:      "T1234xvs",
		Description:         "Salary Payment",
	}

	response := mpesa.BusinessToCustomer(transaction)
	print(response.ThirdPartyConversationID)
}

```

## BUSINESS TO BUSINESS

The B2B API Call is used for business-to-business transactions. Funds from the business’ mobile money wallet will be deducted and transferred to the mobile money wallet of the other business. Use cases for the B2C includes:
- Stock purchases 
- Bill payment 
- Ad-hoc payment

#### Example
```go
import (
	"github.com/elirehema/mpesa"
	"github.com/elirehema/mpesa/datas"
)

func main() {
	makeRequestBusinessToBusiness()
}

func makeRequestBusinessToBusiness() {
	transaction := datas.RequestBusinessToBusiness{
		Amount:         "10",
		Country:        "TZN",
		Currency:       "TZS",
		SenderCode:     "000000",
		ReceiverCode:   "000001",
		ConversationId: "asv02e5958774f7ba228d83d0d689762",
		TransactionRef: "T1234xvs",
		Description:    "Electricity Payment",
	}

	response := mpesa.BusinessToBusiness(transaction)
	print(response.ThirdPartyConversationID)
}

```


## REVERSAL

The Reversal API is used to reverse a successful transaction. Using the Transaction ID of a previously successful transaction,  the OpenAPI will withdraw the funds from the recipient party’s mobile money wallet and revert the funds to the mobile money wallet of the initiating party of the original transaction.

####  Example
```go

import (
	"github.com/elirehema/mpesa"
	"github.com/elirehema/mpesa/datas"
)

func main() {
	makeRequestReversalTransaction()
}

func makeRequestReversalTransaction() {
	transaction := datas.RequestReversalTransaction{
		Amount:              "10",
		Country:             "TZN",
		ServiceProviderCode: "000000",
		ConversationId:      "asv02e5958774f7ba228d83d0d689761",
		TransactionID:       "0000000000001",
	}

	response := mpesa.ReverseTransaction(transaction)
	print(response.ThirdPartyConversationID)
}


```

## QUERY TRANSACTION STATUS

The Query Transaction Status API call is used to query the status of the transaction that has been initiated.

```go

import (
	"github.com/elirehema/mpesa"
	"github.com/elirehema/mpesa/datas"
)

func main() {
	makeRequestQueryTransactionStatus()
}

func makeRequestQueryTransactionStatus() {
	transaction := datas.RequestQueryTransactionStatus{
		QueryReference:      "000000000000000000001",
		ServiceProviderCode: "000000",
		ConversationId:      "asv02e5958774f7ba228d83d0d689761",
		Country:             "TZN",
	}

	response := mpesa.QueryTransactionStatus(transaction)
	print(response.ThirdPartyConversationID)
}


```

## QUERY BENEFICIARY NAME

The Query Beneficiary Name API call is used to query the details about a customer. Information that can be retrieved when using this API is: 1. Customer First Name 2. Customer Last Name

```go

func makeRequestQueryBeneficiaryNameRequest() {
	transaction := datas.RequestQueryBeneficiaryName{
		Msisdn:              "000001",
		ServiceProviderCode: "000000",
		KYCQueryType:        "Name",
		Country:             "TZN",
	}

	response := mpesa.QueryBeneficiaryName(transaction)
	print(response.ThirdPartyConversationID)
}

```

### DEFINITIONS

| Field Name |  Description  | Mandatory | Regex Validation | Possible Values |
|:-----------|:-------------|:---------:|:-----------------|:----------------|
| Amount     | Transaction Amount      | True     | ^\d*\.?\d+$ | 100.0  |
| Country    |  The origin country     | True  | N/A     |  GHA,TZN,LES,DRC | 
| Currency   | Transaction Currency |    True |   ^[a-zA-Z]{1,3}$    | TZS,GHS,LSL,USD   |
| MSISDN   |  Customer Number |   True |  ^[0-9]{12,14}$     | 255707161122      |
| ServiceProviderCode   | Organization Short Code |  True | ^([0-9A-Za-z]{4,12})$  | ORG001 |
| ConversationID   | The third party's transaction reference on their system. | True | ^[0-9a-zA-Z \w+]{1,40}$  | 1e9b774d1da34af78412a498cbc28f5e |
| Description   | Other descriptions |   True |  ^[0-9a-zA-Z \w+]{1,256}$  |  Handbag, Black shoes  |
| TransactionRef   | transaction reference |    True | ^[0-9a-zA-Z \w+]{1,20}$  | T12344C  |
| SenderCode   | Organiation Short Code |   True |   ^([0-9A-Za-z]{4,12})$   |   ORG001    |
| ReceiverCode   |  Organiation Short Code_ |   True |  ^([0-9A-Za-z]{4,12})$    |  ORG002     |
| QueryReference   | The transaction's ID being queried. |    True | ^(?:[A-Za-z0-9]{10,20} |   ............    |
| KYCQueryType   | The type of KYC information to be queried, currently only "Name" is supported. |  True | ^(Name)$ | Name      |


#### QueryReference

The transaction's ID being queried. This can be one of 3 values: 
1. Mobile Money's TransactionID
2. Developer's ThirdPartyConversationID
3. OpenAPI's ConversationID