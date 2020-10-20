package until

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

//定义数据库连接
func initDbDeatil() string {
    dbuser := ReadConf("dbuser")
    dbpass := ReadConf("dbpass")
    dbhost := ReadConf("dbhost")
    dbport := ReadConf("dbport")
    dbname := ReadConf("dbname")
    dbsuffix := ReadConf("dbsuffix")

    dbDeatil := dbuser + ":" + dbpass + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?" + dbsuffix
    return dbDeatil
}

//查询阴历记录
func SelectYinLiData() {
    dbDeatil := initDbDeatil()
    db, err := sql.Open("mysql", dbDeatil)
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

//查询阳历记录
func SelectYangLiData() {
    dbDeatil := initDbDeatil()
    db, err := sql.Open("mysql", dbDeatil)
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

//查询所有记录
func SelectData() {
    dbDeatil := initDbDeatil()
    db, err := sql.Open("mysql", dbDeatil)
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

//根据id查询单条记录
func SelectDataById(id int) {
    dbDeatil := initDbDeatil()
    db, err := sql.Open("mysql", dbDeatil)
    checkErr(err)
    defer db.Close()
    rows, err := db.Query("SELECT * FROM reminder where id=?", id)

    checkErr(err)
    for rows.Next() {
        var id, date, note, alert, yinli string
        err = rows.Scan(&id, &date, &note, &alert, &yinli)
        checkErr(err)
        fmt.Printf(id + "\t" + date + "\t" + note + "\t" + alert + "\t" + yinli + "\n")

    }

}

//删除记录
func DeleteData(id int) {
    dbDeatil := initDbDeatil()
    db, err := sql.Open("mysql", dbDeatil)
    checkErr(err)
    defer db.Close()
    stmt, err := db.Prepare("delete from reminder where id=?")
    checkErr(err)

    res, err := stmt.Exec(id)
    checkErr(err)

    _, err = res.RowsAffected()
    checkErr(err)

    fmt.Println("删除记录完成。")

}

//添加记录
func InsertData(Date, Note string, Alert, Yinli int) {
    dbDeatil := initDbDeatil()
    db, err := sql.Open("mysql", dbDeatil)
    checkErr(err)
    defer db.Close()
    //插入数据
    stmt, err := db.Prepare("INSERT reminder SET date=?,note=?,alert=?,yinli=?")
    checkErr(err)

    res, err := stmt.Exec(Date, Note, Alert, Yinli)
    checkErr(err)

    _, err = res.LastInsertId()
    checkErr(err)
    fmt.Println("添加记录完成。")
}

//更新记录
func UpdateData(Date, Note string, Alert, Yinli, Id int) {
    dbDeatil := initDbDeatil()
    db, err := sql.Open("mysql", dbDeatil)
    defer db.Close()

    //更新数据
    stmt, err := db.Prepare("update reminder set date=?,note=?,alert=?,yinli=?  where id=? ")
    checkErr(err)

    res, err := stmt.Exec(Date, Note, Alert, Yinli, Id)
    checkErr(err)

    _, err = res.RowsAffected()
    checkErr(err)

    fmt.Println("修改完成，如下为修改后的记录")
}

//检查异常
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
