// File: internal/api/auth/user_webhooks.go
// Purpose: User-scoped API endpoints for managing webhooks (create, list, delete, test).

package auth

import (
	"net/http"

	"shorty/internal/core"
	"shorty/internal/models"
	"shorty/internal/web"

	"github.com/go-chi/chi/v5"
)

func handleUserWebhooks(r chi.Router) {
	r.Get("/", listUserWebhooks)
	r.Post("/", createUserWebhook)
	r.Delete("/{id}", deleteUserWebhook)
	r.Post("/{id}/test", testUserWebhook)
}

func listUserWebhooks(w http.ResponseWriter, r *http.Request) {
	user := web.GetUser(r)
	hooks, err := core.ListWebhooks(r.Context(), "user", user.ID)
	if err != nil {
		web.Error(w, r, err)
		return
	}
	web.JSON(w, r, hooks)
}

func createUserWebhook(w http.ResponseWriter, r *http.Request) {
	user := web.GetUser(r)
	var req models.Webhook
	if err := web.Decode(r, &req); err != nil {
		web.Error(w, r, err)
		return
	}
	req.Scope = "user"
	req.OwnerID = user.ID
	if err := core.SaveWebhook(r.Context(), req); err != nil {
		web.Error(w, r, err)
		return
	}
	web.Respond(w, r, http.StatusCreated, req)
}

func deleteUserWebhook(w http.ResponseWriter, r *http.Request) {
	user := web.GetUser(r)
	id := chi.URLParam(r, "id")
	if err := core.DeleteWebhook(r.Context(), "user", user.ID, id); err != nil {
		web.Error(w, r, err)
		return
	}
	web.Respond(w, r, http.StatusNoContent, nil)
}

func testUserWebhook(w http.ResponseWriter, r *http.Request) {
	user := web.GetUser(r)
	id := chi.URLParam(r, "id")
	if err := core.TestWebhook(r.Context(), "user", user.ID, id); err != nil {
		web.Error(w, r, err)
		return
	}
	web.Respond(w, r, http.StatusAccepted, map[string]string{"status": "sent"})
}
