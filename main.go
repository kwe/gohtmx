package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views/templates", ".html")
	config := fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		Prefork:     false,
	}

	app := fiber.New(config)
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(logger.New())

	app.Static("/static", "static")
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/messages", func(c *fiber.Ctx) error {
		// Render message template
		return c.Render("messages", nil, "layouts/empty")
	})

	app.Get("/metrics", monitor.New())

	log.Fatal(app.Listen(":3000"))
}
