package types

type PluginSpec struct {
	Name     string      `json:"name"`
	Version  string      `json:"version"`
	Repo     string      `json:"repo"`
	Releases ReleaseSpec `json:"releases"`
}

type PluginsTypeSpec struct {
	Secrets map[string]PluginSpec `json:"secrets,omitempty"`
	// TODO: Implement other plugin types like auth, database, etc.
}
