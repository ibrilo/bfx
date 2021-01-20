package bfx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	// ErrJSONUnmarshal TOWRITE
	ErrJSONUnmarshal = errors.New("unmarshal json error")
)

func (c *Client) request(method string, url string, b []byte, attempt int) (response []byte, err error) {

	if c.Debug {
		c.logger.Printf("%10s | %s\n", method, url)
		c.logger.Printf("payload: %s\n", string(b))
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(b))
	if err != nil {
		return
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			c.logger.Println("failed to close respons body:", e)
		}
	}()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// REPLACE THIS!!!
const (
	PlatformStatusMaintenance = 0
	PlatformStatusOperative   = 1
)

// PlatformStatus TOWRITE
func (c *Client) PlatformStatus() (int, error) {
	resp, err := c.request("GET", endpointPublicPlatformStatus, nil, 0)
	if err != nil {
		return -1, err
	}

	status := []int{}

	if err := json.Unmarshal(resp, &status); err != nil {
		return -1, err
	}

	if len(status) != 1 {
		return -1, fmt.Errorf("unexpected response platform status")
	}

	return status[0], nil
}

// Tickers TOWRITE
func (c *Client) Tickers(symbols ...string) (Tickers, error) {
	resp, err := c.request("GET", endpointPublicTickers(symbols), nil, 0)
	if err != nil {
		return nil, err
	}

	var data interface{}

	if err := json.Unmarshal(resp, &data); err != nil {
		return nil, err
	}

	tickers, err := parseTickers(data)
	if err != nil {
		return nil, err
	}

	return tickers, nil
}
