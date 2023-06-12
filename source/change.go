package source

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/luoruofeng/fx-tool/util"
)

func replaceContent(path string, replacePair map[string]string) error {
	// 读取文件内容
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// 替换字符串
	newContent := string(content)
	for k, v := range replacePair {
		newContent = strings.ReplaceAll(newContent, k, v)
	}

	// 写入文件
	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func checkDir(ctx context.Context, path string, replacePair map[string]string) {
	fmt.Println("开始检查目录", path)

	isTimeout := false

	go func() {
		<-ctx.Done()
		fmt.Println("执行超时...")
		isTimeout = true
	}()

	filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if isTimeout {
			fmt.Println("检查目录超时,请检查网络连接")
			util.Exit()
		}
		filename := info.Name()
		if err != nil {
			fmt.Println("文件路径不存在", path, err)
			util.Exit()
		}
		if filename == ".git" ||
			filename == ".gitignore" ||
			filename == "LICENSE" {
			fmt.Println("删除不需要的文件", filename)
			RemoveFolder(filename)
		}
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(filename, ".go") ||
			strings.HasSuffix(filename, ".mod") ||
			strings.HasSuffix(filename, "Makefile") ||
			strings.HasSuffix(filename, "LICENSE") {
			fmt.Println("replacing", path)
			if err := replaceContent(path, replacePair); err != nil {
				fmt.Println("文件内容替换失败", filename, err)
				util.Exit()
			}
		}
		return nil
	})
}

func ChangeDirName(ctx context.Context, new string, targetName string, newName string) {
	err := os.Chmod(targetName, 0666)
	if err != nil {
		fmt.Println("项目权限修改失败", targetName, err)
		util.Exit()
	}
	err = os.Rename(targetName, newName)
	if err != nil {
		fmt.Println("项目重命名失败", targetName, err)
		util.Exit()
	}
}

func ReplaceTemplateContent(ctx context.Context, newDirName string, replacePair map[string]string) {
	path, _ := filepath.Abs(newDirName)
	if fi, err := os.Stat(newDirName); err == nil && fi.IsDir() {
		checkDir(ctx, path, replacePair)
	} else {
		fmt.Println("无法从github.com获取模版项目,请检查git设置和网络连接", newDirName, err)
		util.Exit()
	}
}
