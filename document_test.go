package esssdk

import (
	"fmt"
	"testing"
)

func TestCreateDoc(t *testing.T) {
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

	// 读取测试文件
	contentBase64, _, err := readFileBase64("files/proof_cert.pdf")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	doc := &Doc{
		DocName:          "王晓辉的劳动合同-电子数据存证证书",
		DocType:          PDFDocType,
		DocContentBase64: contentBase64,
		UserID:           "ff80808185ba37e70185bdc8bd160a1d",
	}
	docURL, err := c.CreateDoc(doc)
	if err != nil {
		t.Fatalf("CreateDoc failed: %v", err)
	}
	fmt.Println("文档ID: ", doc.UUID)
	fmt.Println("文档URL: ", docURL)
}
