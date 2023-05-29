package git

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func SslVerify(isVerify bool) {
	fmt.Println("git setting sslVerify ", strconv.FormatBool(isVerify))
	cmdItems := strings.Fields(fmt.Sprintf("git config --global http.sslVerify %s", strconv.FormatBool(isVerify)))
	cmd := exec.Command(cmdItems[0], cmdItems[1:]...)
	cmd.Start()
	cmd.Wait()
}
