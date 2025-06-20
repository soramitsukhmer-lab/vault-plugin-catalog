package command

import (
	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/manager"
	"github.com/urfave/cli/v2"
)

const (
	defaultRemotePluginCatalogURL = "https://raw.githubusercontent.com/soramitsukhmer-lab/vault-plugin-catalog/refs/heads/main/catalog.json"
	defaultLocalPluginCatalogPath = "/etc/vault-plugin-catalog/catalog.json"
	defaultVaultPluginDir         = "/etc/vault/plugins"
)

func Run(args []string) error {
	manager := &manager.PluginManager{}

	app := &cli.App{
		Name:  "vault-plugin-catalog",
		Usage: "A cli tool to manage HashiCorp Vault plugins",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "local-catalog-path",
				Usage:       "Path to the local plugin catalog file",
				Value:       defaultLocalPluginCatalogPath,
				EnvVars:     []string{"LOCAL_PLUGIN_CATALOG_PATH"},
				Destination: &manager.LocalCatalogPath,
			},
			&cli.StringFlag{
				Name:        "remote-catalog-url",
				Usage:       "URL of the remote plugin catalog",
				Value:       defaultRemotePluginCatalogURL,
				EnvVars:     []string{"REMOTE_PLUGIN_CATALOG_URL"},
				Destination: &manager.RemoteCatalogURL,
			},
			&cli.StringFlag{
				Name:        "vault-plugin-dir",
				Usage:       "Directory where Vault plugins are stored",
				Value:       defaultVaultPluginDir,
				EnvVars:     []string{"VAULT_PLUGIN_DIR"},
				Destination: &manager.VaultPluginDir,
			},
		},
		Commands: []*cli.Command{
			listPluginCommand(manager),
			installPluginCommand(manager),
			uninstallPluginCommand(manager),
			registerPluginCommand(manager),
			updateCatalogCommand(manager),
		},
	}

	return app.Run(args)
}
