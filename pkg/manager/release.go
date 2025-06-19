package manager

import "github.com/soramitsukhmer-lab/vault-plugin-catalog/pkg/types"

type Release struct {
	PluginName    string
	PluginVersion string
	Release       *types.Release
}
