package app

type buttonArea struct {
	x1, y1, x2, y2 float64
}

type button struct {
	id       int
	active   bool
	drawable bool
	text     string
	area     buttonArea
}


// func (a *speedoApp) touchToScreenXY(x,y float64) (float64, float64) {

// }