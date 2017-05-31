package number

// getNumbers 取出字符串中的数字显示
func getNumbers(s string) []string {
	rs := []rune(s)
	var nums []string
	var n []rune
	for _, v := range rs {
		if v <= '9' && v >= '0' {
			n = append(n, v)
		} else if v == '.' || v == '%' {
			n = append(n, v)
		} else {
			if len(n) > 0 {
				nums = append(nums, string(n))
				n = nil
			}
		}
	}
	//	s := `公积金 1.10 倍 ( 3.58%)`
	//	fmt.Println(nums) // 1.10,3.58%
	return nums
}
