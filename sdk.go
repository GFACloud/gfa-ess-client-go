package esssdk

import (
	"fmt"

	resty "github.com/go-resty/resty/v2"
)

// Response represents the response of http request.
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Options represents the client options.
type Options struct {
	// 接入方编码
	AppKey string

	// 接入方密钥
	AppSecret string

	// 电子印章服务地址
	Addr string
}

// Client represents GFA-ESS client.
type Client struct {
	opts       *Options
	httpClient *resty.Client
}

// NewClient creates an instance of Client.
func NewClient(opts *Options) (c *Client, err error) {
	err = checkOptions(opts)
	if err != nil {
		return nil, err
	}

	// 初始化Http客户端
	httpClient := resty.New()

	c = &Client{opts: opts, httpClient: httpClient}
	return c, nil
}

func checkOptions(opts *Options) error {
	if opts.AppKey == "" {
		return fmt.Errorf("接入方编码为空")
	}
	if opts.AppSecret == "" {
		return fmt.Errorf("接入方密钥为空")
	}
	if opts.Addr == "" {
		return fmt.Errorf("服务地址为空")
	}
	return nil
}

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
