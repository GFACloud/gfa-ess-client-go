package esssdk

import (
	"fmt"
)

type PositionSignModel struct {
	// 印章ID
	SealID string `json:"sealId"`
	// 页码范围,设置签章所在页码,默认“0”，0：第一页，-1：倒数最后一页
	Page int `json:"page"`
	// x轴坐标百分比,示例值(0.2)
	X float64 `json:"x"`
	// y轴坐标百分比,示例值(0.3)
	Y float64 `json:"y"`
	// 印章缩放比例,取值范围百分比：0-1,如果该值为0则使用width和height值,示例值(0.2)
	Zoom float64 `json:"zoom"`
	// 签署原因
	Reason string `json:"reason"`
}

// PositionSignModels represents the position signing info to an ess document.
type PositionSignModels struct {
	// 文档ID
	DocID string `json:"docId"`

	// 备注
	Remark string `json:"remark"`

	// 签署信息
	Signs []*PositionSignModel `json:"signs"`
}

// SignDocByPositionV2 signs an ess document by specified position.
func (c *Client) SignDocByPositionV2(si *PositionSignModels) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/v2/user/doc/sign/position", c.opts.Addr)

	data, err := c.postObject(url, si)
	if err != nil {
		return
	}

	v, ok := data.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("response data invalid: %v", data)
		return
	}

	tv, found := v["id"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	docID, ok := tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	if docID != si.DocID {
		err = fmt.Errorf("response data invalid: id is unmatched")
		return
	}

	tv, found = v["url"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	signedDocURL, ok = tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	return
}

type KeywordSignModel struct {
	// 印章ID
	SealID string `json:"sealId"`
	// 盖章关键字
	Keyword string `json:"keyword"`
	// 关键字查找范围。0：全部，1：指定页
	Scope int `json:"scope"`
	// 关键字查找开始页，scope为1时有效
	Start int `json:"start"`
	// 关键字查找结束页，scope为1时有效
	End int `json:"end"`
	// 印章缩放比例,取值范围：0-1
	Zoom float64 `json:"zoom"`
	// 左右偏移量，正值表示向右偏移，负值表示向左偏移
	OffsetX float64 `json:"offsetX"`
	// 上下偏移量，正值表示向下偏移，负值表示向上偏移
	OffsetY float64 `json:"offsetY"`
}

// KeywordSignModels represents the keyword signing info to an ess document.
type KeywordSignModels struct {
	// 文档ID
	DocID string `json:"docId"`
	// 备注
	Remark string `json:"remark"`
	// 签署信息
	Signs []*KeywordSignModel `json:"signs"`
}

// SignDocByKeywordV2 signs an ess document by specified keywords.
func (c *Client) SignDocByKeywordV2(si *KeywordSignModels) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/v2/user/doc/sign/keyword", c.opts.Addr)

	data, err := c.postObject(url, si)
	if err != nil {
		return
	}

	v, ok := data.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("response data invalid: %v", data)
		return
	}

	tv, found := v["id"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	docID, ok := tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	if docID != si.DocID {
		err = fmt.Errorf("response data invalid: id is unmatched")
		return
	}

	tv, found = v["url"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	signedDocURL, ok = tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	return
}
