package telegramgateway

type RequestStatus struct {
	RequestId          string              `json:"request_id"`
	PhoneNumber        string              `json:"phone_number"`
	RequestCost        float64             `json:"request_cost"`
	IsRefunded         bool                `json:"is_refunded"`
	RemainingBalance   *float64            `json:"remaining_balance"`
	DeliveryStatus     *DeliveryStatus     `json:"delivery_status"`
	VerificationStatus *VerificationStatus `json:"verification_status"`
	Payload            *string             `json:"payload"`
}

type DeliveryState string

const (
	DeliverySent      DeliveryState = "sent"
	DeliveryDelivered DeliveryState = "delivered"
	DeliveryRead      DeliveryState = "read"
	DeliveryExpired   DeliveryState = "expired"
	DeliveryRevoked   DeliveryState = "revoked"
)

type DeliveryStatus struct {
	Status    DeliveryState `json:"status"`
	UpdatedAt int64         `json:"updated_at"`
}

type VerificationState string

const (
	VerificationCodeValid   VerificationState = "code_valid"
	VerificationCodeInvalid VerificationState = "code_invalid"
	VerificationMaxAttempts VerificationState = "code_max_attempts_exceeded"
	VerificationExpired     VerificationState = "expired"
)

type VerificationStatus struct {
	Status      VerificationState `json:"status"`
	UpdatedAt   int64             `json:"updated_at"`
	CodeEntered *string           `json:"code_entered"`
}
