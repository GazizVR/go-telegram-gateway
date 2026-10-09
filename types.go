package telegramgateway

type RequestStatus struct {
	RequestId   string  `json:"request_id"`
	PhoneNumber string  `json:"phone_number"`
	RequestCost float64 `json:"request_cost"`
}
