package main

import (
	mysql "github.com/silverswords/sand/sql/mysql"
)

func main() {
	db := mysql.CreateBuilder("path").Build().Run()
	defer db.Close()
}
