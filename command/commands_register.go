package command

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/pluginmanager"
	"github.com/urfave/cli/v2"
)

func registerPluginCommand(pm *pluginmanager.PluginManager) *cli.Command {
	flags := []cli.Flag{}
	flags = append(flags, genericFlags...)

	return &cli.Command{
		Name:      "register",
		Usage:     "Register a new plugin in the catalog",
		Args:      true,
		ArgsUsage: "<plugin-name>",
		Flags:     flags,
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}
			// Here you would call the RegisterPlugin method from your PluginManager
			if err := pm.RegisterPlugin(name); err != nil {
				return fmt.Errorf("failed to register plugin %s: %w", name, err)
			}

			return nil
		},
	}
}
