package http2

import (
	"wmaster/controller"
	"wmaster/db"
)

type Http struct{}

var d db.MongoDB
var c controller.Ctrl

func init() {
	d = db.MongoDB{"mongodb://localhost:27020", "WServiceMaster", "Master"}
	c = controller.Ctrl{}
}
