package manager

import (
	"fmt"
	"runtime"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/types"
)

type Catalog struct {
	*types.Catalog
}

// NewCatalog initializes a new Catalog instance with the provided catalog data.
func NewCatalog(catalog *types.Catalog) *Catalog {
	return &Catalog{
		Catalog: catalog,
	}
}

func (c *Catalog) GetReleaseByName(pluginName string) (*Release, error) {
	var plugin types.Plugin
	if p, ok := c.Plugins.Secrets[pluginName]; ok {
		plugin = p
	} else {
		return nil, fmt.Errorf("plugin %s not found in secrets catalog", pluginName)
	}

	var platformRelease types.ReleasesPlatform
	switch runtime.GOOS {
	case "linux":
		platformRelease = plugin.Releases.Linux
	case "darwin":
		platformRelease = plugin.Releases.Darwin
	default:
		return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	var release types.Release
	switch runtime.GOARCH {
	case "amd64":
		release = platformRelease.Amd64
	case "arm64":
		release = platformRelease.Arm64
	default:
		return nil, fmt.Errorf("unsupported architecture: %s", runtime.GOARCH)
	}

	r := &Release{
		PluginName:    pluginName,
		PluginVersion: plugin.Version,
		Release:       &release,
	}

	return r, nil
}
