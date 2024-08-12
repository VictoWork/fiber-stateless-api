package application

import "github.com/victowork/fiber-stateless-api/handler"

func (a *App) LoadRoutes() {

	handler := &handler.RedisHandler{
		Rdb: a.rdb,
	}
	a.server.Get("/health", handler.Health)
	a.server.Get("/key/:key", handler.Get)
	a.server.Post("/key", handler.Post)
	a.server.Delete("/key/:key", handler.Delete)
}
