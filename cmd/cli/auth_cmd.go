// File: cmd/cli/auth_cmd.go
// Purpose: CLI commands for authentication: login, logout, token test.

package cli

import (
	"fmt"
	"os"

	"shorty/internal/auth"
	"shorty/internal/cli/util"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)

	authCmd.AddCommand(loginCmd)
	authCmd.AddCommand(logoutCmd)
	authCmd.AddCommand(tokenCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in and store access token",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := auth.LoginInteractive()
		if err != nil {
			util.Fatalf("Login failed: %v", err)
		}
		util.SaveToken(token)
		util.Infof("Login successful")
	},
}

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out and clear token",
	Run: func(cmd *cobra.Command, args []string) {
		util.DeleteToken()
		util.Infof("Logged out")
	},
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Show or test the current token",
	Run: func(cmd *cobra.Command, args []string) {
		token := util.LoadToken()
		if token == "" {
			util.Errorf("No token found")
			os.Exit(1)
		}
		fmt.Printf("Token: %s\n", token)
		if ok := auth.VerifyToken(token); !ok {
			util.Warnf("Token may be invalid or expired")
		} else {
			util.Infof("Token is valid")
		}
	},
}
