package telegramgateway

// SendVerificationMessageRequest holds the parameters of
// Client.SendVerificationMessage. Only PhoneNumber is required; omitted
// fields get the server-side defaults.
type SendVerificationMessageRequest struct {
	// PhoneNumber is the recipient's phone number in E.164 format,
	// for example "+998901234567".
	PhoneNumber string `json:"phone_number"`
	// RequestID is the identifier returned by a previous CheckSendAbility
	// call. Sending with it is not charged a second time.
	RequestID string `json:"request_id,omitempty"`
	// SenderUsername is the username of the Telegram channel the code is
	// sent from.
	SenderUsername string `json:"sender_username,omitempty"`
	// Code is the verification code. Leave it empty and set CodeLength to let
	// Telegram generate the code.
	Code string `json:"code,omitempty"`
	// CodeLength is the length of the code Telegram generates when Code is
	// empty.
	CodeLength int `json:"code_length,omitempty"`
	// CallbackURL is the address that receives delivery reports.
	CallbackURL string `json:"callback_url,omitempty"`
	// Payload is custom data returned in delivery reports and request
	// statuses. It is not shown to the user.
	Payload string `json:"payload,omitempty"`
	// TTL is the time to live of the message in seconds.
	TTL int `json:"ttl,omitempty"`
}

// CheckSendAbilityRequest holds the parameters of Client.CheckSendAbility.
type CheckSendAbilityRequest struct {
	// PhoneNumber is the phone number to check, in E.164 format.
	PhoneNumber string `json:"phone_number"`
}

// CheckVerificationStatusRequest holds the parameters of
// Client.CheckVerificationStatus.
type CheckVerificationStatusRequest struct {
	PhoneNumber string `json:"phone_number"`
	// Code is the code entered by the user. If set, the API checks it
	// against the request.
	Code string `json:"code,omitempty"`
}

// RevokeVerificationMessageRequest holds the parameters of
// Client.RevokeVerificationMessage.
type RevokeVerificationMessageRequest struct {
	// RequestID is the identifier of the verification request to revoke.
	RequestID string `json:"request_id"`
}
