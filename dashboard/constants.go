package dashboard

import (
	"time"

	"github.com/marksaravi/devices-go/colors"
	"github.com/marksaravi/fonts-go/fonts"
)

const (
	BACKGROUNG_COLOR     = colors.WHITE
	LABEL_COLOR          = colors.BLACK
	SPEED_LABEL_COLOR    = LABEL_COLOR
	DISTANCE_LABEL_COLOR = LABEL_COLOR
	DURATION_LABEL_COLOR = LABEL_COLOR
	DATA_COLOR           = colors.RED
	SPEED_DATA_COLOR     = DATA_COLOR
	DISTANCE_DATA_COLOR  = DATA_COLOR
	DURATION_DATA_COLOR  = DATA_COLOR
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
)
