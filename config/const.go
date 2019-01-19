package config

import (
	"github.com/exfly/dockerc/utils"
)

var (
	DockercWorkspace   = "/dockerc"
	CURRENTCONTATINNER = ""
)

func init() {
	CURRENTCONTATINNER = utils.RS()
}
