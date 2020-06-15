package palette

import (
	"image/color/palette"

	"github.com/qiniu/goplus/gop"
)

// I is a Go package instance.
var I = gop.NewGoPackage("image/color/palette")

func init() {
	I.RegisterVars(
		I.Var("Plan9", &palette.Plan9),
		I.Var("WebSafe", &palette.WebSafe),
	)
}
