package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func listCommand() *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all available plugins in the catalog",
		Action: func(c *cli.Context) error {
			localCatalogPath := c.String("local-catalog-path")

			if localCatalogPath == "" {
				fmt.Println("Error: Local catalog path is not specified.")
				return fmt.Errorf("local catalog path is required")
			}

			file, err := os.ReadFile(localCatalogPath)
			if err != nil {
				fmt.Printf("Error reading local catalog file: %v\n", err)
				return err
			}

			var catalog Catalog
			if err := json.Unmarshal(file, &catalog); err != nil {
				fmt.Printf("Error parsing local catalog file: %v\n", err)
				return err
			}

			fmt.Println("Listing all available plugins in the catalog...")
			fmt.Println()

			// TODO: Add support for other plugin types (e.g., auth, database)

			fmt.Println("Available secret plugins:")
			for k, v := range catalog.Plugins.Secrets {
				fmt.Printf("  - %s (%s)\n", k, v.Version)
			}

			return nil
		},
	}
}
