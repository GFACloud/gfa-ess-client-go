package esssdk

import (
	"fmt"
)

// Seal represents an ess seal.
type Seal struct {
	ImgData string `json:"imgData"`
	ImgType string `json:"imgType"`
	Name    string `json:"name"`
	UserID  string `json:"userId"`
	UUID    string `json:"uuid"`
}

// CreateSeal creates an ess seal.
func (c *Client) CreateSeal(seal *Seal) (err error) {
	url := fmt.Sprintf("http://%s/api/seal/quick/create", c.opts.Addr)

	params := map[string]string{
		"imgData": seal.ImgData,
		"imgType": seal.ImgType,
		"name":    seal.Name,
		"userId":  seal.UserID,
		"uuid":    seal.UUID,
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
