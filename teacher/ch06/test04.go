package main

import (
	"fmt"
	"regexp"
)

// regexp正则表达式包
func main() {
	// feewrfr3122@163.com
	// 1. 检索字符串
	// Contains 检查字符串是否包含指定的子字符串,如果包含则返回true，否则返回false
	fmt.Println(regexp.MustCompile("ll").MatchString("hello"))    // true
	fmt.Println(regexp.MustCompile("world").MatchString("hello")) // false

	// 2. 匹配邮箱（实用型正则，涵盖常见邮箱格式）
	email := "feewrfr3@163.com"
	// 说明：
	// - 本地部分：字母数字和 ._%+-
	// - 域名部分：字母数字和 - 与 . 组成，最后是至少2位的字母后缀
	// - 该正则适合教学与常规场景，并非 RFC 5322 的全部覆盖
	emailPattern := `^([A-Za-z0-9._%+-]+)@([A-Za-z0-9.-]+)\.([A-Za-z]{2,})$`
	re := regexp.MustCompile(emailPattern)
	fmt.Println("邮箱是否匹配:", re.MatchString(email)) // 期望 true
	if sub := re.FindStringSubmatch(email); len(sub) > 0 {
		fmt.Printf("local: %s, domain: %s, tld: %s\n", sub[1], sub[2], sub[3])
	}
}
