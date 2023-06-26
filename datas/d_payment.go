package datas

type RequestData struct {
	Amount                 string `json:"input_Amount,omitempty"`
	Country                string `json:"input_Country,omitempty"`
	Currency               string `json:"input_Currency,omitempty"`
	SenderCode             string `json:"input_PrimaryPartyCode,omitempty"`
	ReceiverCode           string `json:"input_ReceiverPartyCode,omitempty"`
	ConversationId         string `json:"input_ThirdPartyConversationID,omitempty"`
	TransactionRef         string `json:"input_TransactionReference,omitempty"`
	Description            string `json:"input_PurchasedItemsDesc,omitempty"`
	Msisdn                 string `json:"input_CustomerMSISDN,omitempty"`
	ServiceProviderCode    string `json:"input_ServiceProviderCode,omitempty"`
	PaymentItemDescription string `json:"input_PaymentItemsDesc,omitempty"`
	QueryReference         string `json:"input_QueryReference,omitempty"`
	OriginalConversationId string `json:"input_OriginalConversationID,omitempty"`
	CustomerFirstName      string `json:"input_CustomerFirstName,omitempty"`
	CustomerLastName       string `json:"input_CustomerLastName,omitempty"`
	ResultCode             string `json:"input_ResultCode,omitempty"`
	ResultDescription      string `json:"input_ResultDesc,omitempty"`
	KYCQueryType           string `json:"input_KYCQueryType,omitempty"`
}
