package client

import (
	"fmt"
)

// GetToken gets access token from ess service.
func (c *Client) GetToken() (token string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/getToken", c.opts.Addr)

	var result Response
	params := map[string]string{
		"appKey":    c.opts.AppKey,
		"appSecret": c.opts.AppSecret,
	}
	resp, err := c.httpClient.R().
		SetQueryParams(params).
		SetResult(&result).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post(url)
	if err != nil {
		return "", err
	}
	if resp.IsError() {
		return "", fmt.Errorf("%v %v", resp.Status(), resp.Error())
	}
	if result.Code != 0 {
		return "", fmt.Errorf("%v %v", result.Code, result.Msg)
	}

	v, ok := result.Data.(map[string]interface{})
	if !ok {
		err = fmt.Errorf("response data invalid: %v", result.Data)
		return
	}

	tv, found := v["accessToken"]
	if !found {
		err = fmt.Errorf("response data invalid: %v", v)
		return
	}

	token, ok = tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	return token, nil
}
