package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	SuggestionURL = "http://sug.music.baidu.com/info/suggestion"
	Fmlink        = "http://music.baidu.com/data/music/fmlink"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Input URL")
		return
	}
	fmt.Println("fetching msg from ", os.Args[1])

	nurl := strings.Replace(os.Args[1], "#/", "", -1)
	response, err := DownLoadString(nurl, nil)
	if err != nil {
		fmt.Println(" wrong in urling", err)
		return
	}
	var path string
	if os.IsPathSeparator("\\") {
		path = "\\"
	} else {
		path = "/"
	}
	dir, _ := os.Getwd()
	dir = dir + path + "songs_dir"

	if _, err := os.Stat(dir); err != nil {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println(" create dir wrong", err)
			return
		}
	}
}
