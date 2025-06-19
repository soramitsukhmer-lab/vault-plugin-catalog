package catalog

import (
	"fmt"
	"runtime"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/types"
)

type Catalog struct {
	*types.Catalog
}

func (c *Catalog) GetReleaseByName(pluginName string) (*types.Release, error) {
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

	return &release, nil
}
