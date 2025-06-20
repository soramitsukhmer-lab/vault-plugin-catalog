package command

import (
	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

func listPluginCommand(pm *manager.PluginManager) *cli.Command {
	flags := []cli.Flag{}
	flags = append(flags, genericFlags...)

	return &cli.Command{
		Name:  "list",
		Usage: "List plugins in the catalog",
		Flags: flags,
		Action: func(c *cli.Context) error {
			return pm.ListPlugins()
		},
	}
}
