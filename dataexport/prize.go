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
