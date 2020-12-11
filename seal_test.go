package esssdk

import (
	"fmt"
	"testing"
)

func TestCreateSeal(t *testing.T) {
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

	// 读取印章文件
	contentBase64, fileType, err := readTestFile("files/seal.png")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	seal := &Seal{
		ImgData: contentBase64,
		ImgType: fileType,
		Name:    "安信签电子证据保全平台印章",
		UserID:  "4028e5e57562ccf00175b14a8f4e000d",
	}
	err = c.CreateSeal(seal)
	if err != nil {
		t.Fatalf("CreateSeal failed: %v", err)
	}
	fmt.Println("SealID: ", seal.UUID)
}

func TestCreatePersonalSeal(t *testing.T) {
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

	seal := &PersonalSeal{
		Name:   "张三的人名章",
		CertSN: "1765f4525a5",
	}
	err = c.CreatePersonalSeal(seal)
	if err != nil {
		t.Fatalf("CreatePersonalSeal failed: %v", err)
	}
	fmt.Println("人名章ID: ", seal.UUID)
}
