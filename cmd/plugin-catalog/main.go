package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	remotePluginCatalogURL = "https://raw.githubusercontent.com/soramitsukhmer-lab/vault-plugin-catalog/refs/heads/main/catalog.json"
	localPluginCatalogPath = "/etc/vault-plugin-catalog/catalog.json"
)

func main() {
	app := &cli.App{
		Name:  "plugin-catalog",
		Usage: "A cli tool to manage HashiCorp Vault plugins",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "local-catalog-path",
				Usage:   "Path to the local plugin catalog file",
				Value:   localPluginCatalogPath,
				EnvVars: []string{"LOCAL_PLUGIN_CATALOG_PATH"},
			},
			&cli.StringFlag{
				Name:    "remote-catalog-url",
				Usage:   "URL of the remote plugin catalog",
				Value:   remotePluginCatalogURL,
				EnvVars: []string{"PLUGIN_CATALOG_URL"},
			},
		},
		Commands: []*cli.Command{
			installCommand(),
			listCommand(),
			updateCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
