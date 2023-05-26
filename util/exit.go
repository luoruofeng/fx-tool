package util

import (
	"os"

	"github.com/luoruofeng/fx-tool/variable"
)

var eixtChan chan struct{}
var NormalFinish bool = false

func init() {
	eixtChan = make(chan struct{})
}

// 程序异常退出
func Exit() {
	eixtChan <- struct{}{}
}

// 程序正常结束
func CloseExit() {
	NormalFinish = true
	close(eixtChan)
}

func ReadyExit() {
	<-eixtChan
	if !NormalFinish {
		defer Delete("./" + variable.ProjectName)
		os.Exit(0)
	}
}

func Delete(path string) {
	if err := os.RemoveAll(path); err != nil {
		panic(err)
	}
}
