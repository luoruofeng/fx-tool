package util

import (
	"fmt"
	"os"
	"sync"

	"github.com/luoruofeng/fx-tool/git"
	"github.com/luoruofeng/fx-tool/variable"
)

var done chan struct{}
var eixtChan chan struct{}
var NormalFinish bool = false
var wg *sync.WaitGroup

func init() {
	eixtChan = make(chan struct{})
	done = make(chan struct{})
}

// 程序异常退出
func Exit() {
	git.SslVerify(true)
	eixtChan <- struct{}{}
	<-done
	fmt.Println("Done")
	wg.Done()
}

// 程序正常结束
func CloseExit() {
	NormalFinish = true
	close(eixtChan)
	CloseSignalChan()
	fmt.Println("Done!")
}

func DeleteProject() {
	if variable.ComponentName == "" && variable.ProjectName == "" {
		Delete(variable.ProjectName)
	} else {
		Delete(variable.ProjectName + "/" + variable.ComponentName)
	}
}

func ReadyExit(wgp *sync.WaitGroup) {
	wg = wgp
	wg.Add(1)
	<-eixtChan
	if !NormalFinish {
		DeleteProject()
		close(eixtChan)
		os.Exit(0)
	}
	close(done)
}

func Delete(path string) {
	if _, err := os.Stat(path); err != nil || os.IsNotExist(err) {
		return
	}

	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}
	fmt.Println("删除目录", path)
}
