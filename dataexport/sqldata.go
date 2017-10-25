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

var (
	purchase = `SELECT
ua.account '手机号',
ui.user_name '用户名',
CASE @uid WHEN u.id THEN @t:=@t+1 ELSE @t:=1 END '投资次数' ,
@uid:=ui.user_id '用户id',
if(ui.inviter=0,'无','有'),
if(ui.inviter=0,'',(SELECT user_name FROM user_infos WHERE user_id =ui.inviter)),
if(ui.inviter=0,'',(SELECT account FROM user_account WHERE user_id =ui.inviter)),
u.platform '渠道',
ui.id_card '身份证号',
b.bid_name'标名',
b.days_limit'期限',
CAST(pr.number/100 AS UNSIGNED) '投资金额',
CAST(pr.pay_number/100 AS UNSIGNED) '实付金额',
pr.create_time '投资时间'

 FROM purchase_record pr
LEFT JOIN user_account ua ON pr.user_id=ua.user_id
LEFT JOIN user_infos ui ON pr.user_id=ui.user_id
LEFT JOIN bids b ON pr.bid_id=b.id
LEFT JOIN users u ON pr.user_id=u.id
, (SELECT @t:=1,@uid:=0) rn
WHERE u.state !=102
ORDER BY pr.user_id,pr.create_time `
	platformPurchase = `SELECT
IF(@uid=a.user_id, @t:=@t+1,IF(a.user_id is NULL ,0,@t:=1))'z',
@uid:=a.user_id '用户id',
a.*
FROM (
SELECT
ua.account '手机号',
ui.user_name '用户名',
pr.user_id,
if(ui.inviter=0,'无','有'),
if(ui.inviter=0,'',(SELECT user_name FROM user_infos WHERE user_id =ui.inviter)),
if(ui.inviter=0,'',(SELECT account FROM user_account WHERE user_id =ui.inviter)),
u.create_time '注册时间',
u.platform ,
ui.id_card '身份证号',
b.bid_name'标名',
b.days_limit'期限',
CAST(pr.number/100 AS UNSIGNED) '投资金额',
CAST(pr.pay_number/100 AS UNSIGNED) '实付金额',
pr.create_time '投资时间'

 FROM users u
LEFT JOIN user_account ua ON u.id=ua.user_id
LEFT JOIN user_infos ui ON u.id=ui.user_id
LEFT JOIN purchase_record pr ON pr.user_id=u.id
LEFT JOIN bids b ON pr.bid_id=b.id
WHERE u.state !=102
AND u.platform >0
ORDER BY pr.user_id,pr.create_time
)a ,(SELECT @t:=1,@uid:=0) rn

ORDER BY platform `
	charge = `SELECT
ui.user_name,
ua.account,
cr.number/100,
cr.create_time ,
cr.pay_no
 FROM charge_record cr
LEFT JOIN user_infos ui ON cr.user_id=ui.user_id
LEFT JOIN user_account ua ON cr.user_id =ua.user_id
LEFT JOIN users u ON cr.user_id=u.id
 where
 cr.pay_way!=1501
AND ua.account!=18000000000
AND u.state !=102
AND DATE(cr.create_time)= ?
ORDER BY  cr.number,cr.pay_no;`
	balanceRecord = `SELECT
ui.user_name,
ua.account,
(cr.after_number-cr.before_number)/100,
cr.create_time,
cr.record_desc
 FROM balance_change_record cr
LEFT JOIN user_infos ui ON cr.user_id=ui.user_id
LEFT JOIN user_account ua ON cr.user_id =ua.user_id
LEFT JOIN users u ON cr.user_id=u.id
 where
cr.from_type=2290
AND ua.account!=18000000000
AND u.state !=102
AND DATE(cr.create_time)= ?
ORDER BY cr.record_desc ,cr.create_time;
`
)

func PurchaseRecord() {
	yestoday := time.Now().Format("2006-01-02")
	fname := "C:/Users/xs253/Desktop/投资" + yestoday + ".xlsx"
	sliceToFile(paramToInterface(sqlData(fname, purchase)), fname)
	sliceToFile(paramToInterface(sqlData(fname, platformPurchase)), fname)
}

func ActivityRecord() {
	yestoday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	fname := "C:/Users/xs253/Desktop/活动" + yestoday + ".xlsx"
	// 充值记录
	cv := sqlData(fname, charge, yestoday)
	// 余额记录
	bv := sqlData(fname, balanceRecord, yestoday)
	sliceToFile(paramToInterface(append(cv, bv...)), fname)
}

func sqlData(fname, sql string, args ...interface{}) (v []orm.ParamsList) {
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
