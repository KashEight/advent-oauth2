package main

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var AuthConfig = &oauth2.Config{
	ClientID:     "<Your ClientID>",
	ClientSecret: "<Your Client Secret>",
	Endpoint:     google.Endpoint,
	Scopes: []string{
		"profile",
		"email",
	},
	RedirectURL: "http://localhost:8080/callback",
}
