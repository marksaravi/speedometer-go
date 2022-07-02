package dashboard

import (
	"time"

	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
)

// basic dimensions
const (
	TOP_MARGIN           int = 0
	LEFT_MARGIN          int = 16
	SPEED_LINE_HEIGHT    int = 64
	DISTANCE_LINE_HEIGHT int = 64
	DURATION_LINE_HEIGHT int = 64
	LABEL_COLUMN         int = 0
	DATA_COLUMN          int = 120
)

// Fonts
var (
	LABEL_FONT          fonts.BitmapFont = fonts.FreeSans9pt7b
	SPEED_LABEL_FONT    fonts.BitmapFont = LABEL_FONT
	DISTANCE_LABEL_FONT fonts.BitmapFont = LABEL_FONT
	DURATION_LABEL_FONT fonts.BitmapFont = LABEL_FONT

	SPEED_DATA_FONT    fonts.BitmapFont = fonts.FreeSans24pt7b
	DISTANCE_DATA_FONT fonts.BitmapFont = fonts.FreeSans18pt7b
	DURATION_DATA_FONT fonts.BitmapFont = fonts.FreeSans18pt7b
)

const (
	SPEED_RESOLUTION    float64       = 0.5
	DISTANCE_RESOLUTION float64       = 10
	DURATION_RESOLUTION time.Duration = time.Second
	SEC_CHANGED         TimeChanged   = 0
	MIN_CHANGED         TimeChanged   = 1
	HOUR_CHANGED        TimeChanged   = 2
)

const (
	LIGHT_LABEL_COLOR = colors.BLACK
	LIGHT_DATA_COLOR  = colors.RED

	DARK_LABEL_COLOR = colors.WHITE
	DARK_DATA_COLOR  = colors.YELLOW
)

var LightTheme = Theme{
	BackgroungColor:    colors.WHITE,
	SpeedLabelColor:    LIGHT_LABEL_COLOR,
	DistanceLabelColor: LIGHT_LABEL_COLOR,
	DurationLabelColor: LIGHT_LABEL_COLOR,
	SpeedDataColor:     LIGHT_DATA_COLOR,
	DistanceDataColor:  LIGHT_DATA_COLOR,
	DurationDataColor:  LIGHT_DATA_COLOR,
}

var DarkTheme = Theme{
	BackgroungColor:    colors.DARKGRAY,
	SpeedLabelColor:    DARK_LABEL_COLOR,
	DistanceLabelColor: DARK_LABEL_COLOR,
	DurationLabelColor: DARK_LABEL_COLOR,
	SpeedDataColor:     DARK_DATA_COLOR,
	DistanceDataColor:  DARK_DATA_COLOR,
	DurationDataColor:  DARK_DATA_COLOR,
}
