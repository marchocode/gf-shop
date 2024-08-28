package main

import (
	_ "gf-shop/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "gf-shop/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
