package mpesa

import (
	"mpesa/datas"
	"mpesa/sandbox"

	"github.com/joho/godotenv"
)

func RefreshSessionKey() {
	godotenv.Load("session.env")
	result := datas.SessionResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).Get("/getSession/")
	if err != nil {
		print(err.Error())
	} else {
		env, _ := godotenv.Unmarshal("API_SESSION_KEY=" + result.SessionId)
		godotenv.Write(env, "./session.env")

		print(&result)
	}
}
func CustomerToBusiness(request datas.RequestData) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post("/c2bPayment/singleStage/")
	if err != nil {
		print(err.Error())
	}
	return result
}
func BusinessToCustomer(request datas.RequestData) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post("/b2cPayment/singleStage/")
	if err != nil {
		print(err.Error())
	}
	return result
}
func BusinessToBusiness(request datas.RequestData) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post("/b2bPayment/singleStage/")
	if err != nil {
		print(err.Error())
	}
	return result
}
func ReverseTransaction(request datas.RequestData) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post("/reversal")
	if err != nil {
		print(err.Error())
	}
	return result
}
func QueryBeneficiaryName(request datas.RequestData) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).SetQueryParams(map[string]string{
		"input_CustomerMSISDN":      request.Msisdn,
		"input_Country":             request.Country,
		"input_ServiceProviderCode": request.ServiceProviderCode,
		"input_KYCQueryType":        request.KYCQueryType,
	}).Post("/queryBeneficiaryName/")
	if err != nil {
		print(err.Error())
	}
	return result
}
