package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	defaultRemotePluginCatalogURL = "https://raw.githubusercontent.com/soramitsukhmer-lab/vault-plugin-catalog/refs/heads/main/catalog.json"
	defaultLocalPluginCatalogPath = "/etc/vault-plugin-catalog/catalog.json"
)

func Run(args []string) error {
	app := &cli.App{
		Name:  "vault-plugin-catalog",
		Usage: "A cli tool to manage HashiCorp Vault plugins",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "local-catalog-path",
				Usage:   "Path to the local plugin catalog file",
				Value:   defaultLocalPluginCatalogPath,
				EnvVars: []string{"LOCAL_PLUGIN_CATALOG_PATH"},
			},
			&cli.StringFlag{
				Name:    "remote-catalog-url",
				Usage:   "URL of the remote plugin catalog",
				Value:   defaultRemotePluginCatalogURL,
				EnvVars: []string{"REMOTE_PLUGIN_CATALOG_URL"},
			},
		},
		Commands: []*cli.Command{
			installPluginCommand(),
			uninstallPluginCommand(),
			registerPluginCommand(),
			updateCatalogCommand(),
		},
		Action: func(*cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	return app.Run(args)
}
