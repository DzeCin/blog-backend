package blog

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Claims struct {
	Email    string `json:"email"`
	Verified bool   `json:"email_verified"`
}

func NewOAuthConfig(oidcUrl string, scopes []string, clientID string, clientSecret string, redirectURL string) (*oauth2.Config, *oidc.Provider) {

	scopes = append(scopes, oidc.ScopeOpenID)

	provider, err := oidc.NewProvider(context.Background(), oidcUrl)

	if err != nil {
		panic(err)
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: scopes,
	}

	return oauth2Config, provider
}

func GetClaims(oauth2Config *oauth2.Config, token string, provider oidc.Provider) Claims {

	verifier := provider.Verifier(&oidc.Config{ClientID: oauth2Config.ClientID})

	oauth2Token, err := (*oauth2Config).Exchange(context.Background(), token)

	if err != nil {
		// handle error
	}

	// Extract the ID Token from OAuth2 token.
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		// handle missing token
	}

	// Parse and verify ID Token payload.
	idToken, err := verifier.Verify(context.Background(), rawIDToken)
	if err != nil {
		// handle error
	}

	// Extract custom claims
	var claims Claims

	if err := idToken.Claims(&claims); err != nil {
		// handle error
	}

	return claims

}
