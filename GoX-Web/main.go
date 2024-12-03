package main

import (
	"Web"
	"log"
	"net/http"
	"time"
)

func onlyForV2() Web.HandleFunc {
	return func(c *Web.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	r := Web.New()
	r.Use(Web.Logger()) // global midlleware
	r.GET("/", func(c *Web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Web</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *Web.Context) {
			// expect /hello/Lzww
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":8080")
}
