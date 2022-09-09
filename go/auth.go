package blog

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"os"
	"strings"
)

type Claims struct {
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

func GetUserClaims(token string) *Claims {

	var claims Claims

	provider, err := oidc.NewProvider(context.Background(), "")
	if err != nil {
		// handle error
	}

	var verifier = provider.Verifier(&oidc.Config{ClientID: ""})

	// Remove "Bearer " in front of the token and verify the token
	fmt.Println(strings.Split(token, "Bearer ")[1])
	idToken, err := verifier.Verify(context.Background(), strings.Split(token, "Bearer ")[1])

	if err != nil {
		fmt.Println(err)
		return &claims
	}

	if err := idToken.Claims(&claims); err != nil {
		os.Exit(1)
	}

	fmt.Println(claims.Name)

	return &claims

}
