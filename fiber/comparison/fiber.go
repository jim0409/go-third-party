package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

/*
➜  go-third-party git:(main) ✗ wrk -t100 -c100 http://127.0.0.1:3000
Running 10s test @ http://127.0.0.1:3000
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.13ms    8.30ms 150.67ms   97.37%
    Req/Sec     1.05k   356.24     5.42k    79.32%
  1041675 requests in 10.08s, 134.11MB read
Requests/sec: 103338.77
Transfer/sec:     13.30MB

➜  go-third-party git:(main) ✗ wrk -t100 -c100 http://127.0.0.1:3000/123
Running 10s test @ http://127.0.0.1:3000/123
  100 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.88ms    0.96ms  54.48ms   99.04%
    Req/Sec     1.19k   135.15     2.25k    92.11%
  1192272 requests in 10.10s, 139.86MB read
Requests/sec: 118050.15
Transfer/sec:     13.85MB
*/

func fiberServer() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/:id", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("id"))
		return c.SendString(msg) // => ✋ register
	})

	go app.Listen(":3000")
}
