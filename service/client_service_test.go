package service

import (
	"reflect"
	"testing"

	"github.com/hirasawayuki/go-oauth-server/model"
	"github.com/hirasawayuki/go-oauth-server/repository"
)

func createClient() *repository.Client {
	return &repository.Client{
		ID:           1,
		ClientID:     "test-client-1",
		ClientSecret: "test-client-secret",
		RedirectURI:  "http://example.com",
		Scope:        "test",
	}
}

type inMemoryClientRepository struct {
}

func (i inMemoryClientRepository) Find(clientID string) (*repository.Client, error) {
	c := createClient()
	if c.ClientID == clientID {
		return c, nil
	}
	return &repository.Client{}, nil
}

func TestGet(t *testing.T) {
	service := &ClientService{ClientRepository: inMemoryClientRepository{}}
	client, err := service.Get("test-client-1")
	if err != nil {
		t.Errorf("Error occured, got %v", err)
	}
	dto := model.ClientDTO{ID: 1, ClientID: "test-client-1", ClientSecret: "test-client-secret", RedirectURI: "http://example.com", Scope: "test"}
	expect := model.NewClient(dto)
	if !reflect.DeepEqual(client, expect) {
		t.Errorf("client does not match %v", client)
	}
}
