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

	app.Use("/hello", compress.New())
	app.Get("/hello", HelloWorldHandler)

	fs := &fasthttp.FS{
		Root:               "./",
		GenerateIndexPages: false,
		Compress:           true,
		CompressBrotli:     true,
		CompressedFileSuffixes: map[string]string{
			"br":   ".br",
			"gzip": ".gz",
		},
	}
	fsh := fs.NewRequestHandler()
	app.Get("/static/*", func(ctx *fiber.Ctx) error {
		fsh(ctx.Context())
		return nil
	})
	log.Fatal(app.Listen(":" + port))
}
