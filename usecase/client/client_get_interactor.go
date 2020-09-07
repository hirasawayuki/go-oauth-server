package usecase

import (
	"github.com/hirasawayuki/go-oauth-server/repository"
)

type ClientGetInteractor struct {
	ClientRepository repository.IClientRepository
}

func (i ClientGetInteractor) Get(input ClientGetInputDTO) (ClientGetOutputDTO, error) {
	clientID := input.ClientID
	c, err := i.ClientRepository.Find(clientID)
	if err != nil {
		return ClientGetOutputDTO{}, err
	}
	return ClientGetOutput(c.ID, c.ClientID, c.ClientSecret, c.RedirectURI, c.Scope), nil
}
