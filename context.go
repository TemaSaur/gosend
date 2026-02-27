package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

type Context struct {
	w  http.ResponseWriter
	r  *http.Request
	db *sql.DB
}

func coloredStatus(status int) string {
	if status >= 400 {
		return fmt.Sprintf("\033[101m\033[30m %v \033[0m", status)
	}
	return fmt.Sprintf("\033[102m\033[30m %v \033[0m", status)
}

func (c *Context) Write(status int, body string) {
	c.w.WriteHeader(status)
	fmt.Fprintf(c.w, body)

	fmt.Printf(
		"[INFO] %v %v | %v | %v %v %v\n",
		time.Now().Format(time.DateOnly),
		time.Now().Format(time.TimeOnly),
		coloredStatus(status),
		c.r.RemoteAddr,
		c.r.Method,
		c.r.URL,
	)
}
