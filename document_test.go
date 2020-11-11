package esssdk

import (
	"fmt"
	"testing"
)

func TestCreateDocument(t *testing.T) {
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

	// 读取测试文件
	contentBase64, _, err := readTestFile("files/proof_cert.pdf")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	doc := &Document{
		DocName:          "李四的劳动合同-电子数据存证证书",
		DocType:          PDFDocType,
		DocContentBase64: contentBase64,
		UserID:           "4028e5e57562ccf00175b14a8f4e000d",
	}
	err = c.CreateDocument(doc)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}
	fmt.Println("DocID: ", doc.UUID)
}
