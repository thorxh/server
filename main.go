package main

import (
	_ "gitee.com/thorxh/thor-api/internal/logic"
	_ "gitee.com/thorxh/thor-api/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"gitee.com/thorxh/thor-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
