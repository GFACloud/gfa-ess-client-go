package client_test

import (
	"fmt"
	"testing"

	ecv2 "github.com/gfacloud/gfa-ess-client-go/v2"
)

func TestSignDocByPosition(t *testing.T) {
	// 新建客户端
	opts := &ecv2.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ecv2.NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &ecv2.PositionSignModels{
		DocID:  "53d8992f87ad00d80187ad02d64f0002",
		Remark: "文档位置盖章测试",
		Positions: []*ecv2.PositionSignModel{
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

	signedDocURL, err := c.SignDocByPosition(si)
	if err != nil {
		t.Fatalf("SignDocByPosition failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByKeyword(t *testing.T) {
	// 新建客户端
	opts := &ecv2.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ecv2.NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &ecv2.KeywordSignModels{
		DocID:  "53d8992f874b903f01874b9e8b470001",
		Remark: "文档关键字盖章测试",
		Keywords: []*ecv2.KeywordSignModel{
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

	signedDocURL, err := c.SignDocByKeyword(si)
	if err != nil {
		t.Fatalf("SignDocByKeyword failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByCrossPage(t *testing.T) {
	// 新建客户端
	opts := &ecv2.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ecv2.NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &ecv2.CrossPageSignModels{
		DocID:  "53d8992f877981db01877984638d0002",
		Remark: "文档骑缝盖章测试",
		CrossPages: []*ecv2.CrossPageSignModel{
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

	signedDocURL, err := c.SignDocByCrossPage(si)
	if err != nil {
		t.Fatalf("SignDocByCrossPage failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDocByAnnotation(t *testing.T) {
	// 新建客户端
	opts := &ecv2.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ecv2.NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &ecv2.AnnotationSignModels{
		DocID:  "53d8992f87ad00d80187ad02d64f0002",
		Remark: "文档注释盖章测试",
		Annotations: []*ecv2.AnnotationSignModel{
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

	signedDocURL, err := c.SignDocByAnnotation(si)
	if err != nil {
		t.Fatalf("SignDocByAnnotation failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}

func TestSignDoc(t *testing.T) {
	// 新建客户端
	opts := &ecv2.Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := ecv2.NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &ecv2.SignModels{
		DocID:  "8ae58ded87e59b4d0187e5a036130003",
		Remark: "文档一站式盖章测试",
		Positions: []*ecv2.PositionSignModel{
			{
				SealID: "12a775fec6767a040008",
				Page:   1,
				X:      0.5,
				Y:      0.6,
				Zoom:   0.8,
			},
		},
		Keywords: []*ecv2.KeywordSignModel{
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
		CrossPages: []*ecv2.CrossPageSignModel{
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
		Annotations: []*ecv2.AnnotationSignModel{
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
