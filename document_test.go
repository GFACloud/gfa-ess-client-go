package esssdk

import (
	"fmt"
	"testing"
)

func TestCreateDocument(t *testing.T) {
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

	// 读取测试文件
	contentBase64, _, err := readTestFile("files/demo.pdf")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	doc := &Document{
		DocName:          "李四的劳动合同-电子数据存证证书",
		DocType:          PDFDocType,
		DocContentBase64: contentBase64,
		UserID:           "4028e5e5765f43630176655cf3c5003c",
	}
	err = c.CreateDocument(doc)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}
	fmt.Println("文档ID: ", doc.UUID)
}

func TestSignDocument(t *testing.T) {
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

	si := &SignInfo{
		DocID:   "4028e5e5765f4363017665834bec003f",
		SealID:  "4028e5e5765f43630176656067f1003d",
		Keyword: "本人声明",
		Scope:   0,
		Start:   0,
		End:     0,
		Zoom:    100,
	}
	signedDocURL, err := c.SignDocument(si)
	if err != nil {
		t.Fatalf("SignDocument failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}
