package esssdk

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	// 新建客户端
	opts := &Options{
		AppKey:    "76c14d72-4338-46af-aa81-9887c91267eb",
		AppSecret: "97c9ab81186dd17a1a8f09b22be4cbbc5a083c36",
		Addr:      "211.88.18.140:2020",
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	token, err := c.GetToken()
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}

	fmt.Println(token)
}
