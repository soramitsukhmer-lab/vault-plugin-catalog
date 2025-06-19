package command

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

func uninstallPluginCommand(m *manager.PluginManager) *cli.Command {
	return &cli.Command{
		Name:      "uninstall",
		Usage:     "Uninstall a plugin by name",
		Args:      true,
		ArgsUsage: "<plugin-name>",
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}
			// Here you would call the UninstallPlugin method from your PluginManager
			fmt.Printf("Uninstalling plugin: %s\n", name)
			return nil
		},
	}
}
