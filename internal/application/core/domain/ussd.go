package domain

import "time"

type UssdRequest struct {
	Payload     string `json:"ussdPayload"`
	Language    string `json:"language"`
	SessionId   string `json:"sessionId"`
	Msisdn      string `json:"msisdn"`
	ServiceCode string `json:"serviceCode"`
	MessageType string `json:"ussdMessageType"`
	CreatedAt   int64  `json:"created_at"`
}

type UssdResponse struct {
	Payload     string `json:"ussdPayload"`
}
func NewUssdRequest(payload string, language string, sessionId string, msisdn string) UssdRequest {
	return UssdRequest{
		Payload:   payload,
		Language:  language,
		SessionId: sessionId,
		Msisdn:    msisdn,
		CreatedAt: time.Now().Unix(),
	}
}
