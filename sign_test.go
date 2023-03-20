package esssdk

import (
	"fmt"
	"testing"
)

func TestSignDocByKeyword(t *testing.T) {
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
	signedDocURL, err := c.SignDocByKeyword(si)
	if err != nil {
		t.Fatalf("SignDocByKeyword failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByPosition(t *testing.T) {
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
		DocID:      "53d8c95886b584690186b59770240003",
		SealID:     "da97e6bbbf2884870001",
		PageNumber: "1",
		X:          330,
		Y:          240,
		Zoom:       40,
		Remark:     "王晓辉的存证证书盖章",
	}
	signedDocURL, err := c.SignDocByPosition(si)
	if err != nil {
		t.Fatalf("SignDocByPosition failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByPercent(t *testing.T) {
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

	si := &PercentSignInfo{
		DocID:      "53d8c95886b584690186b59770240003",
		SealID:     "da97e6bbbf2884870001",
		PageNumber: "1",
		X:          0.5,
		Y:          0.6,
		Zoom:       0.5,
		Remark:     "王晓辉的存证证书盖章",
		Reason:     "存证",
	}
	signedDocURL, err := c.SignDocByPercent(si)
	if err != nil {
		t.Fatalf("SignDocByPercent failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByCrossPage(t *testing.T) {
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

	si := &CrossPageSignInfo{
		DocID:     "53d8992f86fdb6ac0186fdbd18150001",
		SealID:    "11d3c4efee39360a0001",
		BeginPage: 3,
		EndPage:   13,
		PosType:   1,
		Position:  24999,
		Zoom:      100,
		Remark:    "租房合同盖章",
	}
	signedDocURL, err := c.SignDocByCrossPage(si)
	if err != nil {
		t.Fatalf("SignDocByCrossPage failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}
