package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/luoruofeng/fx-tool/cmd"
	"github.com/luoruofeng/fx-tool/source"
	"github.com/luoruofeng/fx-tool/util"
	"github.com/luoruofeng/fx-tool/variable"
)

var wg sync.WaitGroup

func getComponentDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	componetDir := ""
	if filepath.Base(cwd) == variable.ProjectName {
		// 当前目录是项目根目录
		componetDir = filepath.Join("component", variable.ComponentName)
	} else {
		// 当前目录不是项目根目录
		if fis, err := ioutil.ReadDir(cwd); err != nil {
			return "", err
		} else {
			for _, fi := range fis {
				if variable.ProjectName == fi.Name() {
					// 当前目录是项目根目录的父目录
					componetDir = filepath.Join(variable.ProjectName, "component", variable.ComponentName)
					break
				}
			}
			if componetDir == "" {
				panic("当前目录有误")
			}
		}
	}
	return componetDir, nil
}

func addComponent(name string) error {
	fmt.Println("开始下载组件", name)

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(time.Minute)*15)
	defer cancelFunc()

	// 下载模版项目
	componetDir, err := getComponentDir()
	if err != nil {
		return err
	}
	source.Download(ctx, "https://github.com/luoruofeng/fx-component.git", name, componetDir)
	// 替换项目文件内容
	replacePairs := source.NewReplacePairs(
		source.ReplacePair{
			Old: "github.com/luoruofeng/fx-component",
			New: variable.NewURL + "/component",
		},
	)
	source.ReplaceTemplateContent(ctx, componetDir, replacePairs)
	return nil
}

func main() {
	subcmd, url, components := cmd.GetFlag()
	switch subcmd {
	case "initial":
		// 设置项目URL和项目名称
		variable.NewURL = url
		variable.ProjectName = strings.Split(url, "/")[2]
		fmt.Println("global url", variable.NewURL)
		fmt.Println("global project name", variable.ProjectName)
		util.ListenSignal()
		go util.ReadyExit(&wg)
		util.DeleteProject()
		fmt.Printf("你将创建项目: %s.包含组件: %s.\n", url, strings.Join(components, ","))

		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(time.Minute)*25)
		defer cancelFunc()

		// 下载模版项目
		source.Download(ctx, "https://github.com/luoruofeng/fxdemo.git", "basic", "./"+variable.ProjectName)
		// 修改项目文件夹名称
		// source.ChangeDirName(ctx, url, "./fxdemo", variable.ProjectName)
		// 替换项目文件内容
		replacePairs := source.NewReplacePairs(
			source.ReplacePair{
				Old: "github.com/luoruofeng/fxdemo",
				New: variable.NewURL,
			},
			source.ReplacePair{
				Old: "fxdemo",
				New: variable.ProjectName,
			},
		)
		source.ReplaceTemplateContent(ctx, "./"+variable.ProjectName, replacePairs)
		wg.Done()
		wg.Wait()
		util.CloseExit()
		if len(components) > 0 {
			for _, name := range components {
				// 设置组件名称和组件版本
				variable.ComponentName = strings.Split(name, "-")[0]
				variable.ComponentVersion = strings.Split(name, "-")[1]
				fmt.Println("global compoent name", variable.ComponentName)
				fmt.Println("global compoent version", variable.ComponentVersion)

				if err := addComponent(name); err != nil {
					fmt.Println("添加组件出错", err)
				}
			}
		}
	case "add":
		fmt.Println("添加新的组件", components)

		util.ListenSignal()
		go util.ReadyExit(&wg)

		// 设置项目URL和项目名称
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		gomod, err := os.Open(filepath.Join(cwd, "go.mod"))
		if err != nil {
			fmt.Println("当前文件夹下没有go.mod文件", err)
			return
		}
		defer gomod.Close()
		scanner := bufio.NewScanner(gomod)
		if scanner.Scan() {
			oneLine := scanner.Text()
			variable.NewURL = strings.Split(oneLine, " ")[1]
			variable.ProjectName = strings.Split(variable.NewURL, "/")[len(strings.Split(variable.NewURL, "/"))-1]
		} else {
			fmt.Println("当前文件夹下go.mod文件第一行内容没有URL信息", err)
			return
		}
		fmt.Println("global url", variable.NewURL)
		fmt.Println("global project name", variable.ProjectName)

		for _, name := range components {
			// 设置组件名称和组件版本
			variable.ComponentName = strings.Split(name, "-")[0]
			variable.ComponentVersion = strings.Split(name, "-")[1]
			fmt.Println("global compoent name", variable.ComponentName)
			fmt.Println("global compoent version", variable.ComponentVersion)
			// 给当前项目新增组件
			if err := addComponent(name); err != nil {
				fmt.Println("添加组件出错", err)
				continue
			}
		}
		wg.Done()
		wg.Wait()
		fmt.Println("Add components success")
	case "del":
		fmt.Println("删除组件")
	}
}
