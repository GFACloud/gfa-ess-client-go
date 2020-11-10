package esssdk

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCreateSeal(t *testing.T) {
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

func readTestFile(filename string) (contentBase64 string, fileType string, err error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	contentBase64 = base64.StdEncoding.EncodeToString(content)

	items := strings.Split(filename, ".")
	if len(items) >= 2 {
		fileType = items[len(items)-1]
	} else {
		fileType = "png"
	}

	return
}
