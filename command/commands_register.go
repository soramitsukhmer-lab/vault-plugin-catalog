package command

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

func registerPluginCommand(m *manager.PluginManager) *cli.Command {
	return &cli.Command{
		Name:      "register",
		Usage:     "Register a new plugin in the catalog",
		Args:      true,
		ArgsUsage: "<plugin-name>",
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}
			// Here you would call the RegisterPlugin method from your PluginManager
			fmt.Printf("Registering plugin: %s\n", name)
			if err := m.RegisterPlugin(name); err != nil {
				return fmt.Errorf("failed to register plugin %s: %w", name, err)
			}

			return nil
		},
	}
}
