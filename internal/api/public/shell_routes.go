// File: internal/api/public/shell_routes.go
// Purpose: Serves CLI shell completion scripts for supported shells via API routes.

package public

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerShellRoutes(r chi.Router) {
	r.Route("/shell", func(r chi.Router) {
		r.Get("/{shell}", handleShellCompletion)
	})
}

func handleShellCompletion(w http.ResponseWriter, r *http.Request) {
	shell := chi.URLParam(r, "shell")

	completions := map[string]string{
		"bash":      "# bash completion script for shorty-cli\ncomplete -C shorty-cli shorty",
		"zsh":       "# zsh completion script for shorty-cli\ncompdef _shorty shorty",
		"fish":      "# fish completion script for shorty-cli\ncomplete --command shorty",
		"powershell": "# powershell completion script for shorty-cli\nRegister-ArgumentCompleter -CommandName shorty",
		"nushell":   "# nushell completion for shorty-cli\nsource ~/.config/nushell/scripts/shorty.nu",
	}

	script, ok := completions[shell]
	if !ok {
		http.Error(w, "unsupported shell", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(script))
}
