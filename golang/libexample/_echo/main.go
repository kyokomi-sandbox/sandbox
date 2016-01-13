package main

import (
	"log"

	"net/http"

	"strings"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(hoge1())
	e.Use(hoge2())

	// Routes
	e.Get("/hoge/fuga", hello)
	e.Get("/hoge/piyo", hello)
	e.Get("/xxxx/piyo", hello)

	// Start server
	e.Run(":1323")
}

func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func hoge1() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			c.Set("hoge", "fuga")

			path := c.Request().URL.Path
			switch {
			case strings.HasPrefix(path, "/hoge"):
				log.Println("hoge1", path)
			}
			return h(c)
		}
	}
}

func hoge2() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			log.Println(c.Get("hoge"))
			return h(c)
		}
	}
}
