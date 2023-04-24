package esssdk

import (
	"fmt"
)

// PositionSignModel represents the position seal info.
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

// PositionSignModels represents the position signing info.
type PositionSignModels struct {
	// 文档ID
	DocID string `json:"docId"`

	// 备注
	Remark string `json:"remark"`

	// 位置盖章属性
	Positions []*PositionSignModel `json:"positions"`
}

// SignDocByPositionV2 signs the document by the specified position signing info.
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

// KeywordSignModel represents the keyword seal info.
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

// KeywordSignModels represents the keyword signing info.
type KeywordSignModels struct {
	// 文档ID
	DocID string `json:"docId"`
	// 备注
	Remark string `json:"remark"`
	// 关键字盖章属性
	Keywords []*KeywordSignModel `json:"keywords"`
}

// SignDocByKeywordV2 signs the document by the specified keyword signing info.
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

// CrossPageSignModel represents the cross-page seal info.
type CrossPageSignModel struct {
	// 印章ID
	SealID string `json:"sealId"`
	// 骑缝章对齐方式，1：左侧对齐；2：右侧对齐
	Align int `json:"align"`
	// 盖章范围。0：全部，1：指定页
	Scope int `json:"scope"`
	// 盖章开始页，scope为1时有效
	Start int `json:"start"`
	// 盖章结束页，scope为1时有效
	End int `json:"end"`
	// 偏移量 = 以左上角为坐标原点计算的Y坐标值 / 页面高度
	Offset float64 `json:"offset,omitempty"`
	// 印章缩放比例,不设置则按原始尺寸
	Zoom float64 `json:"zoom"`
}

// CrossPageSignModels represents the cross-page signing info.
type CrossPageSignModels struct {
	// 文档ID
	DocID string `json:"docId"`
	// 备注
	Remark string `json:"remark"`
	// 骑缝盖章属性
	CrossPages []*CrossPageSignModel `json:"crosspages"`
}

// SignDocByCrossPageV2 signs the document by the specified cross-page signing info.
func (c *Client) SignDocByCrossPageV2(si *CrossPageSignModels) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/v2/user/doc/sign/crosspage", c.opts.Addr)

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

// AnnotationSignModel represents the annotation seal info.
type AnnotationSignModel struct {
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
	// 是否隐藏
	Hidden bool `json:"hidden"`
}

// AnnotationSignModels represents the annotation signing info.
type AnnotationSignModels struct {
	// 文档ID
	DocID string `json:"docId"`

	// 备注
	Remark string `json:"remark"`

	// 位置盖章属性
	Annotations []*AnnotationSignModel `json:"annotations"`
}

// SignDocByAnnotationV2 signs the document by the specified annotation signing info.
func (c *Client) SignDocByAnnotationV2(si *AnnotationSignModels) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/v2/user/doc/sign/annotation", c.opts.Addr)

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
