package dashboard

import "github.com/marksaravi/devices-go/colors/rgb565"

const (
	max_speed                 float64 = 60  // km/h
	speed_resolution          float64 = 2   // km/h
	speed_curve_start_angle   float64 = 90  // degree
	speed_curve_end_angle     float64 = 300 // degree
	speed_curve_center_x      float64 = 120
	speed_curve_center_y      float64 = 120
	speed_curve_center_radius float64 = 110
	speed_curve_center_width  float64 = 20
)

func (d *dashboardDisplay) drwaSpeedsBackground() {
	d.display.SetColor(rgb565.WHITE)
	d.display.Arc(speed_curve_center_x, speed_curve_center_y, speed_curve_center_radius, toDeg(speed_curve_start_angle), toDeg(speed_curve_end_angle), speed_curve_center_width)
}
