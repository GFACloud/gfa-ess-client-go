package esssdk

import (
	"fmt"
	"testing"
)

func TestSignDocByPositionV2(t *testing.T) {
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

	si := &PositionSignModels{
		DocID:  "53d8992f87ad00d80187ad02d64f0002",
		Remark: "文档位置盖章测试",
		Positions: []*PositionSignModel{
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
	opts := &Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &KeywordSignModels{
		DocID:  "53d8992f874b903f01874b9e8b470001",
		Remark: "OFD文档关键字盖章测试",
		Keywords: []*KeywordSignModel{
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
	opts := &Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &CrossPageSignModels{
		DocID:  "53d8992f877981db01877984638d0002",
		Remark: "文档骑缝盖章测试",
		CrossPages: []*CrossPageSignModel{
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
	opts := &Options{
		AppKey:    APP_KEY_TEST,
		AppSecret: APP_SECRET_TEST,
		Addr:      ADDR_TEST,
	}
	c, err := NewClient(opts)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	si := &AnnotationSignModels{
		DocID:  "53d8992f87ad00d80187ad02d64f0002",
		Remark: "文档注释盖章测试",
		Annotations: []*AnnotationSignModel{
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
