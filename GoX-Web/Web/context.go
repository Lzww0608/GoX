package Web

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer  http.ResponseWriter
	Request *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	//response info
	StatusCode int
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func newContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		Writer:  writer,
		Request: request,
		Path:    request.URL.Path,
		Method:  request.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(status int) {
	c.StatusCode = status
	c.Writer.WriteHeader(status)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) SetContentType(contentType string) {
	c.Writer.Header().Set("Content-Type", contentType)
}
func (c *Context) SetCookie(name, value string, maxAge int) {
	http.SetCookie(c.Writer, &http.Cookie{})
}
func (c *Context) GetCookie(name string) string {
	cookie, err := c.Request.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func (c *Context) DelCookie(name string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	})
}

func (c *Context) Next() {
	c.StatusCode++
}

func (c *Context) String(status int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(status)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) HTML(status int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(status)
	c.Writer.Write([]byte(html))
}

func (c *Context) HTMLBlob(status int, b []byte) {
	c.SetHeader("Content-Type", "application/octet-stream")
	c.Status(status)
	c.Writer.Write(b)
}

func (c *Context) Json(status int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json; charset=utf-8")
	c.Status(status)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Xml(status int, obj interface{}) {
	c.SetHeader("Content-Type", "application/xml; charset=utf-8")
	c.Status(status)
	encoder := xml.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) XmlBlob(status int, b []byte) {
	c.SetHeader("Content-Type", "application/xml; charset=utf-8")
	c.Status(status)
	c.Writer.Write(b)
}

func (c *Context) File(file string) {
	http.ServeFile(c.Writer, c.Request, file)
}

func (c *Context) Redirect(status int, url string) {
	http.Redirect(c.Writer, c.Request, url, status)
}

func (c *Context) Data(status int, data []byte) {
	c.Status(status)
	c.Writer.Write(data)
}
