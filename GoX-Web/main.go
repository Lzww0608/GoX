package main

import (
	"Web"
	"net/http"
)

func main() {
	r := Web.New()
	r.GET("/index", func(c *Web.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *Web.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Web</h1>")
		})

		v1.GET("/hello", func(c *Web.Context) {
			// expect /hello?name=Lzww
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *Web.Context) {
			// expect /hello/Lzww
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *Web.Context) {
			c.Json(http.StatusOK, Web.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":8080")
}
