package main

import (
	"fmt"
	"log"
	//"reflect"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// デーベースのオープン
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/hello")
	// エラーハンドリング
	if err != nil {
		log.Fatal(err)
	}
	// クローズ処理の遅延実行
	defer db.Close()

	// Ping()でデータベースにアクセス可能かどうかチェック
	ping := db.Ping()
	if ping != nil {
		log.Println("database connection is failed.")
	} else {
		log.Println("Database connection is successful.")
	}

	// 行をマップで取得することはできない
	var (
		id   int
		name string
	)

	// 行を返却するSQLではQuery()、返却しないSQLではExec()を使用する
	//rows, err := db.Query("SELECT * FROM user WHERE id = ?", 1)
	stmt, err := db.Prepare("SELECT id, name FROM user WHERE id = ?")
	// エラーハンドリング
	if err != nil {
		log.Fatal("prepare statement error.")
		log.Fatal(fmt.Sprint(err))
	}
	defer stmt.Close()

	rows, err := stmt.Query(1)
	if err != nil {
		log.Fatal("query exeute error")
		log.Fatal(err)
	}
	defer rows.Close()

	// 取得した行セットに対して繰り返し処理
	for rows.Next() {
		// 変数格納時にデータの型変換を裏側で処理してくれる
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 単一行クエリ
	// 1行しか返却しない場合、rows.Next()の処理を省略できる
	err = db.QueryRow("SELECT name FROM user WHERE id=?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(name)

	// データ更新のステートメント
	//stmt2, err := db.Prepare("INSERT INTO user(id, name) VALUES(?, ?)")
	//stmt2, err := db.Prepare("UPDATE user SET id = ?, name = ? WHERE id = ?")
	stmt2, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt2.Exec(3)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId() // 挿入した行のIDを返却
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected() // 影響を受けた行数
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID=%d,affected=%d\n", lastId, rowCnt)

}
