package colormeshop

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
)

const (
	defaultBaseURL = "https://api.shop-pro.jp"
	TokenEnvVar    = "COLORME_TOKEN"
)

type Client struct {
	URL        *url.URL
	HTTPClient *http.Client
	Token      string
}

func NewClient() (*Client, error) {
	u, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		URL:        u,
		HTTPClient: http.DefaultClient,
	}

	token := os.Getenv(TokenEnvVar)
	if token != "" {
		c.Token = token
	}

	return c, nil
}

func (c *Client) get(p string) (*http.Response, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, p)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	return c.HTTPClient.Do(req)
}

func decodeJSON(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
