package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func updateCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update the local plugin catalog from the remote source",
		Action: func(c *cli.Context) error {
			localCatalogPath := c.String("local-catalog-path")
			remoteCatalogURL := c.String("remote-catalog-url")

			// Fetch the remote catalog
			resp, err := http.Get(remoteCatalogURL)
			if err != nil {
				return fmt.Errorf("failed to fetch remote catalog: %w", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("failed to fetch remote catalog: %s", resp.Status)
			}

			// Create or open the local catalog file
			file, err := os.Create(localCatalogPath)
			if err != nil {
				return fmt.Errorf("failed to create local catalog file: %w", err)
			}
			defer file.Close()

			// Copy the remote catalog to the local file
			_, err = io.Copy(file, resp.Body)
			if err != nil {
				return fmt.Errorf("failed to write to local catalog file: %w", err)
			}
			fmt.Printf("Successfully updated local plugin catalog at %s\n", localCatalogPath)
			return nil
		},
	}
}
