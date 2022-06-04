package main

import (
	_ "server/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"server/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
