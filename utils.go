package client

import (
	"fmt"
	"io"
)

func (c *Client) postObject(url string, body interface{}) (interface{}, error) {
	var result Response

	token, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.R().
		SetBody(body).
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
		SetHeader("token_access", token).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%v %v", resp.Status(), resp.Error())
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("%v %v", result.Code, result.Msg)
	}

	return result.Data, nil
}

func (c *Client) postParams(url string, params map[string]string) (interface{}, error) {
	var result Response

	token, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.R().
		SetQueryParams(params).
		SetResult(&result).
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("token_access", token).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%v %v", resp.Status(), resp.Error())
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("%v %v", result.Code, result.Msg)
	}

	return result.Data, nil
}

func (c *Client) postFileForm(url string, formData map[string]string, fileParam string, fileName string, fileReader io.Reader) (interface{}, error) {
	var result Response

	token, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.R().
		SetFileReader(fileParam, fileName, fileReader).
		SetFormData(formData).
		SetResult(&result).
		SetHeader("token_access", token).
		Post(url)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, fmt.Errorf("%v %v", resp.Status(), resp.Error())
	}
	if result.Code != 0 {
		return nil, fmt.Errorf("%v %v", result.Code, result.Msg)
	}

	return result.Data, nil
}
