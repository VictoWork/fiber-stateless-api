package application

import "github.com/victowork/fiber-stateless-api/handler"

func (a *App) LoadRoutes() {

	handler := &handler.RedisHandler{}
	a.server.Get("/health", handler.Health)
}
