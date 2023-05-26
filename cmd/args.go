package cmd

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/luoruofeng/fx-tool/util"
)

func GetFlag() (string, []string) {
	url := flag.String("url", "", "输入项目的URL地址,格式为 [repository_host]/[org]/[project_name],例如github.com/luoruofeng/demo")
	flag.Parse()
	if *url == "" {
		fmt.Println("请输入项目的url地址,如 -url=\"github.com/luoruofeng/demo\"")
		util.Exit()
	}

	pattern := "^([^/]+)/([^/]+)/([^/]+)$"
	re := regexp.MustCompile(pattern)
	if !re.MatchString(*url) {
		fmt.Println("url地址有误,正确的格式如 -url=\"github.com/luoruofeng/demo\"")
		util.Exit()
	}

	args := []string{}
	if len(flag.Args()) > 0 {
		args = flag.Args()
	}
	return *url, args
}
