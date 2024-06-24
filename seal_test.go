package client_test

import (
	"fmt"
	"testing"

	ec "github.com/gfacloud/gfa-ess-client-go"
)

func TestCreateEventCertSeal(t *testing.T) {
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

	// 读取印章文件
	contentBase64, fileType, err := readFileBase64("files/org.png")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	seal := &ec.EventCertSeal{
		ImgData:  contentBase64,
		ImgType:  fileType,
		Name:     "安信签电子证据保全平台印章",
		UserID:   "ff80808185ba37e70185bdc8bd160a1d",
		SealType: "1",
		Validity: 365,
	}
	err = c.CreateEventCertSeal(seal)
	if err != nil {
		t.Fatalf("TestCreateEventCertSeal failed: %v", err)
	}
	fmt.Println("SealID: ", seal.UUID)
}

func TestCreateUserNameSeal(t *testing.T) {
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

	seal := &ec.UserNameSeal{
		Name:   "张三的人名章",
		CertSN: "0186b4d326dd",
	}
	err = c.CreateUserNameSeal(seal)
	if err != nil {
		t.Fatalf("TestCreateUserNameSeal failed: %v", err)
	}
	fmt.Println("SealID: ", seal.UUID)
}

func TestCreateUserImageSeal(t *testing.T) {
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

	// 读取印章文件
	imgBase64, _, err := readFileBase64("files/person.png")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	seal := &ec.UserImageSeal{
		Name:   "张三的图片章",
		CertSN: "0186b4c2388b",
		Image:  imgBase64,
	}
	err = c.CreateUserImageSeal(seal)
	if err != nil {
		t.Fatalf("CreateUserImageSeal failed: %v", err)
	}
	fmt.Println("SealID: ", seal.UUID)
}
