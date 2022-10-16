package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const maxNameLength = 500

var nameParam = "name"
var nameTooLongErrMsg = fmt.Sprintf("Name must be <= %v characters", maxNameLength)

// HelloWorldHandler provides a greeting, using the optional "name" query parameter.
func HelloWorldHandler(ctx *fiber.Ctx) error {
	name := ctx.Query(nameParam, "")
	if name == "" {
		name = "world"
	} else if len(name) > maxNameLength {
		return fiber.NewError(fiber.StatusBadRequest, nameTooLongErrMsg)
	}

	ctx.Set(fasthttp.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	ctx.SendString("Hello, " + name + "!")
	return nil
}
