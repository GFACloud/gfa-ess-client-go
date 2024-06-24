package client

import (
	"fmt"
)

// User represents an ess user.
type User struct {
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	UserName string `json:"userName"`
	RealName string `json:"realName"`
	UnitID   string `json:"unitId"`
	UUID     string `json:"uuid"`
}

// CreateUser creates an ess user.
func (c *Client) CreateUser(user *User) (err error) {
	url := fmt.Sprintf("http://%s/ess/api/user/create", c.opts.Addr)

	params := map[string]string{
		"email":    user.Email,
		"mobile":   user.Mobile,
		"userName": user.UserName,
		"realName": user.RealName,
		"unitId":   user.UnitID,
		"uuid":     user.UUID,
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

	userID, ok := tv.(string)
	if !ok {
		err = fmt.Errorf("response data invalid: %v", tv)
		return
	}

	if userID == "" {
		err = fmt.Errorf("response data invalid: uuid is empty")
		return
	}

	user.UUID = userID

	return
}
