package main

import (
	_ "server_gf/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"server_gf/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
