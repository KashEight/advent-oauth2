package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

	idToken, err := getToken(code)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	email, err := parseJWT(idToken)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	fmt.Fprintf(w, "email: %s", email)
}

const authUrl = "https://oauth2.googleapis.com/token"

type authCallback struct {
	IdToken string `json:"id_token"`
}

func getToken(code string) (string, error) {
	values := url.Values{}

	values.Add("code", code)
	values.Add("client_id", AuthConfig.ClientID)
	values.Add("client_secret", AuthConfig.ClientSecret)
	values.Add("redirect_uri", AuthConfig.RedirectURL)
	values.Add("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", authUrl, strings.NewReader(values.Encode()))

	if err != nil {
		return "", fmt.Errorf("failed building new request")
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := new(http.Client)

	resp, err := client.Do(req)

	if err != nil {
		return "", fmt.Errorf("failed request")
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("failed reading body")
	}

	fmt.Println("------ Callback JSON Data ------")
	fmt.Println(string(b))

	clb := &authCallback{}

	if err := json.Unmarshal(b, clb); err != nil {
		return "", fmt.Errorf("failed unmarshal json data (in getToken())")
	}

	return clb.IdToken, nil
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
