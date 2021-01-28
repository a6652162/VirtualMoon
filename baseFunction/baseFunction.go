package baseFunction

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//字节转换成整形
func BytesToInt(b []byte) int {
	var tb [4]byte
	var i int

	if len(b) < 4 {
		for i = 3; i > 4-len(b)-1; i-- {
			tb[i] = b[len(b)+i-4]
		}
		for ; i > -1; i-- {
			tb[i] = 0
		}
	} else {
		for i = 0; i < 4; i++ {
			tb[i] = b[i]
		}
	}
	bytesBuffer := bytes.NewBuffer(tb[:])

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//16进制数组转字符串显示
func BytesToByteStr(b []byte, ch ...string) string {
	var ss string
	length := len(b)
	for i := 0; i < length; i++ {
		if ch != nil {
			if ss == "" {
				ss = fmt.Sprintf("%s%s%02x", ss, ch[0], b[i])
			} else {
				ss = fmt.Sprintf("%s%s%02x", ss, ch[0], b[i])
			}
		} else {
			ss = fmt.Sprintf("%s%02x", ss, b[i])

		}
	}
	return ss
}

//16进制字符串转16进制字节数组
func StringToByte(value string) []byte {
	retbyte, _ := hex.DecodeString(value)
	return retbyte
}

//字符串转16进制字节数组
func StrToByteArr(value string) []byte {

	dst, _ := hex.DecodeString(value)
	return dst

}

//日期格式转换
func FormatDatetime(t time.Time, format string) string {
	format = strings.ToLower(format)
	//year
	if strings.ContainsAny(format, "y") {

		year := strconv.Itoa(t.Year())

		if strings.Count(format, "yy") == 1 && strings.Count(format, "y") == 2 {
			format = strings.Replace(format, "yy", year[2:], 1)
		} else if strings.Count(format, "yyyy") == 1 && strings.Count(format, "y") == 4 {
			format = strings.Replace(format, "yyyy", year, 1)
		} else {
			panic("format year error! please 'yyyy' or 'yy'")
		}
	}

	//month
	if strings.ContainsAny(format, "m") {

		var month string

		if int(t.Month()) < 10 {
			month = "0" + strconv.Itoa(int(t.Month()))
		} else {
			month = strconv.Itoa(int(t.Month()))
		}

		if strings.Count(format, "mm") == 1 && strings.Count(format, "m") == 2 {
			format = strings.Replace(format, "mm", month, 1)
		} else {
			panic("format month error! please 'mm'")
		}
	}

	//day
	if strings.ContainsAny(format, "d") {

		var day string

		if t.Day() < 10 {
			day = "0" + strconv.Itoa(t.Day())
		} else {
			day = strconv.Itoa(t.Day())
		}

		if strings.Count(format, "dd") == 1 && strings.Count(format, "d") == 2 {
			format = strings.Replace(format, "dd", day, 1)
		} else {
			panic("format day error! please 'dd'")
		}
	}

	//hour
	if strings.ContainsAny(format, "h") {

		var hour string

		if t.Hour() < 10 {
			hour = "0" + strconv.Itoa(t.Hour())
		} else {
			hour = strconv.Itoa(t.Hour())
		}

		if strings.Count(format, "hh") == 1 && strings.Count(format, "h") == 2 {
			format = strings.Replace(format, "hh", hour, 1)
		} else {
			panic("format hour error! please 'hh'")
		}
	}

	//minute
	if strings.ContainsAny(format, "n") {

		var minute string

		if t.Minute() < 10 {
			minute = "0" + strconv.Itoa(t.Minute())
		} else {
			minute = strconv.Itoa(t.Minute())
		}
		if strings.Count(format, "nn") == 1 && strings.Count(format, "n") == 2 {
			format = strings.Replace(format, "nn", minute, 1)
		} else {
			panic("format minute error! please 'nn'")
		}
	}

	//second
	if strings.ContainsAny(format, "s") {

		var second string

		if t.Second() < 10 {
			second = "0" + strconv.Itoa(t.Second())
		} else {
			second = strconv.Itoa(t.Second())
		}

		if strings.Count(format, "ss") == 1 && strings.Count(format, "s") == 2 {
			format = strings.Replace(format, "ss", second, 1)
		} else {
			panic("format second error! please 'ss'")
		}
	}

	//Nanosecond
	if strings.ContainsAny(format, "z") {

		var nanosecond string
		var haomiao int
		haomiao = t.Nanosecond() / 1000000
		if haomiao > 99 {
			nanosecond = strconv.Itoa(haomiao)
		} else {
			if haomiao > 9 {
				nanosecond = "0" + strconv.Itoa(haomiao)
			} else {
				nanosecond = "00" + strconv.Itoa(haomiao)
			}
		}

		if strings.Count(format, "zzz") == 1 && strings.Count(format, "z") == 3 {
			format = strings.Replace(format, "zzz", nanosecond, 1)
		} else {
			panic("format nanosecond error! please 'zzz'")
		}
	}

	return format
}
