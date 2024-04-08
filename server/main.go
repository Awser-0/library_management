package main

import (
	_ "library/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"library/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
