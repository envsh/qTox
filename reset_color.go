package main

import (
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	filepath.Walk("uiorig", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".css") {
			log.Println("===========processing...", path)
			reset_color(path)
		}
		return nil
	})
}

func reset_color(fname string) {
	// fname := "uiorig/window/statusPanel.css"
	bcc, err := ioutil.ReadFile(fname)
	gopp.ErrPrint(err)
	lines := strings.Split(string(bcc), "\n")
	log.Println("lines:", len(lines))
	newlines := []string{}
	for _, line := range lines {
		tline := strings.TrimSpace(line)
		if strings.HasPrefix(tline, "color:") || strings.HasPrefix(tline, "background-color:") ||
			strings.HasPrefix(tline, "background:") || strings.HasPrefix(tline, "border-color:") {
			// log.Println("drop line:", line)
		} else {
			newlines = append(newlines, line)
		}
	}
	newbcc := []byte(strings.Join(newlines, "\n"))
	newpath := strings.Replace(fname, "uiorig/", "ui/", -1)
	err = ioutil.WriteFile(newpath, newbcc, 0755)
	gopp.ErrPrint(err, newpath)
	if err == nil {
		log.Println("rewrite css done:", newpath)
	}
}
