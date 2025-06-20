package pluginmanager

type PluginManager struct {
	RemoteCatalogURL string
	LocalCatalogPath string
	VaultPluginDir   string
}

func NewPluginManager(remoteCatalogURL, localCatalogPath, vaultPluginDir string) *PluginManager {
	return &PluginManager{
		RemoteCatalogURL: remoteCatalogURL,
		LocalCatalogPath: localCatalogPath,
		VaultPluginDir:   vaultPluginDir,
	}
}
