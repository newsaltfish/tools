package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"
	"tools/dataexport"
	"tools/email"

	"golang.org/x/crypto/ssh"

	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
	"github.com/wzshiming/ffmt"

	yml "gopkg.in/yaml.v2"
)

func initSql() {
	bs, err := ioutil.ReadFile("./sql/sql.yml")
	fmt.Println(err)
	res := make(map[string]string)
	yml.UnmarshalStrict(bs, &res)
}
func main() {
	initMysqlCon()
	//	  dataexport.PurchaseRecord()
	//	AddBalance("./投资奖励.xlsx")
	//	sqlmaker.CouponToFile("./sqlmaker/coupon/rate.xlsx",
	//		sqlmaker.GetCouponRRule("./sqlmaker/coupon/加息券.xlsx"))
	//	email.Send()
	//	toFile(couponMRule())
	desktop := "C:\\Users\\xs253\\Desktop\\"
	fname = desktop + "activity\\活动2017-10-25.xlsx"
	email.GetMessage("xs@dev999.com", "253419372@qq.com",
		"zhy@dev999.com", "今日头条数据测试", "")
	email.TencentExmailSend()
}

// Activity 活动数据
func Activity() (fname string) {
	yestoday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	desktop := "C:\\Users\\xs253\\Desktop\\"
	fname = desktop + "activity\\活动" + yestoday + ".xlsx"
	dataexport.ActivityRecord(fname, res["charge"], res["balanceRecord"])
	return
}

// ToutiaoData 头条数据
func ToutiaoData() (fname string) {
	yestoday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	desktop := "C:\\Users\\xs253\\Desktop\\"
	s := [][]interface{}{[]interface{}{"包名", "日期", "人数"}}
	r := dataexport.SqlToSilce(res["ToutiaoRegister"], yestoday)
	s = append(s, r...)
	if len(r) == 0 {
		s = append(s, []interface{}{"无注册数据。。。", " ", " "})
	}
	p := dataexport.SqlToSilce(res["ToutiaoPurchase"], yestoday)
	s = append(s, []interface{}{"-包名-", "-日期-", "-投资人数-", "-投资总额-"})
	s = append(s, p...)
	if len(s) == 0 {
		s = append(s, []interface{}{"无投资数据。。。", " ", " "})
	}
	pmap := map[string]string{
		"1090": "toutiao", "1091": "toutiao2", "1092": "toutiao3",
		"1093": "toutiao4", "1094": "toutiao5", "1095": "toutiao6", "1096": "toutiao7",
		"1097": "toutiao8", "1098": "toutiao9", "1099": "toutiao10",
	}
	for k, v := range s {
		key := fmt.Sprint(v[0])
		if pack, ok := pmap[key]; ok {
			s[k][0] = pack
		}
	}
	fname = desktop + "toutiao\\头条" + yestoday + ".xlsx"
	dataexport.SliceToFile(s, fname)
	return fname
}
func initMysqlCon() {
	//  注册代理
	mysql.RegisterDial("ssh_dial", func(addr string) (conn net.Conn, err error) {
		ffmt.Mark(addr)
		PassWd := []ssh.AuthMethod{ssh.Password("hxh")}
		cli, err := ssh.Dial("tcp", "ssh.hxh-test.wzsm.studio:22", &ssh.ClientConfig{
			User: "hxh",
			Auth: PassWd,
		})
		if err != nil {
			ffmt.Mark(err)
			return
		}
		conn, err = cli.Dial("tcp", "mysql.hexianghang.cn:3306")
		if err != nil {
			ffmt.Mark(err)
			return
		}
		return
	})

	dataSource := "root:toor@tcp(db.gs.wzsm.studio:3306)/wjs_core?charset=utf8mb4&loc=Local"
	dataSource = "hxh:Hxh123123@ssh_dial(mysql.hexianghang.cn:3306)/wjs_core?charset=utf8mb4&loc=Local"

	err := orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		os.Exit(1)
	}

}
