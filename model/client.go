package model

type Client struct {
	id           int
	clientID     string
	clientSecret string
	redirectURI  string
	scope        string
}

func NewClient(id int, clientID string, clientSecret string, redirectURI string, scope string) Client {
	return Client{
		id:           id,
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
		scope:        scope,
	}
}

func (c Client) ID() int {
	return c.id
}

func (c Client) ClientID() string {
	return c.clientID
}

func (c Client) ClientSecret() string {
	return c.clientSecret
}

func (c Client) RedirectURI() string {
	return c.redirectURI
}

func (c Client) Scope() string {
	return c.scope
}
