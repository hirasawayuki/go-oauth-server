package controller

import (
	"github.com/hirasawayuki/go-oauth-server/usecase"
	"github.com/pkg/errors"
)

type ClientController struct {
	ClientGetUseCase usecase.IClientGetUseCase
}

type IClientController interface {
	Get(clientID string) GetClientResponse
}

type Client struct {
	ID           int
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}
type GetClientResponse struct {
	Client Client
	Error  error
}

func (c ClientController) Get(clientID string) GetClientResponse {
	input := usecase.NewClientGetInputDTO(clientID)
	client, err := c.ClientGetUseCase.Get(input)
	if err != nil {
		return GetClientResponse{Client: Client{}, Error: errors.Wrap(err, "")}
	}
	return GetClientResponse{
		Client: Client{ID: client.ID, ClientID: client.ClientID, ClientSecret: client.ClientSecret, RedirectURI: client.RedirectURI, Scope: client.Scope},
		Error:  errors.Wrap(err, ""),
	}
}
