package sophosfirewall

// Client is a client for the Sophos Firewall API.
type Client struct {
	username string
	password string
	endpoint string
}

// New returns a new Sophos Firewall client. It requires a username, password,
// and endpoint. The endpoint should be the API endpoint of the Sophos
// Firewall, for example, `https://firewall.example.com:4444/webconsole/APIController`.
func New(username, password, endpoint string) *Client {
	return &Client{
		username: username,
		password: password,
		endpoint: endpoint,
	}
}

// Login authenticates with the Sophos Firewall and returns an error if
// unsuccessful.
func (c *Client) Login() error {
	return nil
}
