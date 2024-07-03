package client_test

import (
	"fmt"
	"testing"

	ecv2 "github.com/gfacloud/gfa-ess-client-go/v2"
)

func TestGetToken(t *testing.T) {
	// 新建客户端
	opts := &ecv2.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ecv2.NewClient(opts)
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
