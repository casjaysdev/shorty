// File: internal/api/auth/link_routes.go
// Purpose: Authenticated routes for managing user-owned links and their metadata.

package auth

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"shorty/internal/core"
	"shorty/internal/lib/httpx"
	"shorty/internal/models"
)

func RegisterLinkRoutes(r chi.Router) {
	r.Route("/link", func(r chi.Router) {
		r.Post("/", httpx.Handler(createLink))
		r.Get("/", httpx.Handler(listLinks))
		r.Get("/{id}", httpx.Handler(getLink))
		r.Put("/{id}", httpx.Handler(updateLink))
		r.Delete("/{id}", httpx.Handler(deleteLink))
		r.Post("/{id}/clone", httpx.Handler(cloneLink))
		r.Post("/{id}/toggle", httpx.Handler(toggleLink))
		r.Post("/{id}/analytics", httpx.Handler(linkAnalytics))
	})
}

func createLink(w http.ResponseWriter, r *http.Request) error {
	var link models.Link
	if err := httpx.Decode(r, &link); err != nil {
		return err
	}
	return httpx.JSON(w, core.CreateLink(r.Context(), &link))
}

func listLinks(w http.ResponseWriter, r *http.Request) error {
	return httpx.JSON(w, core.ListUserLinks(r.Context()))
}

func getLink(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	return httpx.JSON(w, core.GetUserLink(r.Context(), id))
}

func updateLink(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	var link models.Link
	if err := httpx.Decode(r, &link); err != nil {
		return err
	}
	return httpx.JSON(w, core.UpdateUserLink(r.Context(), id, &link))
}

func deleteLink(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	return httpx.NoContent(core.DeleteUserLink(r.Context(), id))
}

func cloneLink(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	return httpx.JSON(w, core.CloneLink(r.Context(), id))
}

func toggleLink(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	return httpx.JSON(w, core.ToggleLink(r.Context(), id))
}

func linkAnalytics(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	return httpx.JSON(w, core.GetLinkAnalytics(r.Context(), id))
}
