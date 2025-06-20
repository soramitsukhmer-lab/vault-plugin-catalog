package pluginmanager

import (
	"fmt"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/catalog"
)

var message = ` Prepare plugin registration for "%s" plugin:

To register a plugin to HashiCorp Vault, please run the following command:
$ vault plugin register -version=%s -sha256=%s -command=%s %s %s
`

func (p *PluginManager) RegisterPlugin(name string) error {
	// This method should handle the registration of a plugin by its name.
	// It will likely involve checking if the plugin exists in the remote catalog
	// and then registering it with Vault.

	catalog, err := catalog.NewCatalog(p.LocalCatalogPath)
	if err != nil {
		return fmt.Errorf("failed to load catalog: %w", err)
	}

	release, err := catalog.GetReleaseByName(name)
	if err != nil {
		return fmt.Errorf("failed to find plugin %s in catalog: %w", name, err)
	}

	pluginFileName := fmt.Sprintf("%s-%s", release.PluginName, release.PluginVersion)
	fmt.Printf(
		message,
		release.ID,            // <plugin-id>
		release.PluginVersion, // -version
		release.Sha256,        // -sha256
		pluginFileName,        // -command
		release.PluginType,    // <plugin-type>
		release.PluginName,    // <plugin-name>
	)

	return nil
}
