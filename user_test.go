package client_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	ec "github.com/gfacloud/gfa-ess-client-go"
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

	userName := randString(10)
	email := fmt.Sprintf("%s@ec.com.cn", userName)

	user := &ec.User{
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
