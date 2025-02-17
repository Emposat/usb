package usb

//go:generate go run -tags generate gen.go

import (
	"github.com/Emposat/usb/gusb"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/text"
)

func init() {
	log.SetHandler(cli.Default) // nice, but also not nice
	log.SetHandler(text.Default)
	log.SetLevel(log.InfoLevel)
	gusb.SetLogger(log.Log.(*log.Logger))
}
