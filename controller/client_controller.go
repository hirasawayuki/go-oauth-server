package controller

import (
	"github.com/hirasawayuki/go-oauth-server/usecase"
	"github.com/pkg/errors"
)

type ClientController struct {
	ClientUsecase usecase.IClientUsecase
}
type Client struct {
	ID           int    `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
}

type GetClientResponseDTO struct {
	Client Client
	Error  error
}

type IClientController interface {
	Get(clientID string) GetClientResponseDTO
}

func (c ClientController) Get(clientID string) GetClientResponseDTO {
	result := Client{}
	r, err := c.ClientUsecase.Get(clientID)
	if err != nil {
		return GetClientResponseDTO{Client: result, Error: errors.Wrap(err, "")}
	}
	return GetClientResponseDTO{
		Client: Client{ID: r.ID, ClientID: r.ClientID, ClientSecret: r.ClientSecret, RedirectURI: r.RedirectURI, Scope: r.Scope},
		Error:  nil,
	}
}
