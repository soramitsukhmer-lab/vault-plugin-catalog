package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func updateCatalogCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update the local plugin catalog from the remote source",
		Action: func(c *cli.Context) error {
			// Here you would call the UpdateCatalog method from your PluginManager
			fmt.Println("Updating plugin catalog...")
			return nil
		},
	}
}
