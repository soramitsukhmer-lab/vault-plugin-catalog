package command

import (
	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/pluginmanager"
	"github.com/soramitsukhmer-lab/vault-plugin-catalog/version"
	"github.com/urfave/cli/v2"
)

const (
	defaultRemotePluginCatalogURL = "https://raw.githubusercontent.com/soramitsukhmer-lab/vault-plugin-catalog-database/refs/heads/main/catalog.json"
	defaultVaultPluginDir         = "/vault/plugins"
	defaultLocalPluginCatalogPath = defaultVaultPluginDir + "/catalog.json"
)

var pm = &pluginmanager.PluginManager{}

var genericFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "vault-plugin-dir",
		Usage:       "Directory where Vault plugins are stored",
		Value:       defaultVaultPluginDir,
		EnvVars:     []string{"VAULT_PLUGIN_DIR"},
		Destination: &pm.VaultPluginDir,
	},
	&cli.StringFlag{
		Name:        "local-catalog-path",
		Usage:       "Path to the local plugin catalog file",
		Value:       defaultLocalPluginCatalogPath,
		EnvVars:     []string{"LOCAL_PLUGIN_CATALOG_PATH"},
		Destination: &pm.LocalCatalogPath,
	},
	&cli.StringFlag{
		Name:        "remote-catalog-url",
		Usage:       "URL of the remote plugin catalog",
		Value:       defaultRemotePluginCatalogURL,
		EnvVars:     []string{"REMOTE_PLUGIN_CATALOG_URL"},
		Destination: &pm.RemoteCatalogURL,
	},
}

func Run(args []string) error {
	app := &cli.App{
		Name:    "vault-plugin-catalog",
		Usage:   "A cli tool to manage HashiCorp Vault plugins",
		Version: version.Version,
		Authors: []*cli.Author{
			{Name: "Soramitsu Khmer Lab", Email: "https://github.com/soramitsukhmer-lab"},
			{Name: "Socheat Sok", Email: "https://github.com/socheatsok78"},
		},
		Commands: []*cli.Command{
			listPluginCommand(pm),
			installPluginCommand(pm),
			uninstallPluginCommand(pm),
			registerPluginCommand(pm),
			updateCatalogCommand(pm),
		},
	}

	return app.Run(args)
}
