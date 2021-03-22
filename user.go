package pfsense

type UserList struct {
	Data []User `json:"data"`
}

type User struct {
	Username       string `json:"name"`
	Password       string `json:"password"`
	Descr          string `json:"descr"`
	AuthorizedKeys string `json:"authorizedkeys"`
	IpsecPsk       string `json:"ipsecpsk"`
	Disabled       bool   `json:"disabled"`
	Expires        string `json:"expires"`
}

func (c *Client) GetUserList() (*UserList, error) {

	payload := map[string]string{}
	req, err := c.prepareRequest("GET", "/user", payload)

	if err != nil {
		return nil, err
	}

	res := UserList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
