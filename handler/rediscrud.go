package handler

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

type RedisHandler struct {
	Rdb *redis.Client
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (rh *RedisHandler) Health(c fiber.Ctx) error {
	return c.SendString("hello their ! I am healthy")
}

func (rh *RedisHandler) Get(c fiber.Ctx) error {
	key := c.Params("key")
	val, err := rh.Rdb.Get(c.Context(), key).Result()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Key not found",
		})
	}
	return c.JSON(KeyValue{Key: key, Value: val})
}

func (rh *RedisHandler) Post(c fiber.Ctx) error {
	var kv KeyValue
	if err := json.Unmarshal(c.Body(), &kv); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := rh.Rdb.Set(c.Context(), kv.Key, kv.Value, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to set key-value pair",
		})
	}
	return c.JSON(kv)
}

func (rh *RedisHandler) Delete(c fiber.Ctx) error {
	key := c.Params("key")
	err := rh.Rdb.Del(c.Context(), key).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete key-value pair",
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
