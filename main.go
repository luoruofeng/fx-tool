package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/luoruofeng/fx-tool/cmd"
	"github.com/luoruofeng/fx-tool/source"
	"github.com/luoruofeng/fx-tool/util"
	"github.com/luoruofeng/fx-tool/variable"
)

func main() {
	url, components := cmd.GetFlag()
	fmt.Printf("你将创建项目: %s.包含组件: %s.\n", url, strings.Join(components, ","))

	ctx := context.Background()
	source.Download(ctx, "https://github.com/luoruofeng/fxdemo.git", "basic")
	go util.ReadyExit()
	variable.NewURL = url
	variable.ProjectName = strings.Split(url, "/")[2]
	source.ChangeURL(ctx, url)
	util.CloseExit()
	time.Sleep(1 * time.Second)

}
