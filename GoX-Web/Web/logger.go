package Web

import (
	"log"
	"time"
)

func Logger() HandleFunc {
	return func(c *Context) {
		// start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}
