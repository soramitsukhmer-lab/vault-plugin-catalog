package command

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

func installPluginCommand(pm *manager.PluginManager) *cli.Command {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "register",
			Usage: "Register the plugin after installation",
		},
	}

	flags = append(flags, genericFlags...)

	return &cli.Command{
		Name:      "install",
		Usage:     "Install a plugin from the catalog",
		Args:      true,
		ArgsUsage: "<plugin-name>",
		Flags:     flags,
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}

			// Here you would call the InstallPlugin method from your PluginManager
			if err := pm.InstallPlugin(name); err != nil {
				return fmt.Errorf("failed to install plugin %s: %w", name, err)
			}

			return nil
		},
	}
}
