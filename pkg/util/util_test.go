package util

import (
	"fmt"
	"testing"
)

func TestShowSubstr(t *testing.T)  {
	str := "abc我是不可不戒写的eft，测试字符串长度截取的函数，看看到底能否成功按需求截取"
	//fmt.Println(ShowSubstr(str, 26))
	fmt.Println(string([]rune(str)[0:20]))
}