package esssdk

import (
	"fmt"
)

// KeywordSignInfo represents the keyword signing info to an ess document.
type KeywordSignInfo struct {
	DocID   string `json:"docId"`
	SealID  string `json:"sealId"`
	Keyword string `json:"keyword"`
	Scope   int    `json:"scope"`
	Start   int    `json:"start"`
	End     int    `json:"end"`
	Zoom    int    `json:"zoom"`
}

// SignDocByKeyword signs an ess document by the specified keyword.
func (c *Client) SignDocByKeyword(si *KeywordSignInfo) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/doc/sign/keyword", c.opts.Addr)

	params := map[string]string{
		"docId":   si.DocID,
		"sealId":  si.SealID,
		"keyword": si.Keyword,
		"scope":   fmt.Sprintf("%v", si.Scope),
		"start":   fmt.Sprintf("%v", si.Start),
		"end":     fmt.Sprintf("%v", si.End),
		"zoom":    fmt.Sprintf("%v", si.Zoom),
	}
	data, err := c.postParams(url, params)
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

// PositionSignInfo represents the position signing info to an ess document.
type PositionSignInfo struct {
	DocID      string `json:"docId"`
	SealID     string `json:"sealId"`
	PageNumber string `json:"pageNumber"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Zoom       int    `json:"zoom"`
	Reason     string `json:"reason"`
	Remark     string `json:"remark"`
}

// SignDocByPosition signs an ess document by specified postions.
func (c *Client) SignDocByPosition(si *PositionSignInfo) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/doc/sign/position", c.opts.Addr)

	params := map[string]string{
		"docId":               si.DocID,
		"remark":              si.Remark,
		"signs[0].pageNumber": si.PageNumber,
		"signs[0].sealId":     si.SealID,
		"signs[0].x":          fmt.Sprintf("%v", si.X),
		"signs[0].y":          fmt.Sprintf("%v", si.Y),
		"signs[0].zoom":       fmt.Sprintf("%v", si.Zoom),
		"signs[0].reason":     si.Reason,
	}
	data, err := c.postParams(url, params)
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

// PercentSignInfo represents the percent signing info to an ess document.
type PercentSignInfo struct {
	// 文档ID
	DocID string `json:"docId"`
	// 印章ID
	SealID string `json:"sealId"`
	// 页码范围,设置签章所在页码,默认“0”，0：所有页，-1：倒数最后一页，
	//  多个使用“,”分页,例如1-11代表从第1页到第11页,"1,3,5"代表签署第1、3、5页
	PageNumber string `json:"pageNumber"`
	// x轴坐标百分比,示例值(0.2)
	X float64 `json:"x"`
	// y轴坐标百分比,示例值(0.3)
	Y float64 `json:"y"`
	// 印章缩放比例,取值范围百分比：0-1,如果该值为0则使用width和height值,示例值(0.2)
	Zoom float64 `json:"zoom"`
	// 签署原因
	Reason string `json:"reason"`
	// 备注
	Remark string `json:"remark"`
}

// SignDocByPercent signs an ess document by specified percent.
func (c *Client) SignDocByPercent(si *PercentSignInfo) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/doc/sign/percent", c.opts.Addr)

	params := map[string]string{
		"docId":               si.DocID,
		"remark":              si.Remark,
		"signs[0].pageNumber": si.PageNumber,
		"signs[0].sealId":     si.SealID,
		"signs[0].x":          fmt.Sprintf("%f", si.X),
		"signs[0].y":          fmt.Sprintf("%f", si.Y),
		"signs[0].zoom":       fmt.Sprintf("%f", si.Zoom),
		"signs[0].reason":     si.Reason,
	}
	data, err := c.postParams(url, params)
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

// CrossPageSignInfo represents the cross-page signing info to an ess document.
type CrossPageSignInfo struct {
	// 待签文档ID
	DocID string `json:"docId"`
	// 签署印章ID
	SealID string `json:"sealId"`
	// 起始页：从1开始，1表示第一页
	BeginPage int `json:"beginPage"`
	// 结束页：从0开始，0表示最末页
	EndPage int `json:"endPage"`
	// 位置标识，1：'左侧齐缝章'，2：'右侧齐缝章'
	PosType int `json:"posType"`
	// 盖章纵向位置：0-49999，按比列确定位置，24999表示正中央
	Position int `json:"position"`
	// 印章缩放比例,取值范围百分比：10-100,示例值(100)
	Zoom int `json:"zoom"`
	// 签署原因,示例值(人事合同员工签名)
	Remark string `json:"remark"`
}

// SignDocByCrossPage signs the document by the cross-page seal.
func (c *Client) SignDocByCrossPage(si *CrossPageSignInfo) (signedDocURL string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/doc/sign/across/page", c.opts.Addr)

	params := map[string]string{
		"docId":     si.DocID,
		"sealId":    si.SealID,
		"beginPage": fmt.Sprintf("%v", si.BeginPage),
		"endPage":   fmt.Sprintf("%v", si.EndPage),
		"posType":   fmt.Sprintf("%v", si.PosType),
		"position":  fmt.Sprintf("%v", si.Position),
		"zoom":      fmt.Sprintf("%v", si.Zoom),
		"remark":    si.Remark,
	}
	data, err := c.postParams(url, params)
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
