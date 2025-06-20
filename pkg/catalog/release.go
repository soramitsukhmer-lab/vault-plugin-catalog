package catalog

type Release struct {
	ID            string
	PluginName    string
	PluginVersion string
	PluginType    string
	Repo          string
	Url           string
	Sha256        string
}
