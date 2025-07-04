// File: internal/api/auth/domain_routes.go
// Purpose: Authenticated routes for managing custom domains (add, list, delete, validate).

package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"shorty/internal/core"
	"shorty/internal/lib/httpx"
	"shorty/internal/middleware"
	"shorty/internal/models"
)

func RegisterDomainRoutes(r chi.Router) {
	r.Route("/domains", func(r chi.Router) {
		r.Use(middleware.RequireAuth)

		r.Get("/", httpx.Handler(listDomains))
		r.Post("/", httpx.Handler(addDomain))
		r.Delete("/{id}", httpx.Handler(deleteDomain))
		r.Get("/validate", httpx.Handler(validateDomain))
	})
}

func listDomains(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	domains, err := core.ListDomains(ctx)
	if err != nil {
		return err
	}
	return httpx.JSON(w, domains)
}

func addDomain(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	var d models.Domain
	if err := httpx.Decode(r, &d); err != nil {
		return err
	}
	if err := core.AddDomain(ctx, &d); err != nil {
		return err
	}
	return httpx.Created(w)
}

func deleteDomain(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	id := chi.URLParam(r, "id")
	return core.DeleteDomain(ctx, id)
}

func validateDomain(w http.ResponseWriter, r *http.Request) error {
	ctx := httpx.GetContext(r)
	q := r.URL.Query().Get("domain")
	if q == "" {
		return httpx.BadRequestf("missing domain")
	}
	return httpx.JSON(w, map[string]bool{"valid": core.ValidateDomain(q)})
}
