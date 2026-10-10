package telegramgateway

import (
    "testing"
    "time"
)

const testToken = "secret_valid_token"

func TestClient_Default(t *testing.T) {
	client, err := NewClient(testToken)
	if err != nil {
		t.Fatalf("unexpected error: %v",err)
	}
	if client == nil {
		t.Fatal("expected client to not be nil")
	}

	if client.token != testToken {
		t.Errorf("token: got: %q, want: %q", client.token, testToken)
	}
	if client.baseURL != defaultBaseURL {
		t.Errorf("baseURL: got: %q, want: %q", client.baseURL, defaultBaseURL)
	}
	if client.httpClient == nil {
		t.Error("expected client.httpClient to not be nil")
	}
}

func TestClient_EmptyToken(t *testing.T) {
	token := ""
	client, err := NewClient(token)
	if err == nil {
		t.Fatal("expected err to not be nil")
	}
	if client != nil {
		t.Fatalf("client: got: %v, want: nil", client)
	}
}

func TestClient_BlankBaseURL(t *testing.T) {
	baseURL := ""
	option := WithBaseURL(baseURL)
	client, err := NewClient(testToken, option)
	if err == nil {
		t.Fatal("expected err to not be nil")
	}
	if client != nil {
		t.Fatalf("client: got: %v, want: nil", client)
	}
}

func TestClient_InvalidBaseURL(t *testing.T) {
    baseURL := "htpph:/localos:81246/method"
	option := WithBaseURL(baseURL)
	client, err := NewClient(testToken, option)
	if err == nil {
		t.Fatal("expected err to not be nil")
	}
	if client != nil {
		t.Fatalf("client: got: %v, want: nil", client)
	}
}

func TestClient_HttpClientNil(t *testing.T) {
	option := WithHttpClient(nil)
	client, err := NewClient(testToken, option)
	if err == nil {
		t.Fatal("expected err to not be nil")
	}
	if client != nil {
		t.Fatalf("client: got: %v, want: nil", client)
	}
}

func TestClient_CustomHttpClient(t *testing.T) {
    httpClient := &http.Client{
        Timeout: 30 * time.Second,
    }
    option := WithHttpClient(httpClient)
	client, err := NewClient(testToken, option)
	if err != nil {
		t.Fatalf("unexpected error: %v",err)
	}
	if client == nil {
		t.Fatal("expected client to not be nil")
    }
    if client.httpClient != httpClient {
        t.Fatalf("httpClient pointer mismatch: got: %p, want: %p",client.httpClient,httpClient)
    }
}