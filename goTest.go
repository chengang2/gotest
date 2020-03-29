package main

import (
	"fmt"
	"strings"
)

//全角转换半角
func DBCtoSBC(s string) string {
	retstr := ""
	for _, i := range s {
		inside_code := i
		if inside_code == 12288 {
			inside_code = 32
		} else {
			inside_code -= 65248
		}
		if inside_code < 32 || inside_code > 126 {
			retstr += string(i)
		} else {
			retstr += string(inside_code)
		}
	}
	return retstr
}

func main() {


	b := " 总汞（以Hg计） "
	c := "总汞(以Hg计)"
	d := DBCtoSBC(b)
	fmt.Println(strings.TrimSpace(b),b,d,c,d==c,b==c)
}
