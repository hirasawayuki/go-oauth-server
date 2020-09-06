package handler

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/hirasawayuki/go-oauth-server/controller"
)

type AuthorizeHandler struct {
	ClientController controller.IClientController
}

func (h AuthorizeHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state == "" {
		log.Fatal("state is not exists")
	}
	clientID := r.URL.Query().Get("client_id")
	v := h.ClientController.Get(clientID)
	redirect_uri := r.URL.Query().Get("redirect_uri")
	if v.Client.RedirectURI != redirect_uri {
		err := errors.New("mismatched redirect URI")
		log.Fatal(err)
	}
	scope := r.URL.Query().Get("scope")
	rs := strings.Split(scope, " ")
	cs := strings.Split(v.Client.Scope, " ")
	if !reflect.DeepEqual(rs, cs) {
		err := errors.New("invalid scope")
		log.Fatal(err)
	}
	u, err := url.Parse(redirect_uri)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()
	q.Set("state", state)
	q.Set("code", "code")
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), 301)
}
