package utils

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"strings"
)

type StringX struct {
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func (StringX) InterfaceToStr(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func (StringX) Hex2DecStr(str interface{}) string {
	dex16 := TrimLeftZeroes(strings.TrimLeft(HexToString(str), "0x"))
	n := fmt.Sprintf("%x", dex16)
	return n
}

func (StringX) Hex2Dec(str interface{}) uint64 {
	dex16 := TrimLeftZeroes(strings.TrimLeft(HexToString(str), "0x"))
	nUint64, err := strconv.ParseUint(dex16, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return nUint64
}

func TrimLeftZeroes(hex string) string {
	idx := 0
	for ; idx < len(hex); idx++ {
		if hex[idx] != '0' {
			break
		}
	}
	return hex[idx:]
}

func HexToString(str interface{}) string {
	return common.HexToAddress(StringX{}.InterfaceToStr(str)).String()
}

func (StringX) GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}
