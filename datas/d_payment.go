package datas

type PaymentData struct {
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
}
