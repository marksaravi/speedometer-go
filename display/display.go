package display

import (
	"fmt"
	"log"
	"time"

	"github.com/marksaravi/drawings-go/drawings"
	"github.com/marksaravi/drivers-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
	"github.com/marksaravi/speedometer-go/themes"
)

const (
	RESET_BUTTON = 10
)

type area struct {
	x1, x2, y1, y2 float64
}

type button struct {
	id      int
	visible bool
	area    area
	label   string
}

type display struct {
	resetChannel          chan bool
	sketcher              drawings.Sketcher
	theme                 themes.Theme
	xs, ys, width, height float64
	speedArea             *area
	distanceArea          *area
	durationArea          *area
	buttons               map[int]button
}

func NewDisplay(theme themes.Theme, sketcher drawings.Sketcher, margin float64) *display {
	resetChannel := make(chan bool)
	sketcher.SetRotation(drawings.ROTATION_90)
	resetArea := area{x1: 40, y1: 100, x2: 200, y2: 160}
	resetButton := button{
		id: RESET_BUTTON, label: "Reset", area: resetArea, visible: false,
	}
	buttons := make(map[int]button)
	buttons[RESET_BUTTON] = resetButton
	return &display{
		theme:        theme,
		resetChannel: resetChannel,
		sketcher:     sketcher,
		xs:           margin,
		ys:           margin,
		width:        float64(sketcher.ScreenWidth()) - 2*margin,
		height:       float64(sketcher.ScreenHeight()) - 2*margin,
		buttons:      buttons,
	}
}

func (d *display) Initialize() {
	d.sketcher.Clear(colors.BLACK)
	d.sketcher.ClearArea(d.xs, d.ys, d.xs+d.width, d.ys+d.height, d.theme.BackgroungColor)
	d.writeLabels()
	d.sketcher.Update()
}

func (d *display) SetInfo(speed float64, distance float64, duration time.Duration) {
	d.writeSpeed(speed)
	d.writeDistance(distance)
	d.writeDuration(duration)
	if d.buttons[RESET_BUTTON].visible {
		d.drawButton(RESET_BUTTON)
	}
	d.sketcher.Update()
}

func (d *display) setButton(id int, visible bool) {
	btn := d.buttons[id]
	btn.visible = visible
	d.buttons[id] = btn
}

func (d *display) Touched(x, y float64) bool {
	tapped := false
	btn := d.buttons[RESET_BUTTON]
	if !btn.visible {
		d.setButton(RESET_BUTTON, true)
	} else {
		if btn.isTapped(x, y) {
			tapped = true
			fmt.Println("RESET")
		}
		d.setButton(RESET_BUTTON, false)
		d.clearButton(RESET_BUTTON)
	}
	log.Printf("TOUCH: %f,%f\n", x, y)
	return tapped
}

func (d *display) ResetChannel() <-chan bool {
	return d.resetChannel
}

func (d *display) writeLabels() {
	d.write("Distance", fonts.FreeMono9pt7b, d.theme.DurationLabelColor, 4, 20, 1, 1, nil)
	d.write("Speed (km/h)", fonts.FreeMono9pt7b, d.theme.DurationLabelColor, 4, 80, 1, 1, nil)
	d.write("Duration", fonts.FreeMono9pt7b, d.theme.DurationLabelColor, 4, 260, 1, 1, nil)
	d.write("km/h", fonts.FreeMono9pt7b, d.theme.DurationLabelColor, 185, 250, 1, 1, nil)
}

func (d *display) setArea(text string, x, y, xScale, yScale float64, a *area) *area {
	x1, y1, x2, y2 := d.sketcher.GetTextArea(x, y, text, xScale, yScale)
	if a == nil {
		a = &area{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		}
	} else {
		if x1 < a.x1 {
			a.x1 = x1
		}
		if y1 < a.y1 {
			a.y1 = y1
		}
		if x2 > a.x2 {
			a.x2 = x2
		}
		if y2 > a.y2 {
			a.y2 = y2
		}
	}
	return a
}

func (d *display) writeSpeed(speed float64) {
	const x float64 = 20
	const y float64 = 230
	const xScale = 2
	const yScale = 4

	text := fmt.Sprintf("%0.1f", speed)
	d.speedArea = d.write(text, fonts.FreeSans24pt7b, d.theme.SpeedColor, x, y, xScale, yScale, d.speedArea)
}

func (d *display) writeDistance(distance float64) {
	const x float64 = 20
	const y float64 = 60
	const xScale = 1
	const yScale = 1

	text := fmt.Sprintf("%0.1f", distance)
	d.distanceArea = d.write(text, fonts.FreeSans24pt7b, d.theme.SpeedColor, x, y, xScale, yScale, d.distanceArea)
}

func (d *display) writeDuration(duration time.Duration) {
	const x float64 = 20
	const y float64 = 305
	const xScale = 1
	const yScale = 1

	hour := int(duration.Seconds() / 3600)
	minute := int(duration.Seconds()/60) % 60
	second := int(duration.Seconds()) % 60
	text := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	d.durationArea = d.write(text, fonts.FreeSans24pt7b, d.theme.SpeedColor, x, y, xScale, yScale, d.durationArea)
}

func (d *display) write(text string, font fonts.BitmapFont, color colors.Color, x, y float64, xScale, yScale float64, area *area) *area {
	d.sketcher.SetFont(font)
	a := d.setArea(text, x, y, xScale, yScale, area)
	d.sketcher.ClearArea(a.x1, a.y1, a.x2, a.y2, d.theme.BackgroungColor)
	d.sketcher.MoveCursor(x, y)
	d.sketcher.WriteScaled(text, xScale, yScale, d.theme.SpeedColor)
	return a
}

func (d *display) calibrationPoints() {
	var PADDING float64 = 60
	w := d.sketcher.ScreenWidth()
	h := d.sketcher.ScreenHeight()
	d.sketcher.FillCircle(PADDING, PADDING, float64(5), colors.RED)
	d.sketcher.FillCircle(w-PADDING, PADDING, float64(5), colors.RED)
	d.sketcher.FillCircle(PADDING, h-PADDING, float64(5), colors.RED)
	d.sketcher.FillCircle(w-PADDING, h-PADDING, float64(5), colors.RED)
	fmt.Printf("(%5.0f,%5.0f),(%5.0f,%5.0f),(%5.0f,%5.0f),(%5.0f,%5.0f)",
		PADDING, PADDING, w-PADDING, PADDING, PADDING, h-PADDING, w-PADDING, h-PADDING)
}

func (d *display) drawButton(id int) {
	btn := d.buttons[id]
	d.sketcher.FillRectangle(btn.area.x1, btn.area.y1, btn.area.x2, btn.area.y2, colors.YELLOW)
	d.sketcher.SetFont(fonts.FreeSans24pt7b)
	d.sketcher.MoveCursor(btn.area.x1+20, btn.area.y2-14)
	d.sketcher.WriteScaled(btn.label, 1, 1, colors.RED)
}

func (d *display) clearButton(id int) {
	btn := d.buttons[id]
	d.sketcher.FillRectangle(btn.area.x1, btn.area.y1, btn.area.x2, btn.area.y2, d.theme.BackgroungColor)
}

func (b *button) isTapped(x, y float64) bool {
	return x > b.area.x1 && x < b.area.x2 && y > b.area.y1 && y < b.area.y2
}
