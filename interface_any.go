package main

import (
	"fmt"
)

type any interface{}
type dict map[string]any

func main() {
	// interface{}をanyのように使って任意の型を持つMapを定義
	p1 := map[string]interface{}{
		"name": "yamada",
		"age":  26,
	}

	// pythonのdictのような階層構造もどきを定義できる
	p2 := dict{
		"name": "Tanaka",
		"age":  "30",
		"address": dict{
			"zip": "123-4567",
			"tel": "012-3456-7890",
		},
	}

	name := p2["name"]
	tel := p2["address"].(dict)["tel"] // anyをdictにキャストしてから参照

	fmt.Println(name)
	fmt.Println(tel)
}
