package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

// PowerReciprocalsAltHandler computes a converging series n times, using query parameter "n".
func PowerReciprocalsAltHandler(ctx *fiber.Ctx) error {
	nStr := ctx.Query(nParam, "")
	n := 0
	if nStr != "" {
		var err error
		n, err = strconv.Atoi(nStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, nMustBeIntErrMsg)
		}
	}

	result := 0.0
	power := 0.5
	for ; n > 0; n-- {
		power *= 2
		result += 1 / power

		if n > 1 {
			n--
			power *= 2
			result -= 1 / power
		}
	}

	ctx.Set(fasthttp.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	ctx.SendString(strconv.FormatFloat(result, 'f', -1, 64))
	return nil
}
