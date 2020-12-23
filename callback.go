package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func handlerCallback(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	if queries == nil {
		fmt.Fprintf(w, "Invalid URL.")
		return
	}

	fmt.Println("------ Queries ------")

	for k, v := range queries {
		fmt.Println(k, v)
	}

	code := queries.Get("code")

	token, err := AuthConfig.Exchange(context.Background(), code)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	idToken := token.Extra("id_token").(string)

	fmt.Println("------ ID Token ------")
	fmt.Println(idToken)

	email, err := parseJWT(idToken)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	fmt.Fprintf(w, "email: %s", email)
}

type jwtData struct {
	Email string `json:"email"`
}

func parseJWT(token string) (string, error) {
	jwt := strings.Split(token, ".")
	payload := strings.TrimSuffix(jwt[1], "=")
	b, err := base64.RawURLEncoding.DecodeString(payload)

	if err != nil {
		return "", fmt.Errorf("failed decoding base64")
	}

	fmt.Println("------ JWT Data ------")
	fmt.Println(string(b))

	jd := &jwtData{}

	if err := json.Unmarshal(b, jd); err != nil {
		return "", fmt.Errorf("failed unmarshal json data (in parseJWT())")
	}

	return jd.Email, nil
}
