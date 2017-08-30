package main

import (


	"fmt"
	"github.com/go-sql-driver/mysql"

	"tools/dataexport"
	"time"
)

func main() {
	//go dataexport.PurchaseRecord()
	go dataexport.ActivityRecord()
	dsn := "user:passwd@ssh(addt)->u:psw@tcp(addr)/dbname"
	fmt.Println(mysql.ParseDSN(dsn))

	time.Sleep(30 * time.Second)
}
