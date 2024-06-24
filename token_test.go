package client_test

import (
	"fmt"
	"testing"

	ec "github.com/gfacloud/gfa-ess-client-go"
)

func TestGetToken(t *testing.T) {
	// 新建客户端
	opts := &ec.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ec.NewClient(opts)
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
