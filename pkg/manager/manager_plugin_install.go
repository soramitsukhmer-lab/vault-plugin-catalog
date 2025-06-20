package manager

import (
	"fmt"
	"os"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/catalog"
)

func (p *PluginManager) InstallPlugin(name string) error {
	// This method should handle the installation of a plugin by its name.
	// It will likely involve downloading the plugin from the remote catalog
	// and placing it in the specified Vault plugin directory.

	catalog, err := catalog.NewCatalog(p.LocalCatalogPath)
	if err != nil {
		return fmt.Errorf("failed to load catalog: %w", err)
	}

	release, err := catalog.GetReleaseByName(name)
	if err != nil {
		return fmt.Errorf("failed to find plugin %s in catalog: %w", name, err)
	}

	downloadedFilePath, err := catalog.DownloadPlugin(release)
	if err != nil {
		return fmt.Errorf("failed to download plugin %s: %w", name, err)
	}

	pluginFileName := fmt.Sprintf("%s-%s", release.PluginName, release.PluginVersion)
	pluginInstallPath := fmt.Sprintf("%s/%s", p.VaultPluginDir, pluginFileName)

	fmt.Printf("Installing plugin: %s to %s...\n", pluginFileName, pluginInstallPath)
	if err := os.Rename(*downloadedFilePath, pluginInstallPath); err != nil {
		return fmt.Errorf("failed to copy plugin %s to Vault plugin directory: %w", release.PluginName, err)
	}

	return nil
}

func (p *PluginManager) InstallPluginWithRegistration(name string) error {
	// This method should handle the installation of a plugin and its registration with Vault.
	if err := p.InstallPlugin(name); err != nil {
		return fmt.Errorf("failed to install plugin %s: %w", name, err)
	}

	// Here you would typically call a method to register the plugin with Vault.
	// For example:
	// if err := p.RegisterPlugin(name); err != nil {
	// 	return fmt.Errorf("failed to register plugin %s: %w", name, err)
	// }

	fmt.Println("TODO: InstallPluginWithRegistration method is not implemented yet.")
	return nil
}
