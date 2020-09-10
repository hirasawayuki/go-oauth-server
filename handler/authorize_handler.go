package handler

import (
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
		http.Error(w, "state is not exists.", 400)
		return
	}
	clientID := r.URL.Query().Get("client_id")
	v := h.ClientController.Get(clientID)
	redirectURI := r.URL.Query().Get("redirect_uri")
	if v.Client.RedirectURI != redirectURI {
		http.Error(w, "mismatched redirect URI.", 400)
		return
	}
	scope := r.URL.Query().Get("scope")
	rs := strings.Split(scope, " ")
	cs := strings.Split(v.Client.Scope, " ")
	if !reflect.DeepEqual(rs, cs) {
		log.Printf("rs: %v, cs: %v", rs, cs)
		http.Error(w, "invalid scope.", 400)
		return
	}
	u, err := url.Parse(redirectURI)
	if err != nil {
		http.Error(w, "invalid redirect uri.", 400)
		return
	}
	q := u.Query()
	q.Set("state", state)
	q.Set("code", "code")
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), 301)
}
