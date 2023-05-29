package source

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/luoruofeng/fx-tool/git"
)

func showRound(ctx context.Context) {
	c := []string{"|", "/", "-", "\\"}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, v := range c {
				time.Sleep(time.Millisecond * 100)
				fmt.Printf("\r%s", v)
			}
		}
	}
}

func print(ctx context.Context, stdout io.ReadCloser) {
	r := bufio.NewReader(stdout)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			bs, _, err := r.ReadLine()
			if err != nil && err != io.EOF {
				panic(err)
			} else if err != nil && err == io.EOF {
				return
			}
			c := string(bs)
			//显示到终端后是否需要被下一行替换
			if strings.HasPrefix(c, "remote: Counting objects:") ||
				strings.HasPrefix(c, "remote: Compressing objects:") ||
				strings.HasPrefix(c, "Receiving objects:") ||
				strings.HasPrefix(c, "Resolving deltas:") {
				fmt.Printf("\033[2K\r%s ", c)
			} else {
				fmt.Println(string(bs))
			}
		}
	}
}
func Download(ctx context.Context, url string, branch string) {
	git.SslVerify(false)
	defer git.SslVerify(true)

	cmdLine := fmt.Sprintf("git clone -b %s %s --progress", branch, url)
	cmdItems := strings.Fields(cmdLine)
	cmd := exec.Command(cmdItems[0], cmdItems[1:]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()
	defer stdout.Close()
	cmd.Start()
	//显示loading符号
	go showRound(ctx)
	//打印标准输出
	go print(ctx, stdout)
	//打印标准错误
	go print(ctx, stderr)
	//等结束
	cmd.Wait()

}

func Install(ctx context.Context, url string, branch string) {
	cmdLine := fmt.Sprintf("go install ./fxdemo")
	cmdItems := strings.Fields(cmdLine)
	cmd := exec.Command(cmdItems[0], cmdItems[1:]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()
	defer stdout.Close()
	cmd.Start()
	//显示loading符号
	go showRound(ctx)
	//打印标准输出
	go print(ctx, stdout)
	//打印标准错误
	go print(ctx, stderr)
	//等结束
	cmd.Wait()
}
