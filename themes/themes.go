package themes

import "github.com/marksaravi/drivers-go/colors"

type Theme struct {
	BackgroungColor    colors.Color
	SpeedColor         colors.Color
	DistanceColor      colors.Color
	DurationColor      colors.Color
	SpeedLabelColor    colors.Color
	DistanceLabelColor colors.Color
	DurationLabelColor colors.Color
	SeparatorColor     colors.Color
}

// var backgroundColor = colors.DODGERBLUE
var backgroundColor = colors.DEEPSKYBLUE

// var backgroundColor = colors.NAVAJOWHITE
// var backgroundColor = colors.LIGHTSTEELBLUE
// var backgroundColor = colors.LIGHTBLUE

var Default = Theme{
	BackgroungColor:    backgroundColor,
	SpeedColor:         colors.BLACK,
	SpeedLabelColor:    colors.BLACK,
	DistanceColor:      colors.BLACK,
	DistanceLabelColor: colors.BLACK,
	DurationColor:      colors.BLACK,
	DurationLabelColor: colors.BLACK,
	SeparatorColor:     colors.BLACK,
}
