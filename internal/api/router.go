package api

import "github.com/go-chi/chi"

// Router sets up parking-app router
func (a *App) Router() chi.Router {
	r := chi.NewRouter()
	r.Route("parking-app/v1/", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Get("/register", a.handleLogin)
		})
	})
	return r
}
