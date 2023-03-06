package esssdk

import (
	"fmt"
	"testing"
)

func TestCreateDocument(t *testing.T) {
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
	contentBase64, _, err := readTestFile("files/proof_cert.pdf")
	if err != nil {
		t.Fatalf("Read test file failed: %v", err)
	}

	doc := &Document{
		DocName:          "王晓辉的劳动合同-电子数据存证证书",
		DocType:          PDFDocType,
		DocContentBase64: contentBase64,
		UserID:           "ff80808185ba37e70185bdc8bd160a1d",
	}
	docURL, err := c.CreateDocument(doc)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}
	fmt.Println("文档ID: ", doc.UUID)
	fmt.Println("文档URL: ", docURL)
}

func TestSignDocumentForKeyword(t *testing.T) {
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

	si := &KeywordSignInfo{
		DocID:   "4028e5e5768840ab0176888a63fa0009",
		SealID:  "2c92e6a677389f95017738a3ab7c0006",
		Keyword: "本人声明",
		Scope:   0,
		Start:   0,
		End:     0,
		Zoom:    100,
	}
	signedDocURL, err := c.SignDocumentForKeyword(si)
	if err != nil {
		t.Fatalf("SignDocumentForKeyword failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocumentForPosition(t *testing.T) {
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

	si := &PositionSignInfo{
		DocID:      "53d8c95886b584690186b598a14c0004",
		SealID:     "ff80808185beb7db0185bebbac9a0006",
		PageNumber: "1",
		X:          330,
		Y:          240,
		Zoom:       40,
		Remark:     "王晓辉的存证证书盖章",
	}
	signedDocURL, err := c.SignDocumentForPosition(si)
	if err != nil {
		t.Fatalf("SignDocumentForPosition failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}
