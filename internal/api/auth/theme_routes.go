// File: internal/api/auth/theme_routes.go
// Purpose: Authenticated routes for saving and retrieving user/org theme configurations (custom UI/UX theming).

package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"shorty/internal/core"
	"shorty/internal/lib/httpx"
	"shorty/internal/middleware"
	"shorty/internal/models"
)

func RegisterThemeRoutes(r chi.Router) {
	r.Route("/themes", func(r chi.Router) {
		r.Use(middleware.RequireAuth)

		r.Get("/", httpx.Handler(getTheme))
		r.Post("/", httpx.Handler(saveTheme))
	})
}

func getTheme(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	theme, err := core.GetTheme(ctx)
	if err != nil {
		return err
	}
	return httpx.JSON(w, theme)
}

func saveTheme(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	var t models.Theme
	if err := httpx.Decode(r, &t); err != nil {
		return err
	}
	err := core.SaveTheme(ctx, &t)
	if err != nil {
		return err
	}
	return httpx.Success(w)
}
