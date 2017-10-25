package dataexport

import (
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
)

func Test_sendPrize(t *testing.T) {

	o := orm.NewOrm()
	o.Raw("").Prepare()
	getFileValue("C:/Users/xs253/Desktop/奖励.xlsx")
}

func getFileValue(path string) [][]string {
	f, _ := xlsx.OpenFile(path)
	ss := f.Sheets[0]
	res := make([][]string, 0, len(ss.Rows))
	for _, value := range ss.Rows {
		rows := make([]string, 0, len(value.Cells))
		for _, v := range value.Cells {
			rows = append(rows, v.Value)
		}
		res = append(res, rows)
	}
	return res
}
