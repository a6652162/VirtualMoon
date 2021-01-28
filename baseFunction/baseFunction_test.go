package baseFunction

import (
	"fmt"
	"testing"
	"time"
)

func Test_FormatDatetime(t *testing.T) {
  //日期格式转换为字符串
	str := FormatDatetime(time.Now(), "yyyy-mm/dd hh:nn:ss")
	fmt.Println("FormatDateTime:", str)
}

func Test_StrAndByte(t *testing.T) {
  //16进度格式的字符串
	str := "0102030405060708090A0BFF"
  //转换成对应数组
	b := StringToByte(str)
	fmt.Println("byte:", b)
  //转换成16进度格式的字符串
	fmt.Println("string:", BytesToByteStr(b))
  //增加前缀 0x 显示格式
	fmt.Println("string:", BytesToByteStr(b, " 0x"))
}
