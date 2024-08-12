package handler

import "github.com/gofiber/fiber/v3"

type RedisHandler struct {
}

func (rh *RedisHandler) Health(c fiber.Ctx) error {
	return c.SendString("hello their ! I am healthy")
}
