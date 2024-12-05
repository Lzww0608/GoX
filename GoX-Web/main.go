package main

import (
	"Web"
	"net/http"
)

func main() {
	r := Web.Default()
	r.GET("/", func(c *Web.Context) {
		c.String(http.StatusOK, "Hello Lzww\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *Web.Context) {
		names := []string{"Lzww"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8080")
}
