package client

import (
	"fmt"
)

// CreateUserCert creates the certification for an ess user.
func (c *Client) CreateUserCert(userID string) (certSN string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/cert/apply", c.opts.Addr)

	params := map[string]string{
		"userId": userID,
	}
	data, err := c.postParams(url, params)
	if err != nil {
		return
	}

	certSN, ok := data.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", data)
		return
	}

	return
}

// ImportUserCert imports the certification for an ess user.
func (c *Client) ImportUserCert(userID string, cert string, isDefault bool) (certSN string, err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/cert/import/cert", c.opts.Addr)

	params := map[string]string{
		"userId": userID,
		"cert":   cert,
		"defaultCert": func(isDefault bool) string {
			if isDefault {
				return "true"
			} else {
				return "false"
			}
		}(isDefault),
	}
	data, err := c.postParams(url, params)
	if err != nil {
		return
	}

	certSN, ok := data.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", data)
		return
	}

	return
}
