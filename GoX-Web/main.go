package main

import (
	"Web"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := Web.New()
	r.Use(Web.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Lzww", Age: 24}
	stu2 := &student{Name: "Carl", Age: 12}
	r.GET("/", func(c *Web.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *Web.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", Web.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *Web.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", Web.H{
			"title": "GoX-Web",
			"now":   time.Date(2024, 12, 4, 20, 31, 0, 0, time.UTC),
		})
	})

	r.Run(":8080")
}
