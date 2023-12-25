package sophosfirewall

type Client struct {
	username string
	password string
	endpoint string
}

func NewClient(username, password, endpoint string) *Client {
	return &Client{
		username: username,
		password: password,
		endpoint: endpoint,
	}
}

func (c *Client) Login() error {
	return nil
}
