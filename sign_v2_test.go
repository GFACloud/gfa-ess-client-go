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
		DocID:  "53d8c95886b584690186b58c4b820002",
		Remark: "王晓辉的存证证书盖章",
		Signs: []*PercentSignModel{
			{
				SealID: "ff80808185beb7db0185bebbac9a0006",
				Page:   1,
				X:      0.5,
				Y:      0.6,
				Zoom:   0.5,
				Reason: "存证",
			},
		},
	}

	signedDocURL, err := c.SignDocByPercentV2(si)
	if err != nil {
		t.Fatalf("SignDocByPercentV2 failed: %v", err)
	}

	fmt.Println("签名文档路径: ", signedDocURL)
}
