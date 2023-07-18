package main

import (
	"Anniversary_reminder/until"
	"flag"
	"fmt"
)

//cli第一级菜单
//通过判断输入，调用不同函数或进入下级菜单。
func PrintParameter() {
	fmt.Printf("\n\n1. 查询记录\n2. 修改记录\n3. 删除记录\n4. 添加记录\n-----------\n")
	var Parameter int
	fmt.Scanln(&Parameter)
	switch Parameter {
	case 1:
		SelectParameter()
	case 2:
		UpdateParameter()
	case 3:
		DeleteParameter()
	case 4:
		InsertParameter()
	default:
		PrintParameter()
	}

}

//cli查询操作菜单
//通过判断输入，调用不同函数或进入下级菜单。
func SelectParameter() {
	var Parameter int
	fmt.Printf("\n1. 阳历纪念日\n2. 阴历纪念日\n3. 所有纪念日\n")
	fmt.Scanln(&Parameter)
	switch Parameter {
	case 1:
		until.SelectYangLiData()
		PrintParameter()
	case 2:
		until.SelectYinLiData()
		PrintParameter()
	case 3:
		until.SelectData()
		PrintParameter()
	default:
		PrintParameter()
	}
}

//cli添加操作菜单
//通过判断输入，调用不同函数或进入下级菜单。
func InsertParameter() {
	fmt.Printf("\n---------------\n添加记录\n---------------\n\n")
	var Date, Note string
	var Alert, Yinli int
	fmt.Println("输入日期（YYYY-MM-DD）")
	fmt.Scanln(&Date)
	fmt.Println("输入通知内容")
	fmt.Scanln(&Note)
	fmt.Println("输入提前通知天数")
	fmt.Scanln(&Alert)
	fmt.Println("输入阴历阳历（0.阳历/1.阴历）")
	fmt.Scanln(&Yinli)
	until.InsertData(Date, Note, Alert, Yinli)
	PrintParameter()

}

//cli删除操作菜单
//通过判断输入，调用不同函数或进入下级菜单。
func DeleteParameter() {
	until.SelectData()
	fmt.Printf("\n---------------\n删除记录\n---------------\n\n")
	var Id int
	var Parameter string
	fmt.Println("输入id")
	fmt.Scanln(&Id)
	until.SelectDataById(Id)
	fmt.Println("是否删除？(y/n)")
	fmt.Scanln(&Parameter)
	if Parameter == "y" {
		until.DeleteData(Id)
		PrintParameter()

	} else {
		PrintParameter()
	}

}
func UpdateParameter() {
	until.SelectData()
	fmt.Printf("\n---------------\n更新记录\n---------------\n\n")
	var Date, Note string
	var Alert, Yinli, Id int
	fmt.Println("输入id")
	fmt.Scanln(&Id)
	until.SelectDataById(Id)
	fmt.Println("输入日期（YYYY-MM-DD）")
	fmt.Scanln(&Date)
	fmt.Println("输入通知内容")
	fmt.Scanln(&Note)
	fmt.Println("输入提前通知天数")
	fmt.Scanln(&Alert)
	fmt.Println("输入阴历阳历（0.阳历/1.阴历）")
	fmt.Scanln(&Yinli)
	until.UpdateData(Date, Note, Alert, Yinli, Id)
	until.SelectDataById(Id)
	PrintParameter()

}

//帮助显示
func PrintHelp() {
	fmt.Println("./Anniversary_reminder -cli on 开启命令行")
	fmt.Println("./Anniversary_reminder -service start 启动定时任务")
}

//主函数
func main() {
	//定义cli和service变量
	var cli = flag.String("cli", "off", "cli on/off")
	var service = flag.String("service", "stop", "start")

	//获取参数输入
	flag.Parse()

	//判断参数输入

	if *cli == "on" && *service != "start" {
		//显示cli菜单
		PrintParameter()
	} else if *cli != "on" && *service == "start" {
		//开启定时任务进程
		until.AlertCron()
	} else if *cli == "on" && *service == "start" {
		//提示错误输入
		fmt.Println("只能使用cli或service参数中的一个，不可以一起使用。")
	} else {
		//显示帮助
		until.AlertCron()
	}
}
