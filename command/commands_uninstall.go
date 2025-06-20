package command

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

func uninstallPluginCommand(pm *manager.PluginManager) *cli.Command {
	flags := []cli.Flag{}
	flags = append(flags, genericFlags...)

	return &cli.Command{
		Name:      "uninstall",
		Usage:     "Uninstall a plugin by name",
		Args:      true,
		ArgsUsage: "<plugin-name>",
		Flags:     flags,
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}
			// Here you would call the UninstallPlugin method from your PluginManager
			fmt.Printf("Uninstalling plugin: %s\n", name)
			if err := pm.UninstallPlugin(name); err != nil {
				return fmt.Errorf("failed to uninstall plugin %s: %w", name, err)
			}

			return nil
		},
	}
}
