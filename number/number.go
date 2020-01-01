package number

// getNumbers 取出字符串中的数字显示
func getNumbers(s string) []string {
	var nums []string
	var n []rune
	for i := range s {
		if s[i] <= '9' && s[i] >= '0' {
			n = append(n, rune(s[i]))
		} else if s[i] == '.' || s[i] == '%' {
			n = append(n, rune(s[i]))
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
