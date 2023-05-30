package cmd

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/luoruofeng/fx-tool/util"
)

func GetFlag() (string, string, []string) {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	initUrl := initCmd.String("url", "", "输入项目的URL地址,格式为 [repository_host]/[org]/[project_name],例如github.com/luoruofeng/demo")

	// addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	// delCmd := flag.NewFlagSet("del", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("请输入命令,如 init初始化新项目 add添加新组件 del删除组件")
		util.Exit()
	}
	switch os.Args[1] {
	case "init":
		initCmd.Parse(os.Args[2:])
		if *initUrl == "" {
			fmt.Println("请输入项目的url地址,如 -url=\"github.com/luoruofeng/demo\"")
			util.Exit()
		}
		pattern := "^([^/]+)/([^/]+)/([^/]+)$"
		re := regexp.MustCompile(pattern)
		if !re.MatchString(*initUrl) {
			fmt.Println("url地址有误,正确的格式如 -url=\"github.com/luoruofeng/demo\"")
			util.Exit()
		}
		args := []string{}
		if len(initCmd.Args()) > 0 {
			args = initCmd.Args()
		}
		return "init", *initUrl, args

	case "add":
		// addCmd.Parse(os.Args[2:])
		fmt.Println("添加新的组件")
	case "del":
		// delCmd.Parse(os.Args[2:])
		fmt.Println("删除组件")
	}
	fmt.Println("请输入项目的url地址,如 -url=\"github.com/luoruofeng/demo\"")
	util.Exit()
	return "", "", []string{}
}
