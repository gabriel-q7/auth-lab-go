package oidc

import (
	"context"
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	clientID     = "go-client"
	clientSecret = "CLIENT_SECRET"
	redirectURL  = "http://localhost:8080/auth/oidc/callback"
	providerURL  = "http://localhost:8080/realms/auth-lab"

	provider *oidc.Provider
	oauth2Config oauth2.Config
	verifier *oidc.IDTokenVerifier
)

func init() {
	ctx := context.Background()

	var err error
	provider, err = oidc.NewProvider(ctx, providerURL)
	if err != nil {
		log.Fatal(err)
	}

	oauth2Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	verifier = provider.Verifier(&oidc.Config{ClientID: clientID})
}

func GetLoginURL() string {
	return oauth2Config.AuthCodeURL("state")
}

func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return oauth2Config.Exchange(ctx, code)
}

func VerifyToken(ctx context.Context, token *oauth2.Token) (map[string]interface{}, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("no id_token")
	}

	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return nil, err
	}

	claims := map[string]interface{}{}
	if err := idToken.Claims(&claims); err != nil {
		return nil, err
	}

	return claims, nil
}
