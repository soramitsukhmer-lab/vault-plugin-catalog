package command

import (
	"fmt"

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
		Action: func(c *cli.Context) error {
			if err := pm.UpdateCatalog(); err != nil {
				return fmt.Errorf("failed to update catalog: %w", err)
			}
			return nil
		},
	}
}
