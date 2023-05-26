package source

import "os"

func RemoveFolder(path string) {
	os.RemoveAll(path)
}
