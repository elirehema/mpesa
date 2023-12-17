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

### Example
```go

```

## BUSINESS TO CUSTOMER

The B2C API Call is used as a standard business-to-customer funds disbursement. Funds from the business account’s wallet will be deducted and paid to the mobile money wallet of the customer. Use cases for the B2C includes:
-  Salary payments 
-  Funds transfers from business
- Charity pay-out

## BUSINESS TO BUSINESS

The B2B API Call is used for business-to-business transactions. Funds from the business’ mobile money wallet will be deducted and transferred to the mobile money wallet of the other business. Use cases for the B2C includes:
- Stock purchases 
- Bill payment 
- Ad-hoc payment


## REVERSAL

The Reversal API is used to reverse a successful transaction. Using the Transaction ID of a previously successful transaction,  the OpenAPI will withdraw the funds from the recipient party’s mobile money wallet and revert the funds to the mobile money wallet of the initiating party of the original transaction.

## QUERY TRANSACTION STATUS

The Query Transaction Status API call is used to query the status of the transaction that has been initiated.

## QUERY BENEFICIARY NAME

The Query Beneficiary Name API call is used to query the details about a customer. Information that can be retrieved when using this API is: 1. Customer First Name 2. Customer Last Name

