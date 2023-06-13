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

type ReplacePair struct {
	Old string
	New string
}

func NewReplacePairs(pairs ...ReplacePair) []ReplacePair {
	var result []ReplacePair = make([]ReplacePair, 0, len(pairs))
	result = append(result, pairs...)
	return result
}

func replaceContent(path string, replacePair []ReplacePair) error {
	// 读取文件内容
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// 替换字符串
	newContent := string(content)
	for _, v := range replacePair {
		newContent = strings.ReplaceAll(newContent, v.Old, v.New)
	}

	// 写入文件
	err = ioutil.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func deletedir(ctx context.Context, path string) {
	if e, err := os.ReadDir(path); err != nil {
		fmt.Println("删除目录失败", path, err)
		util.Exit()
	} else {
		for _, v := range e {
			if v.Name() == ".git" ||
				v.Name() == ".gitignore" ||
				v.Name() == "LICENSE" {
				fmt.Println("删除不需要的文件", filepath.Join(path, v.Name()))
				if err := os.RemoveAll(filepath.Join(path, v.Name())); err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}
}

func checkDir(ctx context.Context, path string, replacePair []ReplacePair) {
	fmt.Println("开始检查目录", path)

	isTimeout := false

	go func() {
		<-ctx.Done()
		fmt.Println("执行超时....")
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

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(filename, ".go") ||
			strings.HasSuffix(filename, ".mod") ||
			strings.HasSuffix(filename, "Makefile") {
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

func ReplaceTemplateContent(ctx context.Context, newDirName string, replacePair []ReplacePair) {
	path, _ := filepath.Abs(newDirName)
	fi, err := os.Stat(newDirName)
	if err != nil {
		fmt.Println("github连接失败。", newDirName, err)
		util.Exit()
	}
	if fi.IsDir() {
		deletedir(ctx, path)
		checkDir(ctx, path, replacePair)
	} else {
		fmt.Println("请确组件已经从github成功下载。", newDirName, err)
		util.Exit()
	}
}
