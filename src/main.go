package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/valyala/fasthttp"
)

func main() {
	port := "8080"
	if portEnv := os.Getenv("PORT"); portEnv != "" {
		port = portEnv
	}

	app := fiber.New()

	app.Use("/strings", compress.New())
	app.Get("/strings/hello", HelloWorldHandler)
	app.Get("/strings/async-hello", AsyncHelloWorldHandler)
	app.Get("/strings/lines", LinesHandler)

	app.Get("/math/power-reciprocals-alt", PowerReciprocalsAltHandler)

	fs := &fasthttp.FS{
		Root:               "./",
		GenerateIndexPages: false,
		Compress:           true,
		CompressBrotli:     true,
		CompressedFileSuffixes: map[string]string{
			"br": ".br",
		},
	}
	fsh := fs.NewRequestHandler()
	app.Get("/static/*", func(ctx *fiber.Ctx) error {
		fsh(ctx.Context())
		return nil
	})
	log.Fatal(app.Listen(":" + port))
}
