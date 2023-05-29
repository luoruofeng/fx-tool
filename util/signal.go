package util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var sigs chan os.Signal

func ListenSignal() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		fmt.Println("开始监听信号")
		s, ok := <-sigs
		if !ok { // 通道关闭
			return
		}
		fmt.Println("接收到信号", s)
		CloseSignalChan()
		Exit()
	}()
}

func CloseSignalChan() {
	signal.Stop(sigs)
	close(sigs)
	fmt.Println("关闭信号监听")
}
