package types

type Plugin struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Repo     string   `json:"repo"`
	Releases Releases `json:"releases"`
}

type PluginsType struct {
	Secrets map[string]Plugin `json:"secrets,omitempty"`
	// TODO: Implement other plugin types like auth, database, etc.
}
