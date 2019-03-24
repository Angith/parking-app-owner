package api

import (
	"net/http"

	"github.com/google/logger"
)

func (a *App) Serve() {
	r := a.Router()
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		logger.Fatal(err, "Server disconnected")
	}
}
