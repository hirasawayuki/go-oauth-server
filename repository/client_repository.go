package repository

import (
	"github.com/hirasawayuki/go-oauth-server/registry"
)

type Client struct {
	ID           int    `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
}

type ClientRepository struct {
	Store registry.IStores
}

type IClientRepository interface {
	Find(client_id string) (*Client, error)
}

func (r ClientRepository) Find(client_id string) (*Client, error) {
	var client Client
	client.ClientID = client_id
	err := r.Store.Default().DB.First(&client).Error
	if err != nil {
		return nil, err
	}
	return &client, nil
}
