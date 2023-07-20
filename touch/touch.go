package touch

import (
	"github.com/marksaravi/speedometer-go/models"
)

type touch struct {
	touched <-chan models.XY
	ax, bx, ay, by float64
}

func NewTouch(touched chan models.XY, x1t, y1t, x2t, y2t, x1s, y1s, x2s, y2s float64) *touch {
	ax, bx := ConverionFactors(x1t, x2t, x1s, x2s)
	ay, by := ConverionFactors(y1t, y2t, y1s, y2s)
	return &touch{
		touched: touched,
		ax: ax,
		bx: bx,
		ay: ay,
		by: by,
	}
}

func (t *touch) Touched() <-chan models.XY {
	return t.touched
}

func (t *touch) TouchConvert(xy models.XY) (float64, float64) {
	x, y := Convert(xy.X, xy.Y, t.ax, t.bx, t.ay, t.by)
	return x, y
}

func ConverionFactors(x1t, x2t, x1s, x2s float64) (float64, float64) {
	a := (x2s-x1s)/(x2t-x1t)
	b := x1s - a*x1t
	return a, b
}

func Convert(xt,yt, ax, bx, ay, by float64 ) (float64, float64) {
	xs := xt * ax + bx
	ys := yt * ay + by
	return xs, ys
}