package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/valyala/fasthttp"
)

const maxNameLength = 500

var nameParam = "name"
var nameTooLongErrMsg = fmt.Sprintf("Name must be <= %v characters", maxNameLength)
var defaultGreeting = "Hello, world!"
var nParam = "n"
var nMustBeIntErrMsg = "query param n must be an integer"

// AsyncHelloWorldHandler provides a greeting after a delay.
func AsyncHelloWorldHandler(ctx *fiber.Ctx) error {
	time.Sleep(time.Millisecond * 15)
	ctx.SendString(defaultGreeting)
	return nil
}

// HelloWorldHandler provides a greeting, using the optional "name" query parameter.
func HelloWorldHandler(ctx *fiber.Ctx) error {
	name := ctx.Query(nameParam, "")
	if len(name) > maxNameLength {
		return fiber.NewError(fiber.StatusBadRequest, nameTooLongErrMsg)
	}

	ctx.Set(fasthttp.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	if name == "" {
		ctx.SendString(defaultGreeting)
		return nil
	}
	ctx.SendString("Hello, " + name + "!")
	return nil
}

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

func main() {
	port := "8080"
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = portEnv
	}

	app := fiber.New()

	// Apply compression middleware to all paths under /strings/.
	app.Use("/strings/", compress.New())
	app.Get("/strings/hello", HelloWorldHandler)
	app.Get("/strings/async-hello", AsyncHelloWorldHandler)
	app.Get("/strings/lines", LinesHandler)

	app.Get("/math/power-reciprocals-alt", PowerReciprocalsAltHandler)

	// Configure static file serving.
	fs := &fasthttp.FS{
		Root:               "./",
		GenerateIndexPages: false,
		Compress:           true,
		CompressBrotli:     true,
		CompressedFileSuffixes: map[string]string{
			"gzip": ".gz",
			"br":   ".br",
		},
	}

	fsh := fs.NewRequestHandler()
	app.Get("/static/*", func(ctx *fiber.Ctx) error {
		fsh(ctx.Context())
		return nil
	})

	// Serve.
	log.Fatal(app.Listen(":" + port))
}
