package esssdk

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	opts := &Options{
		AppKey:    "4028e5e5765587b801765593f8940003",
		AppSecret: "94073bf0a7d94c4f15a58e7077edaa9d21eacd9c",
		Addr:      "211.88.18.140:30080",
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	if c == nil {
		t.Fatalf("client should not be nil")
	}
}
