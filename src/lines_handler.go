package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var nParam = "n"
var nMustBeIntErrMsg = "query param n must be an integer"

// LinesHandler outputs multiple lines of text based on the optional query parameter "n".
func LinesHandler(ctx *fiber.Ctx) error {
	nStr := ctx.Query(nParam, "")
	n := 0
	if nStr != "" {
		var err error
		n, err = strconv.Atoi(nStr)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, nMustBeIntErrMsg)
		}
	}

	ctx.Set(fasthttp.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	sb := strings.Builder{}
	sb.WriteString("<ol>\n")
	for i := 1; i <= n; i++ {
		sb.WriteString(fmt.Sprintf("  <li>Item number: %v</li>\n", i))
	}
	sb.WriteString("</ol>")
	ctx.SendString(sb.String())
	return nil
}
