package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	httpClient HttpClient
}

func NewClient(opts ...Option) *Client {
	o := Options{
		HttpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return &Client{
		httpClient: o.HttpClient,
	}
}

func (c *Client) JSON(ctx context.Context, url string, v interface{}) error {
	before := func(req *http.Request) {
		req.Header.Set("Accept", "application/json")
	}

	after := func(res *http.Response) error {
		if err := json.NewDecoder(res.Body).Decode(v); err != nil {
			return fmt.Errorf(`%w: %v`, ErrParsingResponse, err)
		}

		return nil
	}

	return c.execute(ctx, url, before, after)
}

func (c *Client) execute(
	ctx context.Context,
	url string,
	before func(*http.Request),
	after func(*http.Response) error,
) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf(`%w: %v`, ErrMakingRequest, err)
	}

	before(req)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf(`%w: %v`, ErrExecutingRequest, err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(`%w %d`, ErrUnexpectedResponseStatusCode, res.StatusCode)
	}

	return after(res)
}
