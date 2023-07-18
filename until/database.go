package until

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// 定义数据库连接
func initDbDeatil() string {
	dbPath := ReadConf("db_path")
	dbDeatil := fmt.Sprintf("%s?_foreign_keys=on", dbPath)
	return dbDeatil
}

// 查询阴历记录
func SelectYinLiData() {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM yinli")
	checkErr(err)

	for rows.Next() {
		var startdate, stopdate, detail string
		err = rows.Scan(&startdate, &stopdate, &detail)
		checkErr(err)
		fmt.Printf(stopdate + "\t" + detail + "\n")
	}
}

// 查询阳历记录
func SelectYangLiData() {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM yangli")
	checkErr(err)

	for rows.Next() {
		var startdate, stopdate, detail string
		err = rows.Scan(&startdate, &stopdate, &detail)
		checkErr(err)
		fmt.Printf(stopdate + "\t" + detail + "\n")
	}
}

// 查询所有记录
func SelectData() {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM reminder")
	checkErr(err)

	for rows.Next() {
		var id, date, note, alert, yinli string
		err = rows.Scan(&id, &date, &note, &alert, &yinli)
		checkErr(err)
		fmt.Printf(id + "\t" + date + "\t" + note + "\t" + alert + "\t" + yinli + "\n")
	}
}

// 根据id查询单条记录
func SelectDataById(id int) {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM reminder WHERE id=?", id)
	checkErr(err)

	for rows.Next() {
		var id, date, note, alert, yinli string
		err = rows.Scan(&id, &date, &note, &alert, &yinli)
		checkErr(err)
		fmt.Printf(id + "\t" + date + "\t" + note + "\t" + alert + "\t" + yinli + "\n")
	}
}

// 删除记录
func DeleteData(id int) {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM reminder WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("删除记录完成。")
}

// 添加记录
func InsertData(Date, Note string, Alert, Yinli int) {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()

	// 插入数据
	stmt, err := db.Prepare("INSERT INTO reminder (date, note, alert, yinli) VALUES (?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(Date, Note, Alert, Yinli)
	checkErr(err)

	_, err = res.LastInsertId()
	checkErr(err)

	fmt.Println("添加记录完成。")
}

// 更新记录
func UpdateData(Date, Note string, Alert, Yinli, Id int) {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	defer db.Close()

	// 更新数据
	stmt, err := db.Prepare("UPDATE reminder SET date=?, note=?, alert=?, yinli=? WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(Date, Note, Alert, Yinli, Id)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("修改完成，如下为修改后的记录")
}

// 检查异常
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
