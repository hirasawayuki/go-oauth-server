package repository

import (
	"github.com/hirasawayuki/go-oauth-server/model"
	"github.com/hirasawayuki/go-oauth-server/registry"
)

type IClientRepository interface {
	Find(client_id string) (model.Client, error)
}

type Client struct {
	ID           int    `gorm:"column:id"`
	ClientID     string `gorm:"column:client_id"`
	ClientSecret string `gorm:"column:client_secret"`
	RedirectURI  string `gorm:"column:redirect_uri"`
	Scope        string `gorm:"column:scope"`
}

type ClientRepository struct {
	Store registry.IStores
}

func (r ClientRepository) Find(client_id string) (model.Client, error) {
	var client Client
	client.ClientID = client_id
	err := r.Store.Default().DB.First(&client).Error
	if err != nil {
		return model.Client{}, err
	}
	return model.NewClient(client.ID, client.ClientID, client.ClientSecret, client.RedirectURI, client.Scope), nil
}
