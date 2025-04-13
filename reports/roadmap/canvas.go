package roadmap

import (
	"fmt"

	gs "google.golang.org/api/slides/v1"
)

const (
	CanvasMaxX = float64(700.0)
	CanvasMaxY = float64(500.0)
)

type SlideCanvasInfo struct {
	BoxFgColor      *gs.RgbColor
	BoxBgColor      *gs.RgbColor
	BoxHeight       float64
	BoxMarginBottom float64
	Canvas          CanvasFloat64
}

type CanvasFloat64 struct {
	MinX float64
	MinY float64
	MaxX float64
	MaxY float64
}

func (c64 *CanvasFloat64) ThisX(this, min, max float64) (float64, error) {
	if min > max {
		return 0.0, fmt.Errorf("min (%v) is larger than max (%v)", min, max)
	} else if this < min || this > max {
		return 0.0, fmt.Errorf("this (%v) is not within min,max (%v, %v)", this, min, max)
	}
	diff := max - min
	plus := this - min
	pct := float64(plus) / float64(diff)
	diffCan := c64.MaxX - c64.MinX
	thisPlus := pct * diffCan
	thisX := c64.MinX + thisPlus
	return thisX, nil
}

type location struct {
	SrcAllMinX int64
	SrcAllMaxX int64
	SrcAllWdtX int64
	SrcBoxMinX int64
	SrcBoxMaxX int64
	SrcBoxWdtX int64
	SrcPctWdtX float64
	OutAllMinX float64
	OutAllMaxX float64
	OutAllWdtX float64
	OutBoxMinX float64
	OutBoxMaxX float64
	OutBoxWdtX float64
	BoxOutPctX float64
}
