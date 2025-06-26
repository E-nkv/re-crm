package app

import (
	"log"
	"net/http"
	"re-crm/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	AuthSvc *services.AuthService
	Router  *chi.Mux
	writers struct {
		http httpWriter
		ws   websocketWriter
	}
}

func NewApp(authSvc *services.AuthService) *App {
	return &App{AuthSvc: authSvc}
}

func (app *App) Mount() {
	mu := chi.NewRouter()

	mu.Use(middleware.Logger)
	mu.Use(corsMdw)
	mu.Get("/", app.HandleHome)
	mu.Route("/api", func(r chi.Router) {
		r.With(app.authMdw).Get("/dashboard", app.HandleDashboard)
		r.With(app.authMdw).Get("/chat", app.HandleChat)
		r.Post("/login", app.HandleLogin)
		r.With(app.authMdw).Put("/logout", app.HandleLogout)
		r.With(app.authMdw, app.isAdminMdw).Post("/create-account", app.HandleCreateAccount)
		r.With(app.authMdw).Get("/me", app.HandleMe)

	})
	app.Router = mu
}

func (app *App) Run() error {
	log.Println("🚀 running on port :8080")
	return http.ListenAndServe("0.0.0.0:8080", app.Router)
}
