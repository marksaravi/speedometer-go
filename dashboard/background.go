package dashboard

func (d *dashboardDisplay) initBackground() {
	d.display.SetBackgroundColor(dashboard_backgroung_color)
	d.display.Clear()
	d.display.Update()
}
