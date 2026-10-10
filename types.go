package telegramgateway

import (
	"encoding/json"
	"time"
)

// RequestStatus describes the state of a verification request.
// Fields documented as nil-able are absent from the API response in some
// cases, so check them before dereferencing.
type RequestStatus struct {
	RequestID   string `json:"request_id"`
	PhoneNumber string `json:"phone_number"`
	// RequestCost is the total cost of the request in credits.
	RequestCost float64 `json:"request_cost"`
	// IsRefunded is true if the request fee was refunded.
	IsRefunded bool `json:"is_refunded"`
	// RemainingBalance is the remaining balance in credits. It is nil unless
	// the request incurred a charge.
	RemainingBalance *float64 `json:"remaining_balance"`
	// DeliveryStatus is the current delivery status. It is nil unless a
	// verification message was sent.
	DeliveryStatus *DeliveryStatus `json:"delivery_status"`
	// VerificationStatus is the current verification status. It is nil unless
	// a code has been checked.
	VerificationStatus *VerificationStatus `json:"verification_status"`
	// Payload is the custom data passed in the request, if any.
	Payload *string `json:"payload"`
}

// UnixTime is a time.Time that is encoded in JSON as a Unix timestamp in
// seconds, the format used by the Gateway API. Decoded values are in UTC.
type UnixTime struct {
	time.Time
}

// MarshalJSON encodes t as a Unix timestamp in seconds.
func (t UnixTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

// UnmarshalJSON decodes a Unix timestamp in seconds. A JSON null leaves t
// unchanged.
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

// DeliveryState is the delivery state of a verification message.
// The API may add new values, so do not assume the list below is exhaustive.
type DeliveryState string

const (
	DeliverySent      DeliveryState = "sent"
	DeliveryDelivered DeliveryState = "delivered"
	DeliveryRead      DeliveryState = "read"
	DeliveryExpired   DeliveryState = "expired"
	DeliveryRevoked   DeliveryState = "revoked"
)

// DeliveryStatus is the delivery status of a verification message.
type DeliveryStatus struct {
	Status DeliveryState `json:"status"`
	// UpdatedAt is the time the status was last updated.
	UpdatedAt UnixTime `json:"updated_at"`
}

// VerificationState is the result of checking a code.
// The API may add new values, so do not assume the list below is exhaustive.
type VerificationState string

const (
	VerificationCodeValid   VerificationState = "code_valid"
	VerificationCodeInvalid VerificationState = "code_invalid"
	VerificationMaxAttempts VerificationState = "code_max_attempts_exceeded"
	VerificationExpired     VerificationState = "expired"
)

// VerificationStatus is the verification status of a request.
type VerificationStatus struct {
	Status VerificationState `json:"status"`
	// UpdatedAt is the time the status was last updated.
	UpdatedAt UnixTime `json:"updated_at"`
	// CodeEntered is the code entered by the user. It is nil if no code was
	// checked.
	CodeEntered *string `json:"code_entered"`
}
