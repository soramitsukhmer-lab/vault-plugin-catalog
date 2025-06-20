package command

import (
	"fmt"
	"os"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/pluginmanager"
	"github.com/urfave/cli/v2"
)

func installPluginCommand(pm *pluginmanager.PluginManager) *cli.Command {
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
		Before: func(c *cli.Context) error {
			vaultPluginDir := c.String("vault-plugin-dir")
			if vaultPluginDir == "" {
				return fmt.Errorf("vault plugin directory cannot be empty")
			}

			// Check if vault plugin directory exists, if not, create it
			if _, err := os.Stat(vaultPluginDir); os.IsNotExist(err) {
				if err := os.MkdirAll(vaultPluginDir, 0755); err != nil {
					return fmt.Errorf("failed to create vault plugin directory: %w", err)
				}
			}

			return nil
		},
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}

			register := c.Bool("register")
			if register {
				// If the register flag is set, we install and register the plugin
				if err := pm.InstallPluginWithRegistration(name); err != nil {
					return fmt.Errorf("failed to install and register plugin %s: %w", name, err)
				}
				fmt.Printf("Plugin %s installed and registered successfully.\n", name)
				return nil
			}

			// Here you would call the InstallPlugin method from your PluginManager
			if err := pm.InstallPlugin(name); err != nil {
				return fmt.Errorf("failed to install plugin %s: %w", name, err)
			}

			return nil
		},
	}
}
