package usecase

import (
	"github.com/hirasawayuki/go-oauth-server/service"
	"github.com/pkg/errors"
)

type ClientUsecase struct {
	ClientService service.IClientService
}

type ClientUsecaseOutput struct {
	ID           int
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}

type IClientUsecase interface {
	Get(clientID string) (ClientUsecaseOutput, error)
}

type Client struct {
	ID           int    `json:"id"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
}

func (u ClientUsecase) Get(clientID string) (ClientUsecaseOutput, error) {
	result, err := u.ClientService.Get(clientID)
	if err != nil {
		return ClientUsecaseOutput{}, errors.Wrap(err, "")
	}
	return ClientUsecaseOutput{
		ID:           result.ID(),
		ClientID:     result.ClientID(),
		ClientSecret: result.ClientSecret(),
		RedirectURI:  result.RedirectURI(),
		Scope:        result.Scope(),
	}, nil
}
