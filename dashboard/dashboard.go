package dashboard

import (
	"fmt"

	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/devices-go/devices/display"
	"github.com/marksaravi/fonts-go/fonts"
)

type clearZone struct {
	x1, y1, x2, y2 int
}
type dashboardDisplay struct {
	display    display.RGBDisplay
	theme      Theme
	clearZones []clearZone
}

func NewDashboardDisplay(display display.RGBDisplay) *dashboardDisplay {
	return &dashboardDisplay{
		display: display,
		theme:   DarkTheme,
		clearZones: []clearZone{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}
}

func (d *dashboardDisplay) Initialise() {
	d.initBackground()
}

func (d *dashboardDisplay) UpdateSpeed(speed float64) {
	x := DATA_X
	y := SPEED_DATA_LINE_Y
	d.printDigits(fmt.Sprintf("%3.1f", speed), SPEED_DATA_FONT, d.theme.SpeedDataColor, x, y, 0)
}

func (d *dashboardDisplay) UpdateDistance(distance float64) {
	x := DATA_X
	y := DISTANCE_DATA_LINE_Y
	d.printDigits(fmt.Sprintf("%4.2f", distance/1000), DISTANCE_DATA_FONT, d.theme.DistanceDataColor, x, y, 1)
}

func (d *dashboardDisplay) UpdateDuration(t int, change int) {
	digits := fmt.Sprintf("%02d", t)
	x := DATA_DUR_X
	y := DURATION_DATA_LINE_Y
	clearIndex := 2
	DIGIT_OFFSET := TIME_DIGIT_WIDTH + TIME_COLON_WIDTH
	if change == MINUTE_CHANGED {
		x += DIGIT_OFFSET
		clearIndex = 3
	}
	if change == SECOND_CHANGED {
		x += 2 * DIGIT_OFFSET
		clearIndex = 4
	}
	d.printDigits(digits, DURATION_DATA_FONT, d.theme.DurationDataColor, x, y, clearIndex)
}

func (d *dashboardDisplay) UpdateDisplay() {
	d.display.Update()
}

func (d *dashboardDisplay) printDigits(digits string, font fonts.BitmapFont, color colors.Color, x, y, clearIndex int) {
	d.setTextSettings(font, color, x, y)
	d.clearArea(digits, x, y, clearIndex)
	d.display.Write(digits)
}

func (d *dashboardDisplay) clearArea(text string, x, y, clearZoneIndex int) {
	cz := d.clearZones[clearZoneIndex]
	d.display.ClearArea(float64(x+cz.x1), float64(y+cz.y1), float64(x+cz.x2), float64(y+cz.y2))
	x1, y1, x2, y2 := d.display.GetTextArea(text)
	d.clearZones[clearZoneIndex] = clearZone{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func (d *dashboardDisplay) printDurationColons() {
	x := DATA_DUR_X + TIME_DIGIT_WIDTH + TIME_COLON_OFFSET
	y := DURATION_DATA_LINE_Y
	d.writeText(
		":",
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		x,
		y,
	)
	d.writeText(
		":",
		DURATION_DATA_FONT,
		d.theme.DurationDataColor,
		x+TIME_DIGIT_WIDTH+TIME_COLON_WIDTH,
		y,
	)
}
