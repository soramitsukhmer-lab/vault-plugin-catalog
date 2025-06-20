package catalog

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func (c *CatalogManager) GetReleaseByName(pluginName string) (*Release, error) {
	var plugin types.PluginSpec
	var pluginType string
	if p, ok := c.Catalog.Plugins.Secrets[pluginName]; ok {
		plugin = p
		pluginType = "secret"
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
		ID:            pluginName,
		PluginName:    plugin.Name,
		PluginVersion: plugin.Version,
		PluginType:    pluginType,
		Repo:          plugin.Repo,
		Url:           release.Url,
		Sha256:        release.Sha256,
	}

	return r, nil
}

func (r *CatalogManager) DownloadPlugin(release *Release) (*string, error) {
	// This method should handle the download of the plugin release.
	// It will likely involve downloading the plugin from the URL and verifying its SHA256 checksum.
	// For now, we will just print a message indicating the download process.

	// Example implementation:
	// 1. Download the file from release.Url
	// 2. Verify the SHA256 checksum
	// 3. Save the file to the appropriate directory

	fmt.Printf("Downloading plugin %s-%s from %s...\n", release.PluginName, release.PluginVersion, release.Repo)

	resp, err := http.Get(release.Url)
	if err != nil {
		return nil, fmt.Errorf("failed to download plugin %s: %w", release.PluginName, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download plugin %s: received status code %d", release.PluginName, resp.StatusCode)
	}

	file, err := os.CreateTemp(os.TempDir(), release.PluginName+"-*")
	tmpFileName := file.Name()
	if err != nil {
		return nil, fmt.Errorf("failed to create output file %s: %w", tmpFileName, err)
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return nil, fmt.Errorf("failed to write plugin data to file %s: %w", tmpFileName, err)
	}

	fmt.Printf("Plugin %s downloaded successfully to %s\n", release.PluginName, tmpFileName)

	return &tmpFileName, nil
}

func (r *CatalogManager) DownloadPluginByName(name string) (*string, error) {
	release, err := r.GetReleaseByName(name)
	if err != nil {
		return nil, fmt.Errorf("failed to get release for plugin %s: %w", name, err)
	}

	return r.DownloadPlugin(release)
}
