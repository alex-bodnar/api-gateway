package app

// InitWorkers ...
func (a *App) initWorkers() []worker {
	workers := []worker{
		serveHTTP,
		serveBroker,
	}

	return workers
}
