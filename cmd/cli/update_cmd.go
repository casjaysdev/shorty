// File: cmd/cli/update_cmd.go
// Purpose: Provides CLI update capabilities for both client and (future) server components.

package cli

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"shorty/internal/cli/update"
	"shorty/internal/cli/util"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.AddCommand(updateClientCmd)
	updateCmd.AddCommand(updateServerCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update CLI or server (future)",
}

var updateClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Update the CLI to the latest release",
	Run: func(cmd *cobra.Command, args []string) {
		util.Infof("Checking for latest release...")

		arch := runtime.GOARCH
		osys := runtime.GOOS
		binName := strings.TrimSuffix(os.Args[0], ".exe")

		err := update.UpdateClient(binName, osys, arch)
		if err != nil {
			util.Fatalf("Update failed: %v", err)
		}
		util.Infof("Client updated successfully")
	},
}

var updateServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Update the server (not implemented yet)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Server update is not yet implemented.")
	},
}
