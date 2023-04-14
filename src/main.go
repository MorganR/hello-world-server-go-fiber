package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/valyala/fasthttp"
)

func main() {
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
	address := ":8080"
	log.Fatal(app.Listen(address))
}
