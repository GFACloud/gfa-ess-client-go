package esssdk

import (
	"fmt"
)

// EventCertSeal represents an ess seal.
type EventCertSeal struct {
	ImgData  string `json:"imgData"`
	ImgType  string `json:"imgType"`
	Name     string `json:"name"`
	UserID   string `json:"userId"`
	SealType string `json:"sealType"`
	UUID     string `json:"uuid"`
	Validity int    `json:"validity"`
}

// CreateEventCertSeal creates an ess seal.
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

// PersonalSeal represents an ess personal seal.
type PersonalSeal struct {
	Name   string `json:"name"`
	CertSN string `json:"certSN"`
	UUID   string `json:"uuid"`
}

// CreatePersonalSeal creates an ess seal.
func (c *Client) CreatePersonalSeal(seal *PersonalSeal) (err error) {
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
