package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func handleAll(c *Context) {
	c.Write(http.StatusBadRequest, "Bad Request")
}

func getIcon(c *Context) {
	c.Write(http.StatusNoContent, "")
}

func postForm(c *Context) {
	err := c.r.ParseForm()
	if err != nil {
		c.Write(http.StatusBadRequest, "Bad Form")
		return
	}
	source := c.r.PathValue("source")

	var buf bytes.Buffer

	for k, v := range c.r.Form {
		fmt.Fprintf(&buf, "[%v]\n%v\n", k, v[0])
	}

	content := buf.String()

	_, err = c.db.Exec(createSql, source, content)
	if err != nil {
		c.Write(http.StatusInternalServerError, err.Error())
		return
	}

	c.Write(http.StatusOK, "Success")
}

func getAll(c *Context) {
	var buf bytes.Buffer

	rows, err := c.db.Query(getSql)
	if err != nil {
		c.Write(http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	first := true
	for rows.Next() {
		if !first {
			fmt.Fprintf(&buf, "\n%v\n\n", strings.Repeat("-", 80))
		}

		first = false
		var source string
		var content string
		if err := rows.Scan(&source, &content); err != nil {
			c.Write(http.StatusInternalServerError, "Internal server error")
			log.Printf("scan: %v\n", err)
			return
		}

		fmt.Fprintf(&buf, "SOURCE: %v\n", source)
		fmt.Fprintf(&buf, "CONTENT:\n")
		fmt.Fprint(&buf, strings.Join(wrap(content), "\n"))
	}
	c.Write(http.StatusOK, buf.String())
}

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(setupSql)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/favicon.ico", handler(getIcon, nil))
	http.HandleFunc("/", handler(handleAll, nil))
	http.HandleFunc("/get", handler(getAll, db))
	http.HandleFunc("/form/{source}", handler(postForm, db))

	port := ":8000"
	fmt.Printf("Listening on %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
