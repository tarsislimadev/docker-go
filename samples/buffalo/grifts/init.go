package grifts

import (
	"buffalo1/actions"
	
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
