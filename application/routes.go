package application

import "github.com/victowork/fiber-stateless-api/handler"

func (a *App) LoadRoutes() {

	handler := &handler.RedisHandler{
		Rdb: a.rdb,
	}
	a.server.Get("/health", handler.Health)
	a.server.Get("/key/:key", handler.Get)
	a.server.Get("/key", handler.Post)
	a.server.Get("/key/:key", handler.Delete)
}
