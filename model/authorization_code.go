package model

import (
	"time"
)

type AuthorizationCode struct {
	id          int
	code        string
	userID      int
	clientID    string
	scope       string
	redirectURI string
	expiresAt   int64
}

const AuthorizationCodeDuration = 600

// 有効期限は 10 分間
func expiresAt() int64 {
	return time.Now().Add(time.Second * time.Duration(AuthorizationCodeDuration)).Unix()
}

func NewAuhtorizationCode(id int, code string, userID int, clientID string, scope string, expiresAt int64) AuthorizationCode {
	return AuthorizationCode{
		id:        id,
		userID:    userID,
		clientID:  clientID,
		scope:     scope,
		expiresAt: expiresAt,
	}
}

func (code AuthorizationCode) ID() int {
	return code.id
}

func (code AuthorizationCode) ClientID() string {
	return code.clientID
}

func (code AuthorizationCode) ClientSecret() string {
	return code.scope
}

func (code AuthorizationCode) RedirectURI() string {
	return code.redirectURI
}

func (code AuthorizationCode) ExpiresAt() int64 {
	return code.expiresAt
}
