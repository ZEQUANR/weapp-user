package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/server"
	"github.com/silverswords/sand/server/web"
)

func main() {
	sand.Instance().Load(
		func() server.Server {
			return web.CreateBuilder("./config/server").Build(
			// web.WithHost("host"),
			// web.WithAddr("0"),
			).Run()
		}(),
	).Run()
}
