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
	BACKGROUNG_COLOR rgb565.RGB565 = ROYAL_BLUE
	SPEED_CURVE_DASH rgb565.RGB565 = GHOST_WHITE
)

// basic dimensions
const (
	SCREEN_WIDTH   float64 = 320
	SCREEN_HEIGHT  float64 = 240
	SCREEN_PADDING float64 = SCREEN_HEIGHT / 48
	VIEW_WIDTH             = SCREEN_WIDTH - SCREEN_PADDING*2
	VIEW_HEIGHT            = SCREEN_HEIGHT - SCREEN_PADDING*2
	VIEW_LEFT              = SCREEN_PADDING
	VIEW_TOP               = SCREEN_PADDING
)

// speed curve
const (
	SPEED_CURVE_RESOLUTION  float64 = 5   // km/h
	SPEED_CURVE_START_ANGLE float64 = 90  // degree
	SPEED_CURVE_END_ANGLE   float64 = 300 // degree
	SPEED_CURVE_WIDTH       float64 = VIEW_HEIGHT / 12
	SPEED_CURVE_RADIUS      float64 = VIEW_HEIGHT/2 - SPEED_CURVE_WIDTH/2
	SPEED_CURVE_CENTER_X    float64 = VIEW_LEFT + VIEW_HEIGHT/2
	SPEED_CURVE_CENTER_Y    float64 = VIEW_TOP + VIEW_HEIGHT/2
	SPEED_CURVE_DASH_LEN    float64 = SPEED_CURVE_WIDTH / 4
)

const (
	MAX_SPEED float64 = 60 // km/h
)
