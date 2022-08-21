package main

import (
	"data-engine/cmd"
	_ "github.com/go-sql-driver/mysql"
)

// @title data-engine API
// @version 1.0.0
// @description data-engine api文档
func main() {
	cmd.Execute()
}
