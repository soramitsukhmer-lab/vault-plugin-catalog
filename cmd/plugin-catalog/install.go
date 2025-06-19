package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/urfave/cli/v2"
)

const defaultVaultPluginsDir = "/vault/plugins"

func installCommand() *cli.Command {
	return &cli.Command{
		Name:      "install",
		Usage:     "Install a plugin from the catalog",
		Args:      true,
		ArgsUsage: "<plugin-name>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "plugin-dir",
				Usage: "Directory where the plugin will be installed",
				Value: defaultVaultPluginsDir,
			},
		},
		Action: func(c *cli.Context) error {
			pluginName := c.Args().Get(0)
			if pluginName == "" {
				fmt.Println("Error: Plugin name is required.")
				return fmt.Errorf("plugin name is required")
			}

			localCatalogPath := c.String("local-catalog-path")

			if localCatalogPath == "" {
				fmt.Println("Error: Local catalog path is not specified.")
				return fmt.Errorf("local catalog path is required")
			}

			file, err := os.ReadFile(localCatalogPath)
			if err != nil {
				fmt.Printf("Error reading local catalog file: %v\n", err)
				return err
			}

			var catalog Catalog
			if err := json.Unmarshal(file, &catalog); err != nil {
				fmt.Printf("Error parsing local catalog file: %v\n", err)
				return err
			}

			// TODO: Add support for other plugin types (e.g., auth, database)

			if plugin, ok := catalog.Plugins.Secrets[pluginName]; ok {
				pluginRelease, err := lookupPluginRelease(&plugin)
				if err != nil {
					fmt.Printf("Error getting download URL for plugin %s: %v\n", pluginName, err)
					return err
				}

				fmt.Printf("Installing \"%s-%s\", sha256 \"%s\"...\n", pluginName, plugin.Version, pluginRelease.Sha256)

				pluginDir := c.String("plugin-dir")
				if pluginDir == "" {
					pluginDir = defaultVaultPluginsDir
				}

				// Create the plugin directory if it doesn't exist
				if err := os.MkdirAll(pluginDir, 0755); err != nil {
					fmt.Printf("Error creating plugin directory %s: %v\n", pluginDir, err)
					return err
				}

				// Download the plugin binary
				pluginBinaryPath := fmt.Sprintf("%s/%s", pluginDir, plugin.Name)
				if err := downloadFile(pluginBinaryPath, pluginRelease.Url); err != nil {
					fmt.Printf("Error downloading plugin %s: %v\n", plugin.Name, err)
					return err
				}

				// Set the executable permission
				if err := os.Chmod(pluginBinaryPath, 0755); err != nil {
					fmt.Printf("Error setting executable permission for %s: %v\n", pluginBinaryPath, err)
					return err
				}

				fmt.Printf("Plugin %s installed successfully at %s\n", plugin.Name, pluginBinaryPath)

			}

			return nil
		},
	}
}

func lookupPluginRelease(plugin *Plugin) (*Release, error) {
	var releasesArch ReleasesArch
	switch runtime.GOOS {
	case "linux":
		releasesArch = plugin.Releases.Linux
	case "darwin":
		releasesArch = plugin.Releases.Darwin
	default:
		return nil, fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	var release Release
	switch runtime.GOARCH {
	case "amd64":
		release = releasesArch.Amd64
	case "arm64":
		release = releasesArch.Arm64
	default:
		return nil, fmt.Errorf("unsupported architecture: %s", runtime.GOARCH)
	}

	return &release, nil
}

func downloadFile(filepath, url string) error {
	// This function should implement the logic to download a file from the given URL
	// and save it to the specified filepath. For now, we will just print the action.
	fmt.Printf("Downloading from %s to %s...\n", url, filepath)
	// Actual download logic would go here (e.g., using http.Get and writing to a file)

	resq, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return err
	}
	defer resq.Body.Close()

	if resq.StatusCode != http.StatusOK {
		fmt.Printf("Error: received status code %d while downloading file\n", resq.StatusCode)
		return fmt.Errorf("failed to download file: status code %d", resq.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", filepath, err)
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resq.Body)
	if err != nil {
		fmt.Printf("Error writing to file %s: %v\n", filepath, err)
		return err
	}

	fmt.Print("Download completed successfully.\n")

	return nil
}
