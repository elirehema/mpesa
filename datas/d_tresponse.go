package datas

type SessionResponseData struct {
	ResponseCode string `json:"output_ResponseCode"`
	ResponseDesc string `json:"output_ResponseDesc,omitempty"`
	SessionId    string `json:"output_SessionID"`
}
