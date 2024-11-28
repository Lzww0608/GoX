package main

import (
	"Web"
	"net/http"
)

func main() {
	r := Web.New()
	r.GET("/", func(c *Web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	r.GET("/hello", func(c *Web.Context) {
		c.String(http.StatusOK, "<h1>Hello %s, you're at $s</h1>", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *Web.Context) {
		c.Json(http.StatusOK, Web.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":8080")
}
