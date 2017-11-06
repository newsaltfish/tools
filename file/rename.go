package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ApkTrim(filename string, newPath string) {
	filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(info.Name(), ".apk") {
			return nil
		}
		ns := strings.Split(info.Name(), "-")
		if len(ns) < 2 {
			return nil
		}
		nf, _ := os.Create(newPath + ns[1] + ".apk")
		defer nf.Close()
		bs, _ := ioutil.ReadFile(path)
		nf.Write(bs)
		return nil
	})
}
