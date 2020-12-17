package main

import "net/http"

func handlerSignIn(w http.ResponseWriter, r *http.Request) {
	url := AuthConfig.AuthCodeURL("state") // 本来は CSRF 対策でランダムな値
	http.Redirect(w, r, url, 302)
}
