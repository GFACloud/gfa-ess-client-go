package client

import (
	"fmt"
)

// EventCertSeal represents an ess event cert seal.
type EventCertSeal struct {
	ImgData  string `json:"imgData"`
	ImgType  string `json:"imgType"`
	Name     string `json:"name"`
	UserID   string `json:"userId"`
	SealType string `json:"sealType"`
	UUID     string `json:"uuid"`
	Validity int    `json:"validity"`
}

// CreateEventCertSeal creates an ess event cert seal.
func (c *Client) CreateEventCertSeal(seal *EventCertSeal) (err error) {
	url := fmt.Sprintf("http://%s/ess/api/seal/event/cert/apply", c.opts.Addr)

	params := map[string]string{
		"imgData":  seal.ImgData,
		"imgType":  seal.ImgType,
		"name":     seal.Name,
		"userId":   seal.UserID,
		"sealType": seal.SealType,
		"uuid":     seal.UUID,
		"validity": fmt.Sprintf("%v", seal.Validity),
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

	tv, found := v["uuid"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	sealID, ok := tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	if sealID == "" {
		err = fmt.Errorf("response data invalid: uuid is empty")
		return
	}

	seal.UUID = sealID

	return
}

// UserNameSeal represents an ess user name seal.
type UserNameSeal struct {
	Name   string `json:"name"`
	CertSN string `json:"certSN"`
	UUID   string `json:"uuid"`
}

// CreateUserNameSeal creates an ess user name seal.
func (c *Client) CreateUserNameSeal(seal *UserNameSeal) (err error) {
	url := fmt.Sprintf("http://%s/ess/api/seal/user_name/cert/apply", c.opts.Addr)

	params := map[string]string{
		"name":   seal.Name,
		"certSN": seal.CertSN,
		"uuid":   seal.UUID,
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

	tv, found := v["uuid"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	sealID, ok := tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	if sealID == "" {
		err = fmt.Errorf("response data invalid: uuid is empty")
		return
	}

	seal.UUID = sealID

	return
}

// UserImageSeal represents an ess user image seal.
type UserImageSeal struct {
	// 印章ID,非必填,会生成默认印章ID,如果填写需要保证唯一性
	UUID string `json:"uuid"`
	// 印章名称
	Name string `json:"name"`
	// 证书序列号
	CertSN string `json:"certSN"`
	// 印章图片
	Image string `json:"image"`
}

// CreateUserImageSeal creates an ess user image seal.
func (c *Client) CreateUserImageSeal(seal *UserImageSeal) (err error) {
	url := fmt.Sprintf("http://%s/ess/api/seal/user_image/cert/apply", c.opts.Addr)

	params := map[string]string{
		"uuid":   seal.UUID,
		"name":   seal.Name,
		"certSN": seal.CertSN,
		"image":  seal.Image,
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

	tv, found := v["uuid"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	sealID, ok := tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	if sealID == "" {
		err = fmt.Errorf("response data invalid: uuid is empty")
		return
	}

	seal.UUID = sealID

	return
}
