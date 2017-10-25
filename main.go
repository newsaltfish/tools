package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"golang.org/x/crypto/ssh"

	"github.com/astaxie/beego/orm"
	"github.com/go-sql-driver/mysql"
	"github.com/wzshiming/ffmt"

	yml "gopkg.in/yaml.v2"
)

func main() {
	bs, err := ioutil.ReadFile("./sql/sql.yml")
	fmt.Println(err)
	res := make(map[string]string)
	yml.UnmarshalStrict(bs, &res)
	fmt.Println(res, yml.UnmarshalStrict(bs, &res))
	//  dataexport.PurchaseRecord()
	//	dataexport.ActivityRecord()
	//	AddBalance("./投资奖励.xlsx")
	//	sqlmaker.CouponToFile("./sqlmaker/coupon/rate.xlsx",
	//		sqlmaker.GetCouponRRule("./sqlmaker/coupon/加息券.xlsx"))
	//	email.Send()
	//	toFile(couponMRule())

}

func initMysql() {
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
