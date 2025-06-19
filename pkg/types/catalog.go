package types

type Catalog struct {
	Version string      `json:"version"`
	Plugins PluginsType `json:"plugins"`
}
