package dataexport

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

func Prize() {
	f, _ := xlsx.OpenFile("C:/Users/xs253/Desktop/奖励.xlsx")
	ss := f.Sheets[0]
	for _, value := range ss.Rows {
		fmt.Println(value.Cells)
	}
}

type Balance struct {
	Uid      int
	FromId   int
	FromType int
	RDesc    string
	Number   int
}

// AddBalance 添加余额
func AddBalance(fname string) {
	f, _ := xlsx.OpenFile(fname)
	res := []Balance{}
	for _, v := range f.Sheets[0].Rows {
		tmp := Balance{}
		tmp.Uid, _ = v.Cells[0].Int()    //用户
		tmp.Number, _ = v.Cells[1].Int() //充值金额
		tmp.RDesc = v.Cells[2].Value     //活动来源
		//		tmp.FromId, _ = v.Cells[1].Int()
		//		tmp.FromType, _ = v.Cells[2].Int()
		//		tmp.BN, _ = v.Cells[3].Int()
		//		tmp.An, _ = v.Cells[4].Int()
		//		tmp.RDesc = v.Cells[5].Value
		//		tmp.Ti = v.Cells[6].Value
		res = append(res, tmp)
	}
	fmt.Println(res)

	sql1 := `
INSERT INTO balance_change_record (user_id,from_id,from_type,before_number,
	after_number,	record_desc,	create_time)
SELECT user_id ,?,2290,number,number+?,?,NOW() FROM balance WHERE user_id =?  LIMIT 1;
`
	sql2 := `UPDATE balance SET number=number+? WHERE user_id = ?;`
	_, _ = sql1, sql2
	o := orm.NewOrm()
	var err error
	o.Begin()
	defer func() {
		if err != nil {
			o.Rollback()
			return
		}
		o.Commit()
	}()

	is := 0
	for _, v := range res {
		_, err = o.Raw(sql1, v.FromId, v.Number, v.RDesc, v.Uid).Exec()
		if err != nil {
			return
		}
		_, err = o.Raw(sql2, v.Number, v.Uid).Exec()
		if err != nil {
			return
		}
		is++
	}
	fmt.Println("更新行数", is)
}