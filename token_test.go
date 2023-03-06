package esssdk

import (
	"fmt"
	"testing"
)

func TestGetToken(t *testing.T) {
	// 新建客户端
	opts := &Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	token, err := c.GetToken()
	if err != nil {
		t.Fatalf("GetToken failed: %v", err)
	}

	if token == "" {
		t.Fatalf("token should not be empty")
	}

	fmt.Println(token)
}
