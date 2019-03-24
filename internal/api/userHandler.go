package api

import (
	"net/http"
)

func (a *App) handleLogin(w http.ResponseWriter, r *http.Request) {
	userData, err := a.Login()
}
