package main

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/luoruofeng/fx-tool/cmd"
	"github.com/luoruofeng/fx-tool/source"
	"github.com/luoruofeng/fx-tool/util"
	"github.com/luoruofeng/fx-tool/variable"
)

var wg sync.WaitGroup

func main() {
	url, components := cmd.GetFlag()
	variable.NewURL = url
	variable.ProjectName = strings.Split(url, "/")[2]
	util.ListenSignal()
	go util.ReadyExit(&wg)
	util.DeleteProject()
	fmt.Printf("你将创建项目: %s.包含组件: %s.\n", url, strings.Join(components, ","))

	ctx := context.Background()
	source.Download(ctx, "https://github.com/luoruofeng/fxdemo.git", "basic")
	source.ChangeURL(ctx, url)
	wg.Done()
	util.CloseExit()
	wg.Wait()
}
