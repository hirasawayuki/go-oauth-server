package usecase

import (
	"github.com/hirasawayuki/go-oauth-server/model"
	"github.com/hirasawayuki/go-oauth-server/repository"
)

type IClientGetUseCase interface {
	Get(ClientGetInputDTO) (ClientGetOutputDTO, error)
}

type ClientGetInputDTO struct {
	ClientID string
}

func NewClientGetInputDTO(client_id string) ClientGetInputDTO {
	return ClientGetInputDTO{
		ClientID: client_id,
	}
}

type ClientGetOutputDTO struct {
	ID           int
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}

func NewClientGetOutputDTO(c model.Client) ClientGetOutputDTO {
	return ClientGetOutputDTO{
		ID:           c.ID(),
		ClientID:     c.ClientID(),
		ClientSecret: c.ClientSecret(),
		RedirectURI:  c.RedirectURI(),
		Scope:        c.Scope(),
	}
}

// ClientGetUsecase
type ClientGetUseCase struct {
	ClientRepository repository.IClientRepository
}

func (i ClientGetUseCase) Get(input ClientGetInputDTO) (ClientGetOutputDTO, error) {
	clientID := input.ClientID
	client, err := i.ClientRepository.Find(clientID)
	if err != nil {
		return ClientGetOutputDTO{}, err
	}
	return NewClientGetOutputDTO(client), nil
}
