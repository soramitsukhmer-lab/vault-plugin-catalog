package command

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

func updateCatalogCommand(m *manager.PluginManager) *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update the local plugin catalog from the remote source",
		Action: func(c *cli.Context) error {
			// Here you would call the UpdateCatalog method from your PluginManager
			fmt.Println("Updating plugin catalog...")
			return nil
		},
	}
}
