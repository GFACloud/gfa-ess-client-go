package esssdk

import (
	"fmt"
	"testing"
)

func TestSignDocByPercentV2(t *testing.T) {
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

	si := &PercentSignModels{
		DocID:  "53d8992f8736c311018736d90ee00003",
		Remark: "OFD文档位置盖章测试",
		Signs: []*PercentSignModel{
			{
				SealID: "ff80808185beb7db0185bebbac9a0006",
				Page:   1,
				X:      0.5,
				Y:      0.6,
				Zoom:   0.5,
				Reason: "OFD文档位置盖章测试",
			},
		},
	}

	signedDocURL, err := c.SignDocByPercentV2(si)
	if err != nil {
		t.Fatalf("SignDocByPercentV2 failed: %v", err)
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
		Signs: []*KeywordSignModel{
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
