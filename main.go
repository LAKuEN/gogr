package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		// fmt.Fprintf(os.Stderr, "need to pass the filepath")
		// os.Exit(1)
		onError(fmt.Errorf("need to pass the filepath and search string"))
	}
	query, filePath := args[1], args[2]

	// TODO リダイレクトされた文字列に対しても処理ができるようにしたい
	//      なので、冗長だがos.Open()とioutil.ReadAll()で実装しておく
	file, err := os.Open(filePath)
	if err != nil {
		onError(err)
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		onError(err)
	}
	lines := strings.Split(string(bytes), "\n")

	// TODO 指定された文字列に合致する文字列が含まれる行を抽出
	// var res []string
	for _, line := range lines {
		if strings.Contains(line, query) {
			// res = append(res, line)
			fmt.Println(line)
		}
	}
}

func onError(err error) {
	fmt.Fprintf(os.Stderr, err.Error())
	os.Exit(1)
}
