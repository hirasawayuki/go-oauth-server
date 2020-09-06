package service

import (
	"github.com/hirasawayuki/go-oauth-server/model"
	"github.com/hirasawayuki/go-oauth-server/repository"
	"github.com/pkg/errors"
)

type IClientService interface {
	Get(client_id string) (model.Client, error)
}

type ClientService struct {
	ClientRepository repository.IClientRepository
}

func (s ClientService) Get(client_id string) (model.Client, error) {
	result, err := s.ClientRepository.Find(client_id)
	if err != nil {
		return model.Client{}, errors.Wrap(err, "")
	}
	m := model.NewClient(model.ClientDTO{
		ID:           result.ID,
		ClientID:     result.ClientID,
		ClientSecret: result.ClientSecret,
		RedirectURI:  result.RedirectURI,
		Scope:        result.Scope,
	})
	return m, nil
}
