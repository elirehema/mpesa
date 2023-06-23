package datas

type ResponseData struct {
	ResponseCode             string `json:"output_ResponseCode"`
	ResponseDescription      string `json:"output_ResponseDesc"`
	TransactionID            string `json:"output_TransactionID"`
	ConversationID           string `json:"output_ConversationID"`
	ThirdPartyConversationID string `json:"output_ThirdPartyConversationID"`
	TransactionStatus        string `json:"output_ResponseTransactionStatus"`
	OriginalTransactionID    string `json:"output_OriginalTransactionID"`
	IsReversed               string `json:"output_Reversed"`
}
