package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/luoruofeng/fx-tool/cmd"
	"github.com/luoruofeng/fx-tool/source"
	"github.com/luoruofeng/fx-tool/util"
	"github.com/luoruofeng/fx-tool/variable"
)

var wg sync.WaitGroup

func main() {
	subcmd, url, components := cmd.GetFlag()
	switch subcmd {
	case "init":
		variable.NewURL = url
		variable.ProjectName = strings.Split(url, "/")[2]
		util.ListenSignal()
		go util.ReadyExit(&wg)
		util.DeleteProject()
		fmt.Printf("你将创建项目: %s.包含组件: %s.\n", url, strings.Join(components, ","))

		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(time.Minute)*15)
		defer cancelFunc()

		source.Download(ctx, "https://github.com/luoruofeng/fxdemo.git", "basic")
		source.ChangeURL(ctx, url)
		wg.Done()
		util.CloseExit()
		wg.Wait()
	case "add":
		fmt.Println("添加新的组件")
	case "del":
		fmt.Println("删除组件")
	}
}
