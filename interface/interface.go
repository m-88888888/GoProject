package main

import (
	"fmt"
)

// 構造体Person、構造体Bookのどちらも`ToString()`と`PrintOut()`メソッドを実装している
// インタフェースを使ってひとつの関数でまとめる

//type Person struct {
// 	name string
// }
//
// func (p Person) ToString() string {
// 	return p.name
// }
// func (p Person) PrintOut() {
// 	fmt.Println(p.ToString())
// }
//
// type Book struct {
// 	title string
// }
//
// func (b Book) ToString() string {
// 	return b.title
// }
// func (b Book) PrintOut() {
// 	fmt.Println(b.ToString())
// }
//
// func main() {
// 	a1 := Person{name: "山田太郎"}
// 	a2 := Book{title: "吾輩は猫である"}
// 	a1.PrintOut()
// 	a2.PrintOut()
// }

// Printableインタフェースの定義
type Printable interface {
	// ToString()メソッドを定義
	ToString() string
}

// Printableインタフェースを実装した構造体であれば利用可能な関数
func PrintOut(p Printable) {
	// PrintableインタフェースのToString()メソッドを呼び出している
	fmt.Println(p.ToString())
}

type Person struct {
	name string
}

// Printableインタフェースにサポートされている関数を定義したので
// 構造体PersonはPrintableインタフェースが自動的に適用される
func (p Person) ToString() string {
	return p.name
}

// 同様
type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

func main() {
	a1 := Person{name: "山田太郎"}
	a2 := Book{title: "吾輩は猫である"}
	PrintOut(a1)
	PrintOut(a2)
}
