package telegramgateway

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrCodeInvalid             = errors.New("code invalid")
	ErrCodeExpired             = errors.New("code expired")
	ErrCodeLengthRequired      = errors.New("code length required")
	ErrCodeLengthInvalid       = errors.New("code length invalid")
	ErrCodeMaxAttemptsExceeded = errors.New("code max attempts exceeded")
	ErrPhoneNumberInvalid      = errors.New("phone number invalid")
	ErrPhoneNumberMismatch     = errors.New("phone number mismatch")
	ErrRequestIDInvalid        = errors.New("request id invalid")
	ErrRequestIDRequired       = errors.New("request id required")
	ErrPayloadInvalid          = errors.New("payload invalid")
	ErrSenderUsernameInvalid   = errors.New("sender username invalid")
	ErrSenderNotVerified       = errors.New("sender not verified")
	ErrSenderNotOwned          = errors.New("sender not owned")
	ErrCallbackURLInvalid      = errors.New("callback URL invalid")
	ErrTTLInvalid              = errors.New("TTL invalid")
	ErrAccessTokenInvalid      = errors.New("access token invalid")
	ErrAccessTokenRequired     = errors.New("access token required")
	ErrMessageAlreadySent      = errors.New("message already sent")
	ErrBalanceNotEnough        = errors.New("balance not enough")
	ErrUnknownMethod           = errors.New("unknown method")
)

type FloodError struct {
	WaitTime time.Duration
}

func (e *FloodError) Error() string {
	suffix := e.WaitTime.String()
	return "flood wait " + suffix
}

type UnknownAPIError struct {
	Code string
}

func (e *UnknownAPIError) Error() string {
	return "unknown api error: " + e.Code
}

var errByCode = map[string]error{
	"CODE_INVALID":               ErrCodeInvalid,
	"CODE_EXPIRED":               ErrCodeExpired,
	"CODE_LENGTH_REQUIRED":       ErrCodeLengthRequired,
	"CODE_LENGTH_INVALID":        ErrCodeLengthInvalid,
	"CODE_MAX_ATTEMPTS_EXCEEDED": ErrCodeMaxAttemptsExceeded,
	"PHONE_NUMBER_INVALID":       ErrPhoneNumberInvalid,
	"PHONE_NUMBER_MISMATCH":      ErrPhoneNumberMismatch,
	"REQUEST_ID_INVALID":         ErrRequestIDInvalid,
	"REQUEST_ID_REQUIRED":        ErrRequestIDRequired,
	"PAYLOAD_INVALID":            ErrPayloadInvalid,
	"SENDER_USERNAME_INVALID":    ErrSenderUsernameInvalid,
	"SENDER_NOT_VERIFIED":        ErrSenderNotVerified,
	"SENDER_NOT_OWNED":           ErrSenderNotOwned,
	"CALLBACK_URL_INVALID":       ErrCallbackURLInvalid,
	"TTL_INVALID":                ErrTTLInvalid,
	"ACCESS_TOKEN_INVALID":       ErrAccessTokenInvalid,
	"ACCESS_TOKEN_REQUIRED":      ErrAccessTokenRequired,
	"MESSAGE_ALREADY_SENT":       ErrMessageAlreadySent,
	"BALANCE_NOT_ENOUGH":         ErrBalanceNotEnough,
	"UNKNOWN_METHOD":             ErrUnknownMethod,
}

func parseError(code string) error {
	if err, ok := errByCode[code]; ok {
		return err
	}
	if strings.HasPrefix(code, "FLOOD_WAIT") {
		var waitDur time.Duration
		numStr := strings.TrimPrefix(code, "FLOOD_WAIT_")
		if num, err := strconv.Atoi(numStr); err == nil {
			waitDur = time.Duration(num) * time.Second
		}
		return &FloodError{WaitTime: waitDur}
	}
	return &UnknownAPIError{Code: code}
}

type UnknownStatusError struct {
	StatusCode int
}

func (e *UnknownStatusError) Error() string {
	return fmt.Sprintf("unexpected http status %d", e.StatusCode)
}
