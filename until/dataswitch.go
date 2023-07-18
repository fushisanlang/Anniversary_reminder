package until

import (
	"regexp"
	"time"

	"github.com/nosixtools/solarlunar"
)

//阳历转阴历
func dataSwitch(solarDate string) string {
	return (solarlunar.SolarToSimpleLuanr(solarDate))

}

//获取今天阳历日期
func getTodayYangli() string {
	timeStr := time.Now().Format("0102")
	return timeStr
}

//获取今天阴历日期
func getTodayYinli() string {
	timeStr := time.Now().Format("2006-01-02")
	timestring := solarlunar.SolarToSimpleLuanr(timeStr)
	//正则替换，上一步结果是2020Y01M01D，通过正则，修改为0101，用于数据库比较
	timeRegexp, _ := regexp.Compile(`^\S{5}|M|D`)
	rep := timeRegexp.ReplaceAllString(timestring, "")
	return rep[:2] + rep[5:7]
}
