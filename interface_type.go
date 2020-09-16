package main

import (
	"fmt"
)

// interface {}型はいわゆるany型と呼ばれるものと同じ働きをする
// 関数hogeはあらゆる型の引数を受け取ることができる
func hoge(fuga interface{}) {
	// .(型名)でinterface{}型の変数をキャストする
	// swith文で型を判断することができる
	switch fuga.(type) {
	case int:
		fmt.Println(fuga.(int))
	case string:
		fmt.Println(fuga.(string))
	}
}

func PrintOut(a interface{}) {
	// aをPrintableインタフェースを実装したオブジェクトに変換してみる
	q, ok := a.(Printable)
	if ok {
		// 変換できたらそのインタフェースを呼び出す
		fmt.Println(q.ToString())
	} else {
		fmt.Println("Not printable.")
	}
}

type Printable interface {
	ToString() string
}

func main() {
	a := 100
	hoge(a) // -> 100
	b := "hogefuga"
	hoge(b) // -> hogefuga

	PrintOut(b) // -> Not printable
}
