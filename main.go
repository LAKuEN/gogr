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

	// TODO コア数の取得: GOMAXPROCSが自動的にCPU数分の値になるので、それを取ればいい
	// fmt.Println("NumCPU:", runtime.NumCPU())
	// fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(-1))
	// TODO コア数で処理対象のデータを分割して並行処理

	// TODO 結果の並び順はファイル内での順序を保ったものにする
	//      前から順番に一定件数を取り出して処理させ、それらを前方の結果から順に連結していけば良い
	// FIXME 文字列そのものを配列に格納するのではなく、ポインタのみを格納したい
	//       コピーを作るのにコストが掛かりそう…
	matchedLines := extractLines(&lines, query, 0, len(lines))

	for _, l := range matchedLines {
		fmt.Println(l)
	}
}

func onError(err error) {
	fmt.Fprintf(os.Stderr, err.Error())
	os.Exit(1)
}

func extractLines(lines *[]string, query string, startIdx, endIdx int) []string {
	var matchedLines []string
	for _, line := range *lines {
		// TODO 正規表現での検索
		if strings.Contains(line, query) {
			matchedLines = append(matchedLines, line)
		}
	}

	return matchedLines
}
