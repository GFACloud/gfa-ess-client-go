package esssdk

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func randString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func TestCreateUser(t *testing.T) {
	t.SkipNow()

	// 新建客户端
	opts := &Options{
		AppKey:    "4028e5e5765587b801765593f8940003",
		AppSecret: "94073bf0a7d94c4f15a58e7077edaa9d21eacd9c",
		Addr:      "211.88.18.140:30080",
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	userName := randString(10)
	email := fmt.Sprintf("%s@ec.com.cn", userName)

	user := &User{
		Email:    email,
		Mobile:   "12345678901",
		UserName: userName,
		RealName: "张三",
	}
	err = c.CreateUser(user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	fmt.Println("用户ID: ", user.UUID)
}

func TestCreateUserCert(t *testing.T) {
	t.SkipNow()

	// 新建客户端
	opts := &Options{
		AppKey:    "4028e5e5765587b801765593f8940003",
		AppSecret: "94073bf0a7d94c4f15a58e7077edaa9d21eacd9c",
		Addr:      "211.88.18.140:30080",
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	userID := "4028e5e5765f43630176655cf3c5003c"
	certSN, err := c.CreateUserCert(userID)
	if err != nil {
		t.Fatalf("CreateUserCert failed: %v", err)
	}
	fmt.Println("证书序列号: ", certSN)
}
