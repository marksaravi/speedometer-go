package dashboard

import (
	"math"

	"github.com/marksaravi/devices-go/colors"
)

const (
	BACKGROUNG_COLOR     = colors.ROYALBLUE
	SPEED_CURVE_DASH     = colors.GHOSTWHITE
	SPEED_FORWARD_SAFE   = colors.GREEN
	SPEED_FORWARD_WARN   = colors.YELLOW
	SPEED_FORWARD_DANGER = colors.RED
	SPEED_BACKWARD       = BACKGROUNG_COLOR
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
	SPEED_CURVE_RESOLUTION  float64 = 5           // km/h
	SPEED_CURVE_START_ANGLE float64 = 0           // math.Pi / 2    // 90 degree
	SPEED_CURVE_END_ANGLE   float64 = 2 * math.Pi // 1.67 * math.Pi // 300 degree
	SPEED_CURVE_WIDTH       float64 = VIEW_HEIGHT / 12
	SPEED_CURVE_RADIUS      float64 = VIEW_HEIGHT/2 - SPEED_CURVE_WIDTH/2
	SPEED_CURVE_CENTER_X    float64 = VIEW_LEFT + VIEW_HEIGHT/2
	SPEED_CURVE_CENTER_Y    float64 = VIEW_TOP + VIEW_HEIGHT/2
	SPEED_CURVE_DASH_LEN    float64 = SPEED_CURVE_WIDTH / 4
)

const (
	MAX_SPEED         float64 = 50 // km/h
	MAX_SAFE_SPEED    float64 = 20
	MAX_WARNING_SPEED float64 = 30
)
