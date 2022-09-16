package blog

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type Claims struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

func GetUserClaims(token string, providerURL string, clientID string) *Claims {

	var claims Claims
	var idToken *oidc.IDToken

	provider, err := oidc.NewProvider(context.Background(), providerURL)
	if err != nil {
		return &claims
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: clientID})

	// Remove "Bearer " in front of the token and verify the token

	if strings.Split(token, "Bearer ")[0] == token {
		fmt.Println(&claims)
		return &claims
	} else {

		idToken, err = verifier.Verify(context.Background(), strings.Split(token, "Bearer ")[1])

		if err != nil {
			claims.Roles = []string{""}
			claims.Name = ""
			return &claims
		}
	}

	if err := idToken.Claims(&claims); err != nil {
		os.Exit(1)
	}

	return &claims

}
