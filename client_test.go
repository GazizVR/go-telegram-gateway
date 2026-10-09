package telegramgateway

import "testing"

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

func TestClient_Default(t *testing.T) {
	token := "secret_valid_token"
	client, err := NewClient(token)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if client == nil {
		t.Fatal("expected client to not be nil")
	}

	if client.token != token {
		t.Errorf("token: got: %q, want: %q", client.token, token)
	}
	if client.baseURL != defaultBaseURL {
		t.Errorf("baseURL: got: %q, want: %q", client.baseURL, defaultBaseURL)
	}
	if client.httpClient == nil {
		t.Error("expected client.httpClient to not be nil")
	}
}

func TestClient_BlankBaseURL(t *testing.T) {
	token := "secret_valid_token"
	baseURL := ""
	option := WithBaseURL(baseURL)
	client, err := NewClient(token, option)
	if err == nil {
		t.Fatal("expected err to not be nil")
	}
	if client != nil {
		t.Fatalf("client: got: %v, want: nil", client)
	}
}
