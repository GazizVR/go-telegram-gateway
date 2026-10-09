package telegramgateway

import (
	"encoding/json"
	"time"
)

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

type UnixTime struct {
	time.Time
}

func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

func (t *UnixTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var sec int64
	if err := json.Unmarshal(data, &sec); err != nil {
		return err
	}
	t.Time = time.Unix(sec, 0).UTC()
	return nil
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
	UpdatedAt UnixTime      `json:"updated_at"`
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
	UpdatedAt   UnixTime          `json:"updated_at"`
	CodeEntered *string           `json:"code_entered"`
}
