package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

// curl http://localhost:3000/fail
// curl http://localhost:3000/success

func main() {
	app := fiber.New()

	app.Use(limiter.New(limiter.Config{
		Max:                    2,
		MaxFunc:                func(_ fiber.Ctx) int { return 1 },
		Expiration:             time.Minute,
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.SlidingWindow{},
	}))

	app.Get("/:status", func(c fiber.Ctx) error {
		if c.Params("status") == "fail" { //nolint:goconst // test
			return c.SendStatus(400)
		}
		return c.SendStatus(200)
	})

	log.Fatal(app.Listen(":3000"))
}
