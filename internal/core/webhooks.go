// File: internal/core/webhooks.go
// Purpose: Core logic for managing and dispatching webhooks across user, org, and system scopes.

package core

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"shorty/internal/db"
	"shorty/internal/models"
)

var (
	ErrWebhookNotFound = errors.New("webhook not found")
	providers          = map[string]func(models.Webhook, any) ([]byte, error){
		"generic": formatGeneric,
		"slack":   formatSlack,
		"discord": formatDiscord,
		"teams":   formatTeams,
		"zapier":  formatGeneric,
		"ifttt":   formatGeneric,
	}
)

func DispatchWebhook(ctx context.Context, scope, event string, payload any) error {
	hooks := []models.Webhook{}
	err := db.Find(ctx, &hooks, "enabled = true AND scope = ? AND event = ?", scope, event)
	if err != nil || len(hooks) == 0 {
		return ErrWebhookNotFound
	}

	for _, hook := range hooks {
		go deliverWebhook(hook, payload)
	}

	return nil
}

func deliverWebhook(hook models.Webhook, payload any) {
	ctx := context.Background()
	format := providers[hook.Provider]
	if format == nil {
		format = formatGeneric
	}

	body, err := format(hook, payload)
	if err != nil {
		return
	}

	req, _ := http.NewRequest("POST", hook.URL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "ShortyWebhook/1.0")

	if hook.Secret != "" {
		h := hmac.New(sha256.New, []byte(hook.Secret))
		h.Write(body)
		signature := hex.EncodeToString(h.Sum(nil))
		req.Header.Set("X-Signature", signature)
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return
	}

	// implement retry logic if needed (e.g. backoff, logging)
}

func formatGeneric(hook models.Webhook, payload any) ([]byte, error) {
	return json.Marshal(payload)
}

func formatSlack(hook models.Webhook, payload any) ([]byte, error) {
	return json.Marshal(map[string]any{
		"text": fmt.Sprintf("Event: %s\n%s", hook.Event, payload),
	})
}

func formatDiscord(hook models.Webhook, payload any) ([]byte, error) {
	return json.Marshal(map[string]any{
		"content": fmt.Sprintf("Event: %s\n%s", hook.Event, payload),
	})
}

func formatTeams(hook models.Webhook, payload any) ([]byte, error) {
	return json.Marshal(map[string]any{
		"text": fmt.Sprintf("Event: %s\n%s", hook.Event, payload),
	})
}
