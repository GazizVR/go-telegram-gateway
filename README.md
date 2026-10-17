# go-telegram-gateway

[![CI](https://github.com/gazizvr/go-telegram-gateway/actions/workflows/ci.yml/badge.svg)](https://github.com/gazizvr/go-telegram-gateway/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/gazizvr/go-telegram-gateway.svg)](https://pkg.go.dev/github.com/gazizvr/go-telegram-gateway)

Unofficial Go client for the [Telegram Gateway API](https://core.telegram.org/gateway/api):
send and verify one-time codes through Telegram. Standard library only, no dependencies.

## Install

```sh
go get github.com/gazizvr/go-telegram-gateway
```

You need an API token from your Telegram Gateway account settings.

## Usage

```go
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	telegramgateway "github.com/gazizvr/go-telegram-gateway"
)

func main() {
	client, err := telegramgateway.NewClient("YOUR_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	status, err := client.SendVerificationMessage(ctx, telegramgateway.SendVerificationMessageRequest{
		PhoneNumber: "+998901234567",
		CodeLength:  6,
	})
	if err != nil {
		if errors.Is(err, telegramgateway.ErrPhoneNumberInvalid) {
			log.Fatal("invalid phone number")
		}
		var floodErr *telegramgateway.FloodError
		if errors.As(err, &floodErr) {
			log.Fatalf("flood wait: %v", floodErr.WaitTime)
		}
		var statusErr *telegramgateway.UnknownStatusError
		if errors.As(err, &statusErr) {
			log.Fatalf("unexpected http status: %d", statusErr.StatusCode)
		}
		log.Fatal(err)
	}
	fmt.Println("request id:", status.RequestID)
}
```

## Methods

| Method | API method |
|---|---|
| `SendVerificationMessage` | `sendVerificationMessage` |
| `CheckSendAbility` | `checkSendAbility` |
| `CheckVerificationStatus` | `checkVerificationStatus` |
| `RevokeVerificationMessage` | `revokeVerificationMessage` |

All methods take a `context.Context` as the first argument.

## Options

```go
client, err := telegramgateway.NewClient(token,
	telegramgateway.WithHTTPClient(&http.Client{Timeout: 10 * time.Second}),
	telegramgateway.WithBaseURL("http://localhost:8080"), // e.g. for tests
)
```

The default HTTP client has no timeout, so set one or pass a context with a deadline.

## Errors

The client returns typed errors that can be inspected with `errors.Is` and `errors.As`:

- **Sentinel errors**: Known API errors are mapped directly to typed sentinel errors such as `ErrPhoneNumberInvalid`, `ErrCodeInvalid`, `ErrBalanceNotEnough`, etc. Check with `errors.Is`.
- `*FloodError`: The request was throttled by rate limits. Contains `WaitTime time.Duration` indicating how long to wait before retrying. Use `errors.As`.
- `*UnknownAPIError`: The API rejected the request with an unrecognized error code or returned `"ok": false` without an error message (`Code string`). Use `errors.As`.
- `*UnknownStatusError`: The server answered with a non-2xx HTTP status code (`StatusCode int`). This includes non-JSON error responses from proxies or upstream servers. Use `errors.As`.
- **Decode errors**: If the server answers with a 2xx status code but the response envelope or `result` cannot be decoded as JSON, a wrapped decoding error is returned (`decode response: ...` or `decode result: ...`).

## Testing

```sh
go test ./...
```

## License

[MIT](LICENSE)
