package usecase

type IClientGetUseCase interface {
	Get(ClientGetInputDTO) ClientGetOutputDTO
}

type ClientGetInputDTO struct {
	ClientID string
}

type ClientGetOutputDTO struct {
	ID           int
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}

type ClientUsecaseOutput struct {
	ID           int
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}

func ClientGetOutput(id int, clientID string, clientSecret string, redirectURI string, scope string) ClientGetOutputDTO {
	return ClientGetOutputDTO{
		ID:           id,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		Scope:        scope,
	}
}
