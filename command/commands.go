package command

import (
	"fmt"
	"os"
	"strings"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/pluginmanager"
	"github.com/urfave/cli/v2"
)

const (
	defaultRemotePluginCatalogURL = "https://raw.githubusercontent.com/soramitsukhmer-lab/vault-plugin-catalog-database/refs/heads/main/catalog.json"
	defaultLocalPluginCatalogPath = "/etc/vault-plugin-catalog/catalog.json"
	defaultVaultPluginDir         = "/vault/plugins"
)

var pm = &pluginmanager.PluginManager{}

var genericFlags = []cli.Flag{
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
	&cli.StringFlag{
		Name:        "vault-plugin-dir",
		Usage:       "Directory where Vault plugins are stored",
		Value:       defaultVaultPluginDir,
		EnvVars:     []string{"VAULT_PLUGIN_DIR"},
		Destination: &pm.VaultPluginDir,
	},
}

func Run(args []string) error {
	app := &cli.App{
		Name:  "vault-plugin-catalog",
		Usage: "A cli tool to manage HashiCorp Vault plugins",
		Commands: []*cli.Command{
			listPluginCommand(pm),
			installPluginCommand(pm),
			uninstallPluginCommand(pm),
			registerPluginCommand(pm),
			updateCatalogCommand(pm),
		},
		Before: func(c *cli.Context) error {
			localCatalogPath := c.String("local-catalog-path")
			if localCatalogPath == "" {
				return fmt.Errorf("local catalog path cannot be empty")
			}

			// Check if local catalog path is exist, if not, create it
			strs := strings.Split(localCatalogPath, "/")
			localCatalogDir := strings.Join(strs[:len(strs)-1], "/")
			if _, err := os.Stat(localCatalogDir); os.IsNotExist(err) {
				if err := os.MkdirAll(localCatalogDir, 0755); err != nil {
					return fmt.Errorf("failed to create local catalog directory: %w", err)
				}
			}

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
	}

	return app.Run(args)
}
