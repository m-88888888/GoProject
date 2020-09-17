package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

//var msg = flag.String("msg", "デフォルト値", "説明")
//var n int

//func init() {
//	flag.IntVar(&n, "n", 1, "回数")
//}

func main() {
	// フラグ（オプション）の設定
	var n = flag.Bool("n", false, "通し番号を付与する")
	flag.Parse()

	var (
		files     = flag.Args()
		path, err = os.Executable()
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
	}

	path = filepath.Dir(path)
	cnt := 1

	for i := 0; i < len(files); i++ {
		sf, err := os.Open(filepath.Join(path, files[i]))
		if err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました", err)
		}

		scanner := bufio.NewScanner(sf)
		for scanner.Scan() {
			if *n {
				fmt.Printf("%v: ", cnt)
				cnt++
			}
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました：", err)
		}
	}
}
