package client_test

import (
	"fmt"
	"testing"

	ec "github.com/gfacloud/gfa-ess-client-go"
)

func TestCreateDocPdf(t *testing.T) {
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

	// 读取测试文件
	contentBase64, _, err := readFileBase64("files/test.pdf")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	doc := &ec.Doc{
		DocName:          "PDF文档测试",
		DocType:          ec.PDFDocType,
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

func TestCreateDocOfd(t *testing.T) {
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

	// 读取测试文件
	contentBase64, _, err := readFileBase64("files/test.ofd")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	doc := &ec.Doc{
		DocName:          "OFD文档测试",
		DocType:          ec.OFDDocType,
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
