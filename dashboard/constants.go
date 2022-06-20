package dashboard

import "github.com/marksaravi/devices-go/colors/rgb565"

// RGB565 colors
const (
	FLORECENT_GREEN rgb565.RGB565 = 0x1FE3
	LAWN_GREEN      rgb565.RGB565 = 0x7FC0
	ROYAL_BLUE      rgb565.RGB565 = 0x3B3B
	GHOST_WHITE     rgb565.RGB565 = 0xF7BF
)

// colors
const (
	backgroung_color  rgb565.RGB565 = LAWN_GREEN
	speed_curve_1kmph rgb565.RGB565 = GHOST_WHITE
)

// basic dimensions
const (
	SCREEN_WIDTH   float64 = 320
	SCREEN_HEIGHT  float64 = 240
	SCREEN_PADDING float64 = 5
	VIEW_WIDTH             = SCREEN_WIDTH - SCREEN_PADDING*2
	VIEW_HEIGHT            = SCREEN_HEIGHT - SCREEN_PADDING*2
	VIEW_LEFT              = SCREEN_PADDING
	VIEW_TOP               = SCREEN_PADDING
)

const (
	MAX_SPEED                 float64 = 60  // km/h
	SPEED_RESOLUTION          float64 = 2   // km/h
	SPEED_CURVE_START_ANGLE   float64 = 90  // degree
	SPEED_CURVE_END_ANGLE     float64 = 300 // degree
	SPEED_CURVE_CENTER_X      float64 = 120
	SPEED_CURVE_CENTER_Y      float64 = 120
	SPEED_CURVE_CENTER_RADIUS float64 = 110
	SPEED_CURVE_CENTER_WIDTH  float64 = 20
)
