package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func installPluginCommand() *cli.Command {
	return &cli.Command{
		Name:      "install",
		Usage:     "Install a plugin by name",
		Args:      true,
		ArgsUsage: "<plugin-name>",
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}
			// Here you would call the InstallPlugin method from your PluginManager
			fmt.Printf("Installing plugin: %s\n", name)
			return nil
		},
	}
}
