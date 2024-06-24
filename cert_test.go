package client_test

import (
	"fmt"
	"testing"

	ec "github.com/gfacloud/gfa-ess-client-go"
)

func TestCreateUserCert(t *testing.T) {
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

	userID := "53d8c95886b4d3270186b51292610003"
	certSN, err := c.CreateUserCert(userID)
	if err != nil {
		t.Fatalf("CreateUserCert failed: %v", err)
	}
	fmt.Println("证书序列号: ", certSN)
}

func TestImportUserCert(t *testing.T) {
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

	// 读取证书文件
	certBase64, _, err := readFile("files/user.cer")
	if err != nil {
		t.Fatalf("Read cert file failed: %v", err)
	}

	userID := "ff80808185ba37e70185bdc8bd160a1d"
	certSN, err := c.ImportUserCert(userID, string(certBase64), false)
	if err != nil {
		t.Fatalf("ImportUserCert failed: %v", err)
	}
	fmt.Println("证书序列号: ", certSN)
}
