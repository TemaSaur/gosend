package main

import (
	"database/sql"
	"net/http"
)

func handler(handleFunc func(*Context), db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		context := &Context{
			w:  w,
			r:  r,
			db: db,
		}

		handleFunc(context)
	}
}
