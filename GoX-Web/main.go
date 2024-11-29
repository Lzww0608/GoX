package main

import (
	"Web"
	"net/http"
)

func main() {
	r := Web.New()
	r.GET("/", func(c *Web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello GoX</h1>")
	})

	r.GET("/hello", func(c *Web.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *Web.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *Web.Context) {
		c.Json(http.StatusOK, Web.H{"filepath": c.Param("filepath")})
	})

	r.Run(":8080")
}
