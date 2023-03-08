package esssdk

import (
	"fmt"
	"testing"
)

func TestCreateSeal(t *testing.T) {
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

	// 读取印章文件
	contentBase64, fileType, err := readTestFile("files/seal.png")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	seal := &EventCertSeal{
		ImgData:  contentBase64,
		ImgType:  fileType,
		Name:     "安信签电子证据保全平台印章",
		UserID:   "ff80808185ba37e70185bdc8bd160a1d",
		SealType: "1",
		Validity: 365,
	}
	err = c.CreateEventCertSeal(seal)
	if err != nil {
		t.Fatalf("CreateSeal failed: %v", err)
	}
	fmt.Println("SealID: ", seal.UUID)
}

func TestCreatePersonalSeal(t *testing.T) {
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

	seal := &PersonalSeal{
		Name:   "张三的人名章",
		CertSN: "0186b4d326dd",
	}
	err = c.CreatePersonalSeal(seal)
	if err != nil {
		t.Fatalf("CreatePersonalSeal failed: %v", err)
	}
	fmt.Println("人名章ID: ", seal.UUID)
}
