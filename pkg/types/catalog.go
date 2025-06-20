package types

type CatalogSpec struct {
	Version string          `json:"version"`
	Plugins PluginsTypeSpec `json:"plugins"`
}
