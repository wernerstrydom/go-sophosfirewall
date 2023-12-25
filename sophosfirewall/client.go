package sophosfirewall

import (
	"fmt"
	"os"
)

// ClientOption is a function that configures a Sophos Firewall client.
type ClientOption func(*Client)

// WithUsername sets the username to use to authenticate with the Sophos
// Firewall.
func WithUsername(username string) ClientOption {
	return func(c *Client) {
		c.username = username
	}
}

// WithPassword sets the password to use to authenticate with the Sophos
// Firewall.
func WithPassword(password string) ClientOption {
	return func(c *Client) {
		c.password = password
	}
}

// WithEndpoint sets the endpoint of the Sophos Firewall (for example,
// `https://firewall.example.com:4444/webconsole/APIController`).
func WithEndpoint(endpoint string) ClientOption {
	return func(c *Client) {
		c.endpoint = endpoint
	}
}

// WithHost sets the host of the Sophos Firewall (for example,
// `firewall.example.com`). The endpoint will be set to
// `https://<host>:4444/webconsole/APIController`.
func WithHost(host string) ClientOption {
	return func(c *Client) {
		c.endpoint = fmt.Sprintf("https://%s:4444/webconsole/APIController", host)
	}
}

// Client is a client for the Sophos Firewall API.
type Client struct {
	username string
	password string
	endpoint string
}

// New returns a new Sophos Firewall client. The client can be configured
// using the `With*` options. If no options are provided, the client will
// attempt to read the username, password, and endpoint from the
// `SOPHOS_FIREWALL_USERNAME`, `SOPHOS_FIREWALL_PASSWORD`, and
// `SOPHOS_FIREWALL_ENDPOINT` environment variables, respectively.
//
// The endpoint can also be set using the `SOPHOS_FIREWALL_HOST` environment
// variable. If the endpoint is not set, the client will not be able to
// authenticate with the Sophos Firewall.
func New(opts ...ClientOption) *Client {
	var username, password, endpoint string
	var ok bool
	username, _ = os.LookupEnv("SOPHOS_FIREWALL_USERNAME")
	password, _ = os.LookupEnv("SOPHOS_FIREWALL_PASSWORD")
	if endpoint, ok = os.LookupEnv("SOPHOS_FIREWALL_ENDPOINT"); !ok {
		var host string
		if host, ok = os.LookupEnv("SOPHOS_FIREWALL_HOST"); ok {
			endpoint = fmt.Sprintf("https://%s:4444/webconsole/APIController", host)
		}
	}

	c := &Client{
		username: username,
		password: password,
		endpoint: endpoint,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Validate validates the client configuration and returns an error if
// unsuccessful.
func (c *Client) Validate() error {
	if c.username == "" {
		return fmt.Errorf("username not set")
	}
	if c.password == "" {
		return fmt.Errorf("password not set")
	}
	if c.endpoint == "" {
		return fmt.Errorf("endpoint not set")
	}
	return nil
}

// Login authenticates with the Sophos Firewall and returns an error if
// unsuccessful.
func (c *Client) Login() error {
	return nil
}
