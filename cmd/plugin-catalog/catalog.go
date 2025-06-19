package main

type Catalog struct {
	Version string      `json:"version"`
	Plugins PluginsType `json:"plugins"`
}

type PluginsType struct {
	Secrets map[string]Plugin `json:"secrets,omitempty"`
	// TODO: Implement other plugin types like auth, database, etc.
}

type Plugin struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Repo     string   `json:"repo"`
	Releases Releases `json:"releases"`
}

type Releases struct {
	Linux  ReleasesArch `json:"linux,omitempty"`
	Darwin ReleasesArch `json:"darwin,omitempty"`
}

type ReleasesArch struct {
	Amd64 Release `json:"amd64,omitempty"`
	Arm64 Release `json:"arm64,omitempty"`
}

type Release struct {
	Url    string `json:"url"`
	Sha256 string `json:"sha256"`
}
