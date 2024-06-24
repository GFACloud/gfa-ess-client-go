package client_test

import (
	"fmt"
	"testing"

	ec "github.com/gfacloud/gfa-ess-client-go"
)

func TestSignDocByPositionV2(t *testing.T) {
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

	si := &ec.PositionSignModels{
		DocID:  "53d8992f87ad00d80187ad02d64f0002",
		Remark: "文档位置盖章测试",
		Positions: []*ec.PositionSignModel{
			{
				SealID: "db4bea687357a1b60003",
				Page:   1,
				X:      0.5,
				Y:      0.6,
				Zoom:   0.5,
				Reason: "文档位置盖章测试",
			},
		},
	}

	signedDocURL, err := c.SignDocByPositionV2(si)
	if err != nil {
		t.Fatalf("SignDocByPositionV2 failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByKeywordV2(t *testing.T) {
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

	si := &ec.KeywordSignModels{
		DocID:  "53d8992f874b903f01874b9e8b470001",
		Remark: "文档关键字盖章测试",
		Keywords: []*ec.KeywordSignModel{
			{
				SealID:  "db09f65231610d7a0002",
				Keyword: "长恨歌",
				Zoom:    0.8,
				OffsetX: 10,
			},
			{
				SealID:  "db09f65231610d7a0002",
				Keyword: "长歌行",
				Zoom:    0.5,
				OffsetY: 5,
			},
			{
				SealID:  "db09f65231610d7a0002",
				Keyword: "警世贤文",
				Zoom:    0.3,
			},
		},
	}

	signedDocURL, err := c.SignDocByKeywordV2(si)
	if err != nil {
		t.Fatalf("SignDocByKeywordV2 failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByCrossPageV2(t *testing.T) {
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

	si := &ec.CrossPageSignModels{
		DocID:  "53d8992f877981db01877984638d0002",
		Remark: "文档骑缝盖章测试",
		CrossPages: []*ec.CrossPageSignModel{
			{
				// SealID: "db4bea6873582a6d0004",
				SealID: "db4bea687357a1b60003",
				Align:  1,
				Scope:  1,
				Start:  2,
				End:    10,
				Offset: 0.2,
				Zoom:   0.5,
			},
		},
	}

	signedDocURL, err := c.SignDocByCrossPageV2(si)
	if err != nil {
		t.Fatalf("SignDocByCrossPageV2 failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByAnnotationV2(t *testing.T) {
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

	si := &ec.AnnotationSignModels{
		DocID:  "53d8992f87ad00d80187ad02d64f0002",
		Remark: "文档注释盖章测试",
		Annotations: []*ec.AnnotationSignModel{
			{
				SealID: "db4bea6873582a6d0004",
				Page:   1,
				X:      0.5,
				Y:      0.6,
				Zoom:   0.5,
				Reason: "文档注释盖章测试",
				Hidden: true,
			},
		},
	}

	signedDocURL, err := c.SignDocByAnnotationV2(si)
	if err != nil {
		t.Fatalf("SignDocByAnnotationV2 failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDoc(t *testing.T) {
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

	si := &ec.SignModels{
		DocID:  "8ae58ded87e59b4d0187e5a036130003",
		Remark: "文档一站式盖章测试",
		Positions: []*ec.PositionSignModel{
			{
				SealID: "12a775fec6767a040008",
				Page:   1,
				X:      0.5,
				Y:      0.6,
				Zoom:   0.8,
			},
		},
		Keywords: []*ec.KeywordSignModel{
			{
				SealID:  "12a775fec20227070003",
				Keyword: "长恨歌",
				Zoom:    0.8,
				OffsetX: 10,
			},
			{
				SealID:  "12a775fec20227070003",
				Keyword: "长歌行",
				Zoom:    0.5,
				OffsetY: 5,
			},
			{
				SealID:  "12a775fec20227070003",
				Keyword: "警世贤文",
				Zoom:    0.3,
			},
		},
		CrossPages: []*ec.CrossPageSignModel{
			{
				// SealID: "db4bea6873582a6d0004",
				SealID: "12a775fec6767a040008",
				Align:  2,
				Scope:  1,
				Start:  2,
				End:    10,
				Offset: 0.2,
				Zoom:   0.5,
			},
		},
		Annotations: []*ec.AnnotationSignModel{
			{
				SealID: "db4bea6873582a6d0004",
				Page:   2,
				X:      0.5,
				Y:      0.6,
				Zoom:   0.8,
				Hidden: false,
			},
		},
	}

	signedDocURL, err := c.SignDoc(si)
	if err != nil {
		t.Fatalf("SignDoc failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}
