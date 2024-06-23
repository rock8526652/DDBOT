package huya

import (
	"github.com/rock8526652/DDBOT/lsp/concern"
)

type GroupConcernConfig struct {
	concern.IConfig
}

func NewGroupConcernConfig(g concern.IConfig) *GroupConcernConfig {
	return &GroupConcernConfig{g}
}
