package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"gin-swagger-demo/pkg/constants"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"gin-swagger-demo/pkg/setting"
)

// Setup Initialize the util
func Setup() {
	jwtUserSecret = []byte(setting.AppSetting.JwtUserSecret)
	jwtAdminSecret = []byte(setting.AppSetting.JwtAdminSecret)
}

func UUID() string {
	id := uuid.New()
	return strings.Replace(id.String(), "-", "", -1)
}

//转换时间戳
func StrToTime(toTime string) int {
	if toTime == "" {
		return 0
	}

	//时间 to 时间戳
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, _ := time.ParseInLocation(constants.TimeLayout, toTime, loc)
	return int(tt.Unix())
}

//时间戳转换时间
func TimeToStrTime(toTime int) string {
	if toTime < 1 {
		return ""
	}
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeobj := time.Unix(int64(toTime), 0)
	return timeobj.In(cstSh).Format(constants.TimeLayout)
}

//时间戳转换时间 1608880464 -> 12.1 14:30
func TimeToStrTimeShort(toTime int) string {
	timeLayout := "01.02 15:04"
	if toTime < 1 {
		return ""
	}
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeobj := time.Unix(int64(toTime), 0)
	return timeobj.In(cstSh).Format(timeLayout)
}

//时间戳转换时间 1608880464 -> 12月1日 周三
func TimeToWeek(toTime int) string {
	if toTime < 1 {
		return ""
	}
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeobj := time.Unix(int64(toTime), 0)
	date := timeobj.In(cstSh).Format("1月2日")

	weekMap := map[string]string{
		"Mon" : "周一",
		"Tue" : "周二",
		"Wed" : "周三",
		"Thu" : "周四",
		"Fri" : "周五",
		"Sat" : "周六",
		"Sun" : "周日",
	}
	week := timeobj.In(cstSh).Format("Mon")
	return date + " " + weekMap[week]
}

//时间戳转换日期
func TimeToStrDate(toTime int) string {
	if toTime < 1 {
		return ""
	}
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	timeobj := time.Unix(int64(toTime), 0)
	return timeobj.In(cstSh).Format(constants.DateLayout)
}

//时间戳转换
func TimeToStrShow(toTime int) string {
	ret := ""
	if toTime < 1 {
		return ret
	}
	nowTime := int(time.Now().Unix())
	diff := nowTime - toTime
	if diff < 0 {
		return ret
	}

	if diff < 60 {
		ret = "刚刚"
	} else if 60 <= diff && diff <= 3600 {
		ret = strconv.Itoa(diff/60) + "分钟前"
	} else if 3600 <= diff && diff <= 86400 {
		ret = strconv.Itoa(diff/3600) + "小时前"
	} else if 86400 <= diff && diff <= 86400*30 {
		ret = strconv.Itoa(diff/86400) + "天前"
	} else {
		var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
		timeobj := time.Unix(int64(toTime), 0)
		ret = timeobj.In(cstSh).Format(constants.DateLayout)
	}
	return ret
}

//md5 加密
func Md5String(str string) string {
	if str == "" {
		return ""
	}

	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return md5str
}

//去重切片
func RemoveDuplicateElement(languages []string) []string {
	if len(languages) < 1 {
		return nil
	}
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func RemoveDuplicateInt(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func OssUrl(key string) string {
	return setting.AliyunSetting.AliyunOssDomain + "/" + key
}

func DomainFullUrl(key string) string {
	return setting.AppSetting.PrefixUrl + "/" + key
}

// RandomString returns a random string with a fixed length
func RandomString(n int, allowedChars ...[]rune) string {
	var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var letters []rune

	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func GetEventMinFeeStr(fee int) string {
	if fee == 0 {
		return "免费"
	} else {
		res := fmt.Sprintf("%.2f", float64(fee)/100)
		return res + "元起"
	}
}

func FeeToStr(fee int) string {
	if fee == 0 {
		return "0"
	} else {
		res := fmt.Sprintf("%.2f", float64(fee)/100)
		return res
	}
}

func FilterEmoji(content string) string {
	new_content := ""
	for _, value := range content {
		_, size := utf8.DecodeRuneInString(string(value))
		if size <= 3 {
			new_content += string(value)
		}
	}
	return new_content
}


func ExcelExportData(header []string, data []map[string]interface{}, sheetName string, fileName string) string {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	style := &xlsx.Style{}
	style.Fill = *xlsx.NewFill("solid", "EFEFDE", "EFEFDE")
	style.Border = xlsx.Border{RightColor: "FF"}
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet(sheetName)
	row = sheet.AddRow()
	for i := 0; i < len(header); i++ { //looping from 0 to the length of the array
		cell = row.AddCell()
		cell.Value = header[i]
		cell.SetStyle(style)
	}
	for _, obj := range data {
		row = sheet.AddRow()
		for i := 0;i <len(header);i++{
			switch obj[header[i]].(type) {
			case string:
				obj[header[i]] = obj[header[i]]
				break
			case int:
				obj[header[i]] = strconv.Itoa(obj[header[i]].(int))
				break
			case float64:
				obj[header[i]] = strconv.FormatFloat(obj[header[i]].(float64), 'E', -1, 64)
				break
			}
			cell = row.AddCell()
			cell.Value = obj[header[i]].(string)
		}
	}
	url := fileName+".xlsx"
	file.Save(url)
	return url
}


// GetUidFromHeader 获取当前登录用户的 UID
func GetUidFromHeader(c *gin.Context) (uid int,isAdmin bool,err error) {
	const BEARER_SCHEMA_USER = "Bearer "
	//const BEARER_SCHEMA_ADMIN = "Aearer "

	authHeader := c.GetHeader("Authorization")

	if len(authHeader) < len(BEARER_SCHEMA_USER) {
		return 0,false,fmt.Errorf("Token is empty")
	}

	token := authHeader[len(BEARER_SCHEMA_USER):]
	if token == "" {
		return 0,false,nil
	}

	tokenType := authHeader[0:1] //第一个字节用来标识是用户的 token 还是 admin 的

	if tokenType == "A" {
		res, err := ParseAdminToken(token)
		if err != nil {
			return  0,false,err
		}
		return res.Id,true,nil
	} else{
		res, err := ParseUserToken(token)
		if err != nil {
			return  0,false,err
		}
		return res.Id,false,nil
	}
}

// 10进制转成26进制字母 A:0,B:1
func DecToChar(num int64) string {
	s26 := strconv.FormatInt(num, 26)

	res := ""
	for _,c := range  s26 {
		if c <= 57 {
			res += string(c+17)
		} else{
			res += string(c-22)
		}
	}
	return res
}

func ShowSubstr(s string, l int) string {
	if len(s) <= l {
		return s
	}
	ss, sl, rl, rs := "", 0, 0, []rune(s)
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			rl = 1
		} else {
			rl = 2
		}

		if sl + rl > l {
			break
		}
		sl += rl
		ss += string(r)
	}
	return ss
}