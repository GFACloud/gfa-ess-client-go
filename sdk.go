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

	resp, err := c.httpClient.R().
		SetBody(body).
		SetResult(&result).
		SetHeader("Content-Type", "application/json").
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

// GetToken gets access token from ess service.
func (c *Client) GetToken() (token string, err error) {
	url := fmt.Sprintf("http://%s/api/gettoken", c.opts.Addr)

	var result Response
	params := map[string]string{
		"appkey":    c.opts.AppKey,
		"appsecret": c.opts.AppSecret,
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
