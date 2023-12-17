package mpesa

import (
	"github.com/elirehema/mpesa/datas"
	"github.com/elirehema/mpesa/sandbox"
	"github.com/joho/godotenv"
)

const (
	SESSION                  = "/getSession/"
	C2BPAYMENT               = "/c2bPayment/singleStage/"
	B2CPAYMENT               = "/b2cPayment/singleStage/"
	B2BPAYMENT               = "/b2bPayment/singleStage/"
	REVERSAL                 = "/reversal"
	QUERY_BENEFICIARY_NAME   = "/queryBeneficiaryName/"
	QUERY_TRANSACTION_STATUS = "/queryTransactionStatus/"
)

func RefreshSessionKey() {
	godotenv.Load("session.env")
	result := datas.SessionResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).Get(SESSION)
	if err != nil {
		print(err.Error())
	} else {
		env, _ := godotenv.Unmarshal("API_SESSION_KEY=" + result.SessionId)
		godotenv.Write(env, "./session.env")

		print(&result)
	}
}
func CustomerToBusiness(request datas.RequestCustomerToBusiness) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post(C2BPAYMENT)
	if err != nil {
		print(err.Error())
	}
	return result
}
func BusinessToCustomer(request datas.RequestBusinessToCustomer) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post(B2CPAYMENT)
	if err != nil {
		print(err.Error())
	}
	return result
}
func BusinessToBusiness(request datas.RequestBusinessToBusiness) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post(B2BPAYMENT)
	if err != nil {
		print(err.Error())
	}
	return result
}
func ReverseTransaction(request datas.RequestReversalTransaction) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).Post(REVERSAL)
	if err != nil {
		print(err.Error())
	}
	return result
}
func QueryBeneficiaryName(request datas.RequestQueryBeneficiaryName) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).SetQueryParams(map[string]string{
		"input_CustomerMSISDN":      request.Msisdn,
		"input_Country":             request.Country,
		"input_ServiceProviderCode": request.ServiceProviderCode,
		"input_KYCQueryType":        request.KYCQueryType,
	}).Post(QUERY_BENEFICIARY_NAME)
	if err != nil {
		print(err.Error())
	}
	return result
}
func QueryTransactionStatus(request datas.RequestQueryTransactionStatus) datas.ResponseData {
	result := datas.ResponseData{}
	_, err := sandbox.MpesaClient.R().SetResult(&result).SetBody(&request).SetQueryParams(map[string]string{
		"input_QueryReference":           request.QueryReference,
		"input_Country":                  request.Country,
		"input_ServiceProviderCode":      request.ServiceProviderCode,
		"input_ThirdPartyConversationID": request.ConversationId,
	}).Get(QUERY_TRANSACTION_STATUS)
	if err != nil {
		print(err.Error())
	}
	return result
}
