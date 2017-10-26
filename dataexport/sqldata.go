package dataexport

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/tealeg/xlsx"
	"github.com/wzshiming/ffmt"
)

func PurchaseRecord(fname string, purchase string, platformPurchase string) {
	sliceToFile(paramToInterface(sqlData(fname, purchase)), fname)
	sliceToFile(paramToInterface(sqlData(fname, platformPurchase)), fname)
}

// SqlToFile sql查询结果导出
func SqlToFile(filename string, sql string, args ...interface{}) {
	sliceToFile(paramToInterface(sqlData(sql)), filename)
}

// SqlToSilce sql返回数组
func SqlToSilce(sql string, args ...interface{}) [][]interface{} {
	return paramToInterface(sqlData(sql, args...))
}

// SliceToFile 数组保持到文件
func SliceToFile(req [][]interface{}, fname string) {
	sliceToFile(req, fname)
}

// ActivityRecord 活动
func ActivityRecord(fname string, charge string, balanceRecord string) {
	yestoday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	// 充值记录
	cv := sqlData(charge, yestoday)
	// 余额记录
	bv := sqlData(balanceRecord, yestoday)
	sliceToFile(paramToInterface(append(cv, bv...)), fname)
}

func sqlData(sql string, args ...interface{}) (v []orm.ParamsList) {
	o := orm.NewOrm()
	_, err := o.Raw(sql, args...).ValuesList(&v)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	return
}

func paramToInterface(r []orm.ParamsList) (res [][]interface{}) {
	for _, value := range r {
		vl := reflect.ValueOf(value)
		if vl.Kind() != reflect.Slice {
			continue
		}
		tmp := []interface{}{}
		for i := 0; i < vl.Len(); i++ {
			tmp = append(tmp, vl.Index(i))
		}
		res = append(res, tmp)
	}
	return
}

func sliceToFile(s [][]interface{}, fname string) {
	f, err := xlsx.OpenFile(fname)
	if err != nil {
		ffmt.Mark(err)
		f = xlsx.NewFile()
	}
	sname := "Sheet" + strconv.Itoa(len(f.Sheets)+1)
	f.AddSheet(sname)
	for _, v := range s {
		row := f.Sheet[sname].AddRow()
		for _, value := range v {
			cell := row.AddCell()
			cell.Value = fmt.Sprint(value)
		}
	}
	f.Save(fname)
	fmt.Println(fname + " .......... done")
}
