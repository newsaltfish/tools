package str

import (
	"regexp"
	"strings"
)

// CamelName 蛇形转驼峰
func CamelName(base string) string {
	var r = make([]rune, 0, len(base))
	var b = []rune(base)
	for i := 0; i < len(b); i++ {
		if i == 0 && b[i] == '_' {
			continue
		}
		if i == len(b)-1 && b[i] == '_' {
			continue
		}
		if b[i] == '_' && i < len(b)-1 {
			if i > 0 && b[i] == '_' && i < len(b)-1 {
				if (b[i+1] <= 'Z' && b[i+1] >= 'A') || (b[i+1] >= 'a' && b[i+1] <= 'z') {
					r = append(r, b[i+1]-32)
					i++
				}
				continue
			}
			continue
		}
		r = append(r, b[i])
	}
	return string(r)
}

// SnakeName 蛇形转驼峰
func SnakeName(base string) string {
	var r = make([]byte, 0, len(base)*2)
	var b = []byte(base)
	for i := 0; i < len(b); i++ {
		if i > 0 && b[i] >= 'A' && b[i] <= 'Z' {
			r = append(r, '_', b[i]+32)
			continue
		}
		if i == 0 && b[i] >= 'A' && b[i] <= 'Z' {
			r = append(r, b[i]+32)
			continue
		}
		r = append(r, b[i])
	}
	return string(r)
}

// CamelNameB 驼峰转蛇形
func CamelNameB(s string) string {
	return strings.TrimLeft(regexp.MustCompile(`[A-Z]`).ReplaceAllStringFunc(s, func(v string) string {
		return "_" + strings.ToLower(v)
	}), "_")
}
