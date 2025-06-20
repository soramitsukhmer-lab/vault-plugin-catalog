package pluginmanager

import (
	"fmt"
	"io"
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
	if err := crossDeviceCopy(*downloadedFilePath, pluginInstallPath); err != nil {
		return fmt.Errorf("failed to copy plugin %s to Vault plugin directory: %w", release.PluginName, err)
	}

	// Fix file permissions
	if err := os.Chmod(pluginInstallPath, 0755); err != nil {
		return fmt.Errorf("failed to set permissions for plugin file %s: %w", pluginInstallPath, err)
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

// Using os.Rename() with mounted volume in Docker or similar environments
// will result in a invalid cross-device link error.
func crossDeviceCopy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file from %s to %s: %w", src, dst, err)
	}

	if err := dstFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync destination file %s: %w", dst, err)
	}

	err = os.Remove(src)
	if err != nil {
		return fmt.Errorf("failed to remove source file %s after copy: %w", src, err)
	}

	return nil
}
