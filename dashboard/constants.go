package dashboard

import (
	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
)

// basic dimensions
const (
	SEG_HEIGHT            int = 24
	SEG_WIDTH             int = 32
	LABEL_X               int = 16
	LABEL_GAP_M           int = 2 * SEG_HEIGHT
	SPEED_LABEL_LINE_Y    int = 4*SEG_HEIGHT - 12
	DISTANCE_LABEL_LINE_Y int = SPEED_LABEL_LINE_Y + LABEL_GAP_M
	DURATION_LABEL_LINE_Y int = DISTANCE_LABEL_LINE_Y + LABEL_GAP_M
	DATA_X                int = SEG_WIDTH*5 + 2
	DATA_DUR_X            int = SEG_WIDTH*5 + 2
	SPEED_DATA_LINE_Y     int = SPEED_LABEL_LINE_Y + 8
	DISTANCE_DATA_LINE_Y  int = SPEED_DATA_LINE_Y + LABEL_GAP_M
	DURATION_DATA_LINE_Y  int = DISTANCE_DATA_LINE_Y + LABEL_GAP_M
	TIME_DIGIT_WIDTH      int = 36
	TIME_COLON_WIDTH      int = 14
	TIME_COLON_OFFSET     int = 4
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
	SECOND_CHANGED int = 0
	MINUTE_CHANGED int = 1
	HOUR_CHANGED   int = 2
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
