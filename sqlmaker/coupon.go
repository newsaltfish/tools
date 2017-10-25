package sqlmaker

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"
)

// Coupon 优惠券
type Coupon struct {
	Id         int       //优惠券id
	CouponType int       // 类型
	CouponRule string    // 规则
	CouponName string    // 名称
	CouponDesc string    // 描述
	CreateTime time.Time // 创建时间
}

// MCouponRule 优惠券规则
type MCouponRule struct {
	MoneyMin       int
	Minus          int
	TimeNow        int
	DayMin         int
	TimeRangeBegin string // 开始有效时间
	TimeRangeEnd   string // 结束有效时间
}

// RCouponRule 加息券规则
type RCouponRule struct {
	MoneyMin       int
	TimeNow        int
	DayMin         int
	RateRanges     []rateRule
	TimeRangeBegin string // 开始有效时间
	TimeRangeEnd   string // 结束有效时间
}

// rateRule 加息规则
type rateRule struct {
	Day  int
	Rate int
}

// GetCouponRRule 读取加息券
func GetCouponRRule(fname string) []Coupon {
	f, _ := xlsx.OpenFile(fname)
	res := make([]Coupon, 0)
	tmap := make(map[string]interface{})
	for key, value := range f.Sheets[0].Rows {
		if key == 0 {
			continue
		}
		t := RCouponRule{}
		r := rateRule{}
		t.DayMin, _ = value.Cells[1].Int()
		t.MoneyMin, _ = value.Cells[2].Int()
		r.Rate, _ = value.Cells[3].Int()
		t.TimeNow, _ = value.Cells[4].Int()
		//金额*100
		t.MoneyMin = t.MoneyMin * 100
		r.Rate = r.Rate * 1000
		t.RateRanges = []rateRule{r}
		//有效期
		days := " (有效期" + strconv.Itoa(t.TimeNow) + "天)"
		rule := ""
		b, _ := json.Marshal(t)
		json.Unmarshal(b, &tmap)
		if t.TimeNow == 0 {
			t.TimeRangeBegin = value.Cells[6].Value
			t.TimeRangeEnd = value.Cells[7].Value
			days = "(有效期至" + t.TimeRangeEnd[:10] + ")"
			tmap["TimeRangeBegin"] = value.Cells[6].Value
			tmap["TimeRangeEnd"] = value.Cells[7].Value
			delete(tmap, "TimeNow")
		}
		b, _ = json.Marshal(tmap)
		rule = string(b)

		desc := "(" + value.Cells[5].Value + ") " + strconv.Itoa(r.Rate/1000) + "%加息券 "
		if t.DayMin != 0 {
			desc += strconv.Itoa(t.DayMin) + "天及以上的标 "
		}
		if t.MoneyMin != 0 {
			desc += " 满" + strconv.Itoa(t.MoneyMin/100) + "元可用 "
		}
		desc += days
		tmp := Coupon{
			CreateTime: time.Now(),
			CouponDesc: desc,
			CouponName: strconv.Itoa(r.Rate/1000) + "%抵用券",
			CouponType: 502,
			CouponRule: rule,
		}
		tmp.Id, _ = value.Cells[0].Int()
		res = append(res, tmp)
	}
	return res

}

// GetCouponMRule 读取优惠券
func GetCouponMRule(fname string) []Coupon {
	f, _ := xlsx.OpenFile(fname)
	res := make([]Coupon, 0)
	tmap := make(map[string]interface{})
	for key, value := range f.Sheets[0].Rows {
		if key == 0 {
			continue
		}
		t := MCouponRule{}
		t.DayMin, _ = value.Cells[1].Int()
		t.MoneyMin, _ = value.Cells[2].Int()
		t.Minus, _ = value.Cells[3].Int()
		t.TimeNow, _ = value.Cells[4].Int()
		//金额*100
		t.MoneyMin = t.MoneyMin * 100
		t.Minus = t.Minus * 100
		//有效期
		days := " (有效期" + strconv.Itoa(t.TimeNow) + "天)"
		rule := ""
		b, _ := json.Marshal(t)
		json.Unmarshal(b, &tmap)
		if t.TimeNow == 0 {
			t.TimeRangeBegin = value.Cells[6].Value
			t.TimeRangeEnd = value.Cells[7].Value
			days = "(有效期至" + t.TimeRangeEnd[:10] + ")"
			tmap["TimeRangeBegin"] = value.Cells[6].Value
			tmap["TimeRangeEnd"] = value.Cells[7].Value
			delete(tmap, "TimeNow")
		}
		b, _ = json.Marshal(tmap)
		rule = string(b)

		desc := "(" + value.Cells[5].Value + ") " + strconv.Itoa(t.Minus/100) + "元 " +
			strconv.Itoa(t.DayMin) + "天及以上的标 "
		if t.MoneyMin != 0 {
			desc += " 满" + strconv.Itoa(t.MoneyMin/100) + "元可用 "
		}
		desc += days
		tmp := Coupon{
			CreateTime: time.Now(),
			CouponDesc: desc,
			CouponName: strconv.Itoa(t.Minus/100) + "元抵用券",
			CouponType: 501,
			CouponRule: rule,
		}
		tmp.Id, _ = value.Cells[0].Int()
		res = append(res, tmp)
	}
	return res
}

// CouponToFile 优惠券
func CouponToFile(fname string, p []Coupon) {
	f, _ := xlsx.OpenFile(fname)
	if f == nil {
		f = xlsx.NewFile()
	}
	sname := "Sheet" + strconv.Itoa(len(f.Sheets)+1)
	f.AddSheet(sname)
	for _, value := range p {
		row := f.Sheet[sname].AddRow()
		row.AddCell().SetInt(value.Id)
		row.AddCell().SetInt(value.CouponType)
		row.AddCell().Value = value.CouponRule
		row.AddCell().Value = value.CouponName
		row.AddCell().Value = value.CouponDesc
		row.AddCell().SetDateWithOptions(value.CreateTime, xlsx.DateTimeOptions{Location: time.Local, ExcelTimeFormat: "yyyy-mm-dd hh:mm:ss"})
	}
	f.Save(fname)
}
