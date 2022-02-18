package app

// onTaskStarted fires event when task started with specific Event struct.
func (c *App) onTaskStarted() {
	c.fireEvent(&Event{
		IsProxyEnabled: true,
	})
}

// onTaskCompleted fires event when task completed with specific Event struct.
func (c *App) onTaskCompleted() {
	c.fireEvent(&Event{
		IsProxyEnabled: false,
	})
}
