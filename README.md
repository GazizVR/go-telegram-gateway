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
		var apiErr *telegramgateway.APIError
		if errors.As(err, &apiErr) {
			log.Fatalf("api error: %s", apiErr.Message)
		}
		log.Fatal(err)
	}
	fmt.Println("request id:", status.RequestId)
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
	telegramgateway.WithHttpClient(&http.Client{Timeout: 10 * time.Second}),
	telegramgateway.WithBaseURL("http://localhost:8080"), // e.g. for tests
)
```

The default HTTP client has no timeout, so set one or pass a context with a deadline.

## Errors

- `*APIError`: the API rejected the request (`StatusCode`, `Message`). Use `errors.As`.
- `ErrNetwork`: the request could not be sent or the response not read. Check
  for `context.Canceled` / `context.DeadlineExceeded` before retrying.
- `ErrResponseParsing`: the API answered successfully but the response could
  not be decoded. The request may have been processed, so do not blindly retry.

## Testing

```sh
go test ./...
```

## License

[MIT](LICENSE)
