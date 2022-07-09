package dashboard

import "github.com/marksaravi/devices-go/colors"

type DisplayData struct {
	Speed           float64
	SpeedChanged    bool
	Distance        float64
	DistanceChanged bool
	Sec             int
	SecChanged      bool
	Min             int
	MinChanged      bool
	Hour            int
	HourChanged     bool
}

type Theme struct {
	BackgroungColor    colors.Color
	SpeedLabelColor    colors.Color
	DistanceLabelColor colors.Color
	DurationLabelColor colors.Color
	SpeedDataColor     colors.Color
	DistanceDataColor  colors.Color
	DurationDataColor  colors.Color
}
