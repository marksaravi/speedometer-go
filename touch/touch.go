package touch

type XY struct {
	X, Y float64
}

type touch struct {
	touched chan XY
}

func NewTouch() *touch {
	touched := make(chan XY)
	return &touch{
		touched: touched,
	}
}

func (t *touch) Touched() <-chan XY {
	return t.touched
}
