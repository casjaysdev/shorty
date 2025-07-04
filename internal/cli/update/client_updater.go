// File: internal/cli/update/client_updater.go
// Purpose: Implements update logic for downloading and installing the latest CLI binary.

package update

import (
	"archive/zip"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"shorty/internal/config"
)

const (
	githubOrg  = "casjaysdev"
	project    = "shorty"
	cliName    = "shorty-cli"
	jsonURLVar = "SHORTY_CLI_RELEASE_JSON"
	urlVar     = "SHORTY_CLI_RELEASE_URL"
)

func UpdateClient(binary, osys, arch string) error {
	url := os.Getenv(urlVar)
	jsonURL := os.Getenv(jsonURLVar)

	if url == "" && jsonURL != "" {
		res, err := http.Get(jsonURL)
		if err != nil {
			return fmt.Errorf("failed to fetch JSON manifest: %w", err)
		}
		defer res.Body.Close()

		var obj struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(res.Body).Decode(&obj); err != nil {
			return fmt.Errorf("invalid JSON manifest: %w", err)
		}
		url = obj.URL
	}

	if url == "" {
		url = fmt.Sprintf("https://github.com/%s/%s/releases/latest/download/%s-%s-cli-%s.tar.gz",
			githubOrg, project, osys, project, arch)
	}

	tmpFile := filepath.Join(os.TempDir(), "shorty-cli-update")
	out, err := os.Create(tmpFile)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile)
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if strings.HasSuffix(url, ".gz") {
		gz, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer gz.Close()
		_, err = io.Copy(out, gz)
	} else if strings.HasSuffix(url, ".zip") {
		zr, err := zip.NewReader(resp.Body, resp.ContentLength)
		if err != nil {
			return err
		}
		for _, f := range zr.File {
			rc, err := f.Open()
			if err != nil {
				return err
			}
			defer rc.Close()
			io.Copy(out, rc)
			break
		}
	} else {
		_, err = io.Copy(out, resp.Body)
	}
	if err != nil {
		return err
	}

	out.Chmod(0755)

	installPath := config.ResolveBinInstallPath()
	if installPath == "" {
		installPath = filepath.Join(os.Getenv("HOME"), ".local", "bin")
	}
	finalPath := filepath.Join(installPath, cliName)

	if err := os.Rename(tmpFile, finalPath); err != nil {
		return fmt.Errorf("install failed: %w", err)
	}

	return nil
}
