package esssdk

import (
	"fmt"
)

type PercentSignModel struct {
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

// PercentSignModels represents the percent signing info to an ess document.
type PercentSignModels struct {
	// 文档ID
	DocID string `json:"docId"`

	// 备注
	Remark string `json:"remark"`

	// 签署信息
	Signs []*PercentSignModel `json:"signs"`
}

// SignDocByPercentV2 signs an ess document by specified percent.
func (c *Client) SignDocByPercentV2(si *PercentSignModels) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/v2/user/doc/sign/percent", c.opts.Addr)

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
