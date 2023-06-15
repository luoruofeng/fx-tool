package source

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/luoruofeng/fx-tool/git"
	"github.com/luoruofeng/fx-tool/util"
)

func DeleteModuleFiles(path string) {
	de, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("删除模块失败", err)
		util.Exit()
	}
	for _, v := range de {
		if !v.IsDir() &&
			(v.Name() == "go.mod" ||
				v.Name() == "go.sum" ||
				v.Name() == "go.sum" ||
				v.Name() == ".git" ||
				v.Name() == ".gitignore" ||
				v.Name() == "LICENSE") {
			err := os.Remove(filepath.Join(path, v.Name()))
			if err != nil {
				fmt.Println("删除模块失败", err)
				util.Exit()
			}
			fmt.Println("成功删除文件", filepath.Join(path, v.Name()))
		}
	}
}

func GetRequirement(ctx context.Context, componentDir string) {
	file, err := os.Open(filepath.Join(componentDir, "requirement.txt"))
	if err != nil {
		fmt.Println("读取requirement.txt失败", err)
		util.Exit()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		fmt.Println("开始下载component所需依赖", url)
		cmdLine := fmt.Sprintf("go get -u %s", url)
		fmt.Println(cmdLine)
		cmdItems := strings.Fields(cmdLine)
		cmd := exec.Command(cmdItems[0], cmdItems[1:]...)
		cmd.Dir = componentDir
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
	fmt.Println("成功下载component所需依赖")
}

func showRound(ctx context.Context) {
	c := []string{"|", "/", "-", "\\"}
	for {
		select {
		case <-ctx.Done():
			fmt.Println("执行超时.....")
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
			fmt.Println("执行超时.......")
			return
		default:
			bs, _, err := r.ReadLine()
			if err != nil && err != io.EOF {
				return
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
func Download(ctx context.Context, url string, branch string, dir string) {
	git.SslVerify(false)
	defer git.SslVerify(true)

	cmdLine := fmt.Sprintf("git clone -b %s %s %s --progress", branch, url, dir)
	cmdItems := strings.Fields(cmdLine)
	cmd := exec.Command(cmdItems[0], cmdItems[1:]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	defer stderr.Close()
	defer stdout.Close()
	cmd.Start()
	go func() {
		<-ctx.Done()
		fmt.Println("执行超时..")
		err := cmd.Process.Signal(os.Interrupt)
		if err != nil {
			err := cmd.Process.Kill()
			if err != nil {
				fmt.Println("killing process...", err)
			}
		}
		util.Exit()
	}()
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
