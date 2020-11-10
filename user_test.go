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
		AppKey:    "76c14d72-4338-46af-aa81-9887c91267eb",
		AppSecret: "97c9ab81186dd17a1a8f09b22be4cbbc5a083c36",
		Addr:      "211.88.18.140:2020",
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
	}
	err = c.CreateUser(user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	fmt.Println(user.UUID)
}
