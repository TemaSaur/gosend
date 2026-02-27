package main

import _ "embed"

//go:embed sql/setup.sql
var setupSql string

//go:embed sql/get.sql
var getSql string

//go:embed sql/create.sql
var createSql string
