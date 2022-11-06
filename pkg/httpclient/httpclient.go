package httpclient

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	client http.Client
}

func NewClient() *Client {
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	return &Client{client: httpClient}
}

func (c *Client) NewRequest(url string) (string, error) {
	response, err := c.client.Get(url)
	if err != nil {
		return "", nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	return string(body), nil
}
