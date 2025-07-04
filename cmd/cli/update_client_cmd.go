// File: cmd/cli/update_client_cmd.go
// Purpose: CLI command to update the client binary from remote release or manifest.

package cli

import (
	"fmt"
	"os"
	"runtime"

	"shorty/internal/cli/update"

	"github.com/spf13/cobra"
)

func newUpdateClientCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "client",
		Short: "Update the shorty CLI to the latest version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("🔄 Checking for latest version of shorty CLI...")

			err := update.UpdateClient("shorty-cli", runtime.GOOS, runtime.GOARCH)
			if err != nil {
				fmt.Fprintf(os.Stderr, "❌ Update failed: %v\n", err)
				return err
			}

			fmt.Println("✅ shorty CLI updated successfully.")
			return nil
		},
	}
	return cmd
}
