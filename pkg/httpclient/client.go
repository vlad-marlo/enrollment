package httpclient

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

func New(addr string) (*Client, error) {
	cli := new(Client)
	cli.client = resty.New().SetBaseURL(addr)

	if resp, err := cli.client.R().Get("/ping"); err != nil || !resp.IsSuccess() {
		return nil, fmt.Errorf("bad connection: %w", err)
	}

	return cli, nil
}
