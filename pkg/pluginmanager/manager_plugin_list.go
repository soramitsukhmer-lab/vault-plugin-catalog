package pluginmanager

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/catalog"
)

func (p *PluginManager) ListPlugins() error {
	// This method should return a list of plugin names available in the catalog.
	// Implementation will depend on how the catalog is structured (e.g., JSON, YAML).

	catalog, err := catalog.NewCatalog(p.LocalCatalogPath)
	if err != nil {
		return fmt.Errorf("failed to load catalog: %w", err)
	}

	releases, err := catalog.GetReleases()
	if err != nil {
		return fmt.Errorf("failed to get releases from catalog: %w", err)
	}

	fmt.Println("Listing plugins in the catalog...")
	for _, release := range releases {
		fmt.Printf(
			"  - %s (ver=%s type=%s, SHA256=%s)\n",
			release.PluginName, release.PluginVersion, release.PluginType, release.Sha256,
		)
	}

	return nil
}
