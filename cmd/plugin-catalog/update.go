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
			// Logic to update the local plugin catalog
			// This would typically involve fetching the remote catalog
			// and saving it to the local path.
			fmt.Println("Pulling the latest plugin catalog to", localPluginCatalogPath, "from", remotePluginCatalogURL)
			// Here you would implement the actual logic to fetch the remote catalog

			res, err := http.Get(remotePluginCatalogURL)
			if err != nil {
				return fmt.Errorf("failed to fetch remote catalog: %w", err)
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				return fmt.Errorf("failed to fetch remote catalog: received status code %d", res.StatusCode)
			}

			// Read the response body and write it to the local file
			// For simplicity, we are not actually reading the body here.

			file, err := os.Create(localPluginCatalogPath)
			if err != nil {
				return fmt.Errorf("failed to create local catalog file: %w", err)
			}
			defer file.Close()

			// Here you would copy the response body to the file
			_, err = io.Copy(file, res.Body)
			if err != nil {
				return fmt.Errorf("failed to write to local catalog file: %w", err)
			}

			return nil
		},
	}
}
