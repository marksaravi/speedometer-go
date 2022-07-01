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
	DATA_COLOR           = colors.BEIGE
	SPEED_DATA_COLOR     = DATA_COLOR
	DISTANCE_DATA_COLOR  = DATA_COLOR
	DURATION_DATA_COLOR  = DATA_COLOR
)

// basic dimensions
const (
	TOP_MARGIN           int = 32
	LEFT_MARGIN          int = 16
	SPEED_LINE_HEIGHT    int = 32
	DISTANCE_LINE_HEIGHT int = 48
	DURATION_LINE_HEIGHT int = 48
	LABEL_COLUMN         int = 0
	DATA_COLUMN          int = 200
)

// Fonts
var (
	SPEED_FONT    fonts.BitmapFont = fonts.FreeMono24pt7b
	DISTANCE_FONT fonts.BitmapFont = fonts.FreeMono12pt7b
	DURATION_FONT fonts.BitmapFont = fonts.FreeMono12pt7b
)

const (
	SPEED_RESOLUTION    float64       = 0.5
	DISTANCE_RESOLUTION float64       = 10
	DURATION_RESOLUTION time.Duration = time.Second
)
