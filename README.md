# mpesa
M-PESA Golang dependency

## Getting Started 
This dependency is a quick start to work with [M-PESA Open API's](https://openapiportal.m-pesa.com/api-documentation).

## Usage 
To start use the depenency get mpesa dependency by

```golang
go get github.com/elirehema/mpesa
```
Create a `.env` file in the root of your project and add the following properties

```.env
BASE_URL=https://openapi.m-pesa.com/sandbox/ipg/v2
COUNTRY={{YOUR COUNTRY e.g vodacomTZN}}
API_PUBLIC_KEY={{API PUBLIC KEY}}
API_KEY={{API KEY}}
```