package main

import (
	"github.com/gogf/gf-demos/v2/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	var (
		ctx = gctx.New()
	)
	cmd.Main.Run(ctx)
}
