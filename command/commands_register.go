package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func registerPluginCommand() *cli.Command {
	return &cli.Command{
		Name:  "register",
		Usage: "Register a new plugin in the catalog",
		Action: func(c *cli.Context) error {
			name := c.Args().First()
			if name == "" {
				return fmt.Errorf("plugin name is required")
			}
			// Here you would call the RegisterPlugin method from your PluginManager
			fmt.Printf("Registering plugin: %s\n", name)
			return nil
		},
	}
}
