package telegramgateway

import (
	"testing"
)

func TestParseReponse(t *testing.T) {
	resps := []struct {
		name       string
		statusCode int
		body       string
		wantError  bool
	}{
		{"normal", 200, `{"ok": true}`, false},
		{"internal-err", 502, "", true},
		{"empty", 200, "", true},
	}
	for _, resp := range resps {
		t.Run(resp.name, func(t *testing.T) {
			var result string
			err := parseResponse(resp.statusCode, []byte(resp.body), result)
			if resp.wantError {
				if err == nil {
					t.Fatal("expected err to not be nil")
				}
			} else {
				if err != nil {
					t.Fatalf("response parsing error: %v", err)
				}
			}
		})
	}
}
