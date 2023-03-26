package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// AsyncHelloWorldHandler provides a greeting after a delay.
func AsyncHelloWorldHandler(ctx *fiber.Ctx) error {
	time.Sleep(time.Millisecond * 15)
	ctx.SendString("Hello, world!")
	return nil
}
