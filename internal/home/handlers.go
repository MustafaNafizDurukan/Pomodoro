package home

func (c *app) onTaskStarted() {
	c.fireEvent(&Event{
		IsProxyEnabled: true,
	})
}

func (c *app) onTaskCompleted() {
	c.fireEvent(&Event{
		IsProxyEnabled: false,
	})
}
