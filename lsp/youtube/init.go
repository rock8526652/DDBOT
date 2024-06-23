package youtube

import (
	"github.com/rock8526652/DDBOT/lsp/concern"
)

func init() {
	concern.RegisterConcern(NewConcern(concern.GetNotifyChan()))
}
