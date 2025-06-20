package manager

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/types"
)

type CatalogManager struct {
	Catalog *types.CatalogSpec
}

// NewCatalog initializes a new Catalog instance with the provided catalog data.
func NewCatalog(filepath string) (*CatalogManager, error) {
	if filepath == "" {
		return nil, fmt.Errorf("catalog file path cannot be empty")
	}

	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read catalog file: %w", err)
	}

	catalog := types.CatalogSpec{}
	if err := json.Unmarshal(file, &catalog); err != nil {
		return nil, fmt.Errorf("failed to unmarshal catalog data: %w", err)
	}

	cm := &CatalogManager{
		Catalog: &catalog,
	}

	return cm, nil
}

func (c *CatalogManager) GetReleaseByName(pluginName string) (*Release, error) {
	var plugin types.PluginSpec
	var pluginType string
	if p, ok := c.Catalog.Plugins.Secrets[pluginName]; ok {
		plugin = p
		pluginType = "secrets"
	} else {
		return nil, fmt.Errorf("plugin %s not found in secrets catalog", pluginName)
	}

	var platformRelease types.ReleasePlatformSpec
	switch runtime.GOOS {
	case "linux":
		platformRelease = plugin.Releases.Linux
	case "darwin":
		platformRelease = plugin.Releases.Darwin
	default:
		return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	var release types.ReleaseArchitectureSpec
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
		PluginType:    pluginType,
		Url:           release.Url,
		Sha256:        release.Sha256,
	}

	return r, nil
}

func (c *CatalogManager) GetReleases() ([]Release, error) {
	var releases []Release

	for pluginName := range c.Catalog.Plugins.Secrets {
		r, err := c.GetReleaseByName(pluginName)
		if err != nil {
			return nil, fmt.Errorf("failed to get release for plugin %s: %w", pluginName, err)
		}
		if r != nil {
			releases = append(releases, *r)
		}
	}

	return releases, nil
}
