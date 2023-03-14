package day25

import (
	"fmt"
	"math"
	"strings"
)

//*数值判断；常规循环实现;使用标记法：标记E/e；小数点；数字；
func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if s == "" {
		return false
	}

	dot_flag := false //?小数点标记
	e_flag := false   //?科学计数法e/E标记
	num_flag := false //?数字标记

	for i := 0; i < len(s); i++ {
		if s[i] == '-' || s[i] == '+' { //*只能出现在第一位或e/E后一位
			if i != 0 && (strings.ToLower(string(s[i-1])) != "e") {
				return false
			}
		} else if s[i] == '.' { //?小数点只能出现一次，并不能再e/E后面出现
			if dot_flag || e_flag { //*小数点或者e/E已经有了
				return false
			}
			dot_flag = true
		} else if s[i] == 'e' || s[i] == 'E' { //*e/E只能出现一次，且不能再开头或结尾
			if e_flag || !num_flag || i == len(s)-1 { //*e前面必须有数字
				return false
			}
			e_flag, num_flag = true, false
		} else if isDigit(s[i]) {
			num_flag = true
		} else {
			return false //*不符合为非法字符
		}
	}
	return true

}

//?面试题67：把字符串转化整数（不能使用库函数）
//*创建自己的Aoti函数
func strToInt(str string) int {
	if len(str) == 0 {
		return 0
	}
	i := 0    //*计数器
	sign := 1 //*符号位
	res := 0  //*结果
	//*去除左边空字符
	for str[i] == ' ' && i < len(str) {
		i++
		if i == len(str) {
			return 0
		}
	}
	//?符号判别
	if str[i] == '-' {
		sign = -1
		i++
	} else if str[i] == '+' {
		i++
	}

	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			res = res*10 + int(str[i]-'0')
			if res > (math.MaxInt32+1)/10 {
				if sign == 1 {
					return math.MaxInt32
				}
				return math.MinInt32
			}
		} else {
			break
		}
	}
	return sign * res
}

func isDigit(s byte) bool {
	n := "0123456789"
	for i := 0; i < len(n); i++ {
		if n[i] == s {
			return true
		}
	}
	return false
}

func strToInt2(str string) int {
	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i := 0; i < len(str); i++ {
		v := str[i]
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
			fmt.Println(result)
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}

		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * result
}
