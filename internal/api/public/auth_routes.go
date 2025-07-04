// File: internal/api/public/auth_routes.go
// Purpose: Public API routes for user authentication and identity operations.

package public

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"shorty/internal/core"
	"shorty/internal/lib/httpx"
	"shorty/internal/models"
)

func RegisterAuthRoutes(r chi.Router) {
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", httpx.Handler(register))
		r.Post("/login", httpx.Handler(login))
		r.Post("/logout", httpx.Handler(logout))
		r.Post("/refresh", httpx.Handler(refresh))
		r.Get("/me", httpx.Handler(me))
	})
}

func register(w http.ResponseWriter, r *http.Request) error {
	var u models.User
	if err := httpx.Decode(r, &u); err != nil {
		return err
	}
	return core.RegisterUser(r.Context(), &u)
}

func login(w http.ResponseWriter, r *http.Request) error {
	var cred models.Credentials
	if err := httpx.Decode(r, &cred); err != nil {
		return err
	}
	token, err := core.Login(r.Context(), &cred)
	if err != nil {
		return err
	}
	return httpx.JSON(w, token)
}

func logout(w http.ResponseWriter, r *http.Request) error {
	return core.Logout(r.Context())
}

func refresh(w http.ResponseWriter, r *http.Request) error {
	token, err := core.RefreshToken(r.Context())
	if err != nil {
		return err
	}
	return httpx.JSON(w, token)
}

func me(w http.ResponseWriter, r *http.Request) error {
	user, err := core.GetAuthenticatedUser(r.Context())
	if err != nil {
		return err
	}
	return httpx.JSON(w, user)
}
