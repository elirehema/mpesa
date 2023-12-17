package datas

type RequestCustomerToBusiness struct {
	Amount              string `json:"input_Amount"  binding:"required"`
	Msisdn              string `json:"input_CustomerMSISDN"  binding:"required"`
	Country             string `json:"input_Country"  binding:"required"`
	Currency            string `json:"input_Currency"  binding:"required"`
	ServiceProviderCode string `json:"input_ServiceProviderCode"  binding:"required"`
	TransactionRef      string `json:"input_TransactionReference"  binding:"required"`
	ConversationId      string `json:"input_ThirdPartyConversationID"  binding:"required"`
	Description         string `json:"input_PurchasedItemsDesc"  binding:"required"`
}

type RequestBusinessToCustomer struct {
	Amount              string `json:"input_Amount"  binding:"required"`
	Msisdn              string `json:"input_CustomerMSISDN"  binding:"required"`
	Country             string `json:"input_Country"  binding:"required"`
	Currency            string `json:"input_Currency"  binding:"required"`
	ServiceProviderCode string `json:"input_ServiceProviderCode"  binding:"required"`
	TransactionRef      string `json:"input_TransactionReference"  binding:"required"`
	ConversationId      string `json:"input_ThirdPartyConversationID"  binding:"required"`
	Description         string `json:"input_PurchasedItemsDesc"  binding:"required"`
}

type RequestBusinessToBusiness struct {
	Amount         string `json:"input_Amount"  binding:"required"`
	ReceiverCode   string `json:"input_ReceiverPartyCode"  binding:"required"`
	Country        string `json:"input_Country"  binding:"required"`
	Currency       string `json:"input_Currency"  binding:"required"`
	SenderCode     string `json:"input_PrimaryPartyCode"  binding:"required"`
	TransactionRef string `json:"input_TransactionReference"  binding:"required"`
	ConversationId string `json:"input_ThirdPartyConversationID"  binding:"required"`
	Description    string `json:"input_PurchasedItemsDesc"  binding:"required"`
}
type RequestReversalTransaction struct {
	Amount              string `json:"input_Amount"  binding:"required"`
	Country             string `json:"input_Country"  binding:"required"`
	TransactionID       string `json:"input_TransactionID"  binding:"required"`
	ServiceProviderCode string `json:"input_ServiceProviderCode"  binding:"required"`
	ConversationId      string `json:"input_ThirdPartyConversationID"  binding:"required"`
}
type RequestQueryTransactionStatus struct {
	QueryReference      string `json:"input_QueryReference"  binding:"required"`
	Country             string `json:"input_Country"  binding:"required"`
	ServiceProviderCode string `json:"input_ServiceProviderCode"  binding:"required"`
	ConversationId      string `json:"input_ThirdPartyConversationID"  binding:"required"`
}
type RequestQueryBeneficiaryName struct {
	Msisdn              string `json:"input_CustomerMSISDN"  binding:"required"`
	Country             string `json:"input_Country"  binding:"required"`
	ServiceProviderCode string `json:"input_ServiceProviderCode"  binding:"required"`
	KYCQueryType        string `json:"input_KYCQueryType"  binding:"required"`
}
