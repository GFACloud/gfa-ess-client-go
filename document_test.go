package esssdk

import (
	"fmt"
	"testing"
)

func TestCreateDocument(t *testing.T) {
	//t.SkipNow()

	// 新建客户端
	opts := &Options{
		AppKey:    "2c92e6a6777afd8501777b3a10440004",
		AppSecret: "f6db396833cdd152c862d88a567c7c5e2686daac",
		Addr:      "211.88.18.10:8080",
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
		UserID:           "2c92e6a67780b91e017780cd44150310",
	}
	err = c.CreateDocument(doc)
	if err != nil {
		t.Fatalf("CreateDocument failed: %v", err)
	}
	fmt.Println("文档ID: ", doc.UUID)
}

//
//func TestSignDocumentForKeyword(t *testing.T) {
//	t.SkipNow()
//
//	// 新建客户端
//	opts := &Options{
//		AppKey:    "4028e5e5765587b801765593f8940003",
//		AppSecret: "94073bf0a7d94c4f15a58e7077edaa9d21eacd9c",
//		Addr:      "211.88.18.140:30080",
//	}
//	c, err := NewClient(opts)
//	if err != nil {
//		t.Fatalf("NewClient failed: %v", err)
//	}
//
//	si := &KeywordSignInfo{
//		DocID:   "4028e5e5768840ab0176888a63fa0009",
//		SealID:  "2c92e6a677389f95017738a3ab7c0006",
//		Keyword: "本人声明",
//		Scope:   0,
//		Start:   0,
//		End:     0,
//		Zoom:    100,
//	}
//	signedDocURL, err := c.SignDocumentForKeyword(si)
//	if err != nil {
//		t.Fatalf("SignDocumentForKeyword failed: %v", err)
//	}
//
//	fmt.Println("签名文档路径: ", signedDocURL)
//}

func TestSignDocumentForPosition(t *testing.T) {
	//t.SkipNow()

	// 新建客户端
	opts := &Options{
		AppKey:    "2c92e6a6777afd8501777b3a10440004",
		AppSecret: "f6db396833cdd152c862d88a567c7c5e2686daac",
		Addr:      "211.88.18.10:8080",
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &PositionSignInfo{
		DocID:      "2c92e6a67784510c0177845a5aab0008",
		SealID:     "2c92e6a67784433c017784455c3d0002",
		PageNumber: "1",
		X:          330,
		Y:          240,
		Zoom:       40,
		Remark:     "李四的存证证书盖章",
	}
	signedDocURL, err := c.SignDocumentForPosition(si)
	if err != nil {
		t.Fatalf("SignDocumentForPosition failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}
