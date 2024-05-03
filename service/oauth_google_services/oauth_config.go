package oauthgoogleservices

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOauthAppConfig() (*oauth2.Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("(From oauth) Error loading .env file:", err)
		return nil, err
	}

	clientid := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	conf := &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/api/v1/auth/callback",
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
	return conf, nil
}
