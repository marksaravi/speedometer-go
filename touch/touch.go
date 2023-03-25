package touch

import (
	"github.com/marksaravi/speedometer-go/models"
)

type touch struct {
	touched <-chan models.XY
}

func NewTouch(touched chan models.XY) *touch {
	return &touch{
		touched: touched,
	}
}

func (t *touch) Touched() <-chan models.XY {
	return t.touched
}
