package telegramgateway

type SendVerificationMessageRequest struct {
	PhoneNumber    string `json:"phone_number"`
	RequestID      string `json:"request_id,omitempty"`
	SenderUsername string `json:"sender_username,omitempty"`
	Code           string `json:"code,omitempty"`
	CodeLength     int    `json:"code_length,omitempty"`
	CallbackURL    string `json:"callback_url,omitempty"`
	Payload        string `json:"payload,omitempty"`
	TTL            int    `json:"ttl,omitempty"`
}

type CheckSendAbilityRequest struct {
	PhoneNumber string `json:"phone_number"`
}
