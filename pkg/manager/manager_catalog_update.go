package manager

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (p *PluginManager) UpdateCatalog() error {
	// This method should handle the update of the plugin catalog.
	// It will likely involve fetching the latest catalog from a remote source
	// and updating the local catalog file.

	if p.RemoteCatalogURL == "" {
		return fmt.Errorf("remote catalog URL is not set")
	}

	// Check if local catalog file exists, if it does, back it up
	if _, err := os.Stat(p.LocalCatalogPath); os.IsExist(err) {
		backupPath := p.LocalCatalogPath + ".bak"
		if err := os.Rename(p.LocalCatalogPath, backupPath); err != nil {
			return fmt.Errorf("failed to backup existing catalog file: %w", err)
		}
	}

	fmt.Println("Updating catalog from remote source...")

	// Fetch the latest catalog from the remote URL
	resp, err := http.Get(p.RemoteCatalogURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch catalog: %s", resp.Status)
	}

	// Here you would typically read the response body and update the local catalog file.
	file, err := os.Create(p.LocalCatalogPath)
	if err != nil {
		return fmt.Errorf("failed to create local catalog file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write to local catalog file: %w", err)
	}

	fmt.Println("Catalog updated successfully.")

	return nil
}
