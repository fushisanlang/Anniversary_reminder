package until

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/robfig/cron"
)

//实际通知函数，通过日期去数据库查询数据
func Alert() {
	yangli := getTodayYangli()
	yinli := getTodayYinli()
	selectYangliAlertData(yangli)
	selectYinliAlertData(yinli)
}

//通过日期查询阳历数据，有符合条件的，推送到微信中
func selectYangliAlertData(yanglidate string) {
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT note,stopdate FROM yangli where startdate<=? and stopdate >=?", yanglidate, yanglidate)

	checkErr(err)
	for rows.Next() {
		var dbnote, dbstopdate string
		err = rows.Scan(&dbnote, &dbstopdate)
		checkErr(err)
		//推送内容
		alertString := dbnote + "\t" + dbstopdate + "\n"
		//推送
		SendWx(alertString)

	}
}

//通过日期查询阴历数据，有符合条件的，推送到微信中
func selectYinliAlertData(yinlidate string) {
	fmt.Println(yinlidate)
	dbDeatil := initDbDeatil()
	db, err := sql.Open("sqlite3", dbDeatil)
	checkErr(err)
	defer db.Close()
	rows, err := db.Query("SELECT note,stopdate FROM yinli where startdate<=? and stopdate >=?", yinlidate, yinlidate)

	fmt.Println("SELECT note,stopdate FROM yinli where startdate<=? and stopdate >=?", yinlidate, yinlidate)
	checkErr(err)
	for rows.Next() {
		var dbnote, dbstopdate string
		err = rows.Scan(&dbnote, &dbstopdate)
		checkErr(err)
		alertString := dbnote + "\t" + dbstopdate + "\n"
		//推送
		SendWx(alertString)
	}

}

//定时任务函数
func AlertCron() {
	cronvalue := ReadConf("cronvalue")
	log.Println("Starting Cron...")
	c := cron.New()
	//创建一个定时任务
	c.AddFunc(cronvalue, Alert)
	c.Start()
	//死循环
	select {}
}
