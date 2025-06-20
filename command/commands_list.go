package command

import (
	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

func listPluginCommand(pm *manager.PluginManager) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List plugins in the catalog",
		Action: func(c *cli.Context) error {
			return pm.ListPlugins()
		},
	}
}
