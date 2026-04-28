package handlers

import (
	"context"
	"fmt"
	"net/http"

	"auth-lab-go/internal/oidc"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func OIDCLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oidc.GetLoginURL()
	http.Redirect(w, r, url, http.StatusFound)
}

func OIDCCallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	token, err := oidc.ExchangeCode(context.Background(), code)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	claims, err := oidc.VerifyToken(context.Background(), token)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	fmt.Fprintf(w, "User info: %v", claims)
}
