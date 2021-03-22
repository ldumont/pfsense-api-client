package pfsense

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const (
	BaseURLV1 = "/api/v1"
)

type Client struct {
	BaseURL     string
	clientId    string
	clientToken string
	HTTPClient  *http.Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Return  int         `json:"return"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewClient(host string, clientId string, clientToken string) *Client {
	return &Client{

		BaseURL:     host + BaseURLV1,
		clientId:    clientId,
		clientToken: clientToken,
		HTTPClient: &http.Client{

			Timeout: time.Minute,
		},
	}
}

func (c *Client) injectCredentials(v *map[string]string) error {
	(*v)["client-id"] = c.clientId
	(*v)["client-token"] = c.clientToken
	return nil
}

func (c *Client) prepareRequest(method string, endpoint string, payload map[string]string) (*http.Request, error) {
	c.injectCredentials(&payload)

	final_payload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(final_payload))
	return req, err
}

func (c *Client) sendRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Content-Type", "application/json")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := successResponse{
		Data: v,
	}

	// debugging
	// if resDump, err := httputil.DumpResponse(res, true); err != nil {
	// 	fmt.Printf(err.Error())
	// } else {
	// 	fmt.Printf("%s\n\n", resDump)
	// }

	if err = json.NewDecoder(res.Body).Decode(&fullResponse.Data); err != nil {
		return err
	}

	return nil
}
