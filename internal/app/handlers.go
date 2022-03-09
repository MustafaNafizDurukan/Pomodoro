package app

// onTaskStarted fires event when task started with specific Event struct.
func (a *App) onTaskStarted() {
	a.fireEvent(&Event{
		IsProxyEnabled: true,
	})
}

// onTaskCompleted fires event when task completed with specific Event struct.
func (a *App) onTaskCompleted() {
	a.fireEvent(&Event{
		IsProxyEnabled: false,
	})
}
