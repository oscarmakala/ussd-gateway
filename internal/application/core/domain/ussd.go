package domain

type UssdRequest struct {
	Payload     string `json:"ussdPayload"`
	Language    string `json:"language"`
	SessionId   string `json:"sessionId"`
	Msisdn      string `json:"msisdn"`
	ServiceCode string `json:"serviceCode"`
	MessageType string `json:"ussdMessageType"`
}

