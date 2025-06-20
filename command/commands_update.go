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
			if err := m.UpdateCatalog(); err != nil {
				return fmt.Errorf("failed to update catalog: %w", err)
			}
			return nil
		},
	}
}
