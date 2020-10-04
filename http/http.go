package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	//"html"
	//"log"
	"net/http"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//	http.Handle("/foo", fooHandler)
	//
	//	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//	})

	// /hogeに対する処理
	http.HandleFunc("/hoge", hogeHandler)
	http.HandleFunc("/fuga", fugaHandler)
	http.HandleFunc("/testJson", jsonHandler)

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

// /hogeに対する処理内容
func hogeHandler(w http.ResponseWriter, r *http.Request) {

	// HTTPメソッドをチェック
	// RequestのMethodプロパティをチェックすれば判別できる
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("POST ONLY"))
		return
	}
	fmt.Fprint(w, "Hello World from Go.")

}

func fugaHandler(w http.ResponseWriter, r *http.Request) {

	// クエリパラメータ取得してみる
	fmt.Fprintf(w, "クエリ：%s\n", r.URL.RawQuery)

	// Bodyデータを扱う場合には、事前にパースを行う
	r.ParseForm()

	// Formデータを取得.
	form := r.PostForm
	fmt.Fprintf(w, "フォーム：\n%v\n", form)

	// または、クエリパラメータも含めて全部.
	params := r.Form
	fmt.Fprintf(w, "フォーム0：\n%v\n", params)
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// request bodyの読み取り
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("io error")
		return
	}

	// jsonのdecode
	jsonBytes := ([]byte)(b)
	data := new(InputJsonSchema)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error: ", err)
		return
	}
	fmt.Fprintln(w, "JSONでーた: "+data.Name)

	// DB処理
	insert(data)
}

type InputJsonSchema struct {
	Name string `json:"name"`
}

func insert(jsonSchema *InputJsonSchema) {
	// デーベースのオープン
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/hello")
	// エラーハンドリング
	if err != nil {
		log.Fatal(err)
		return
	}
	// クローズ処理の遅延実行
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(id, name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	res, err := stmt.Exec(4, jsonSchema.Name)
	if err != nil {
		log.Fatal(err)
		return
	}
}
