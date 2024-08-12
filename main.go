package main

import (
	_ "gf-shop/internal/packed"

	_ "gf-shop/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
