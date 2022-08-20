package app

// InitWorkers ...
func (a *App) initWorkers() []worker {
	workers := []worker{
		serveHTTP,
	}

	return workers
}
