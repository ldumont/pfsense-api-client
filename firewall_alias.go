package pfsense

type FirewallAliasList struct {
	Data []FirewallAlias `json:"data"`
}

type FirewallAlias struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Descr   string `json:"descr"`
	Detail  string `json:"detail"`
}

func (c *Client) CreateFirewallAlias(v FirewallAlias) (*FirewallAlias, error) {

	payload := map[string]string{
		"name":    v.Name,
		"type":    v.Type,
		"address": v.Address,
		"descr":   v.Descr,
		"detail":  v.Detail,
	}
	req, err := c.prepareRequest("POST", "/firewall/alias", payload)

	if err != nil {
		return nil, err
	}

	res := FirewallAlias{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteFirewallAlias(v FirewallAlias) (*FirewallAlias, error) {

	payload := map[string]string{
		"id": v.Name,
	}
	req, err := c.prepareRequest("DELETE", "/firewall/alias", payload)

	if err != nil {
		return nil, err
	}

	res := FirewallAlias{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetFirewallAliasList() (*FirewallAliasList, error) {

	payload := map[string]string{}

	req, err := c.prepareRequest("GET", "/firewall/alias", payload)

	if err != nil {
		return nil, err
	}

	res := FirewallAliasList{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateFirewallAlias(v FirewallAlias) (*FirewallAlias, error) {

	payload := map[string]string{
		"id":      v.Name,
		"name":    v.Name,
		"type":    v.Type,
		"address": v.Address,
		"descr":   v.Descr,
		"detail":  v.Detail,
	}
	req, err := c.prepareRequest("PUT", "/firewall/alias", payload)

	if err != nil {
		return nil, err
	}

	res := FirewallAlias{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
