package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/pluginmanager"
	"github.com/urfave/cli/v2"
)

func updateCatalogCommand(pm *pluginmanager.PluginManager) *cli.Command {
	flags := []cli.Flag{}
	flags = append(flags, genericFlags...)

	return &cli.Command{
		Name:  "update",
		Usage: "Update the local plugin catalog from the remote source",
		Flags: flags,
		Before: func(c *cli.Context) error {
			localCatalogPath := c.String("local-catalog-path")
			if localCatalogPath == "" {
				return fmt.Errorf("local catalog path cannot be empty")
			}

			// Check if local catalog path is exist, if not, create it
			strs := strings.Split(localCatalogPath, "/")
			localCatalogDir := strings.Join(strs[:len(strs)-1], "/")
			if _, err := os.Stat(localCatalogDir); os.IsNotExist(err) {
				if err := os.MkdirAll(localCatalogDir, 0755); err != nil {
					return fmt.Errorf("failed to create local catalog directory: %w", err)
				}
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			if err := pm.UpdateCatalog(); err != nil {
				return fmt.Errorf("failed to update catalog: %w", err)
			}
			return nil
		},
	}
}
