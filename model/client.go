package model

type Client struct {
	id            int    `json:"id"`
	client_id     string `json:"client_id"`
	client_secret string `json:"client_secret"`
	redirect_uri  string `json:"redirect_uri"`
	scope         string `json:"scope"`
}

type ClientDTO struct {
	ID           int
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}

func NewClient(c ClientDTO) Client {
	return Client{
		id:            c.ID,
		client_id:     c.ClientID,
		client_secret: c.ClientSecret,
		redirect_uri:  c.RedirectURI,
		scope:         c.Scope,
	}
}

func (c Client) ID() int {
	return c.id
}

func (c Client) ClientID() string {
	return c.client_id
}

func (c Client) ClientSecret() string {
	return c.client_secret
}

func (c Client) RedirectURI() string {
	return c.redirect_uri
}

func (c Client) Scope() string {
	return c.scope
}
