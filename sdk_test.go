package esssdk

import (
	"testing"
)

const (
	// For 145
	APP_KEY_TEST    = "ff80808185a616240185aa352dc409a6"
	APP_SECRET_TEST = "2d291657812b0c56a8a197c307f1dfa2116f501d"
	ADDR_TEST       = "localhost:8080"

	// For 109
	// APP_KEY_TEST    = "8ae58ded77c7a5a90177c7a63d260002"
	// APP_SECRET_TEST = "e09ea9ae4927005e747763e014a78fe84d75024c"
	// ADDR_TEST       = "10.101.13.109:8800"
)

func TestNewClient(t *testing.T) {
	opts := &Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	if c == nil {
		t.Fatalf("client should not be nil")
	}
}
