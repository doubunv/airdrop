package utils

import (
	"math/rand"
	"time"
)

func RandStr(n int) string {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetDayTimeString() (startString string, endString string) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local)
	startString = startOfDay.Format("2006-01-02 15:04:05")
	endString = endOfDay.Format("2006-01-02 15:04:05")
	return
}

func GetDayTimeInt() (startTimestamp int64, endTimestamp int64) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	endOfDay := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.Local)
	startTimestamp = startOfDay.Unix()
	endTimestamp = endOfDay.Unix()
	return
}

func DifferenceStringArr(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	for _, s := range slice1 {
		m[s] = true
	}
	for _, s := range slice2 {
		delete(m, s)
	}
	result := make([]string, 0, len(m))
	for s := range m {
		result = append(result, s)
	}
	return result
}

func TimeStrToInt(timeStr string) int64 {
	if timeStr == "" {
		return int64(0)
	}
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return timeObj.Unix()
}

func DifferenceIntArr(int1, int2 []int64) []int64 {
	m := make(map[int64]bool)
	for _, s := range int1 {
		m[s] = true
	}
	for _, s := range int2 {
		delete(m, s)
	}
	result := make([]int64, 0, len(m))
	for s := range m {
		result = append(result, s)
	}
	return result
}

func GetBeforeDayTimeStart(days int) int64 {
	now := time.Now()
	sevenDaysAgo := now.AddDate(0, 0, -days)
	sevenDaysAgoMidnight := time.Date(sevenDaysAgo.Year(), sevenDaysAgo.Month(), sevenDaysAgo.Day(), 0, 0, 0, 0, sevenDaysAgo.Location())
	return sevenDaysAgoMidnight.Unix()
}

func GetBeforeDayTimeEnd(days int) int64 {
	now := time.Now()
	sevenDaysAgo := now.AddDate(0, 0, -days)
	sevenDaysAgoMidnight := time.Date(sevenDaysAgo.Year(), sevenDaysAgo.Month(), sevenDaysAgo.Day(), 23, 59, 59, 0, sevenDaysAgo.Location())
	return sevenDaysAgoMidnight.Unix()
}

func TimeIntToStr(timeInt int64) string {
	if timeInt == 0 {
		return ""
	}
	return time.Unix(timeInt, 0).Format("2006-01-02")
}

func TimeIntToStrSce(timeInt int64) string {
	if timeInt == 0 {
		return ""
	}
	return time.Unix(timeInt, 0).Format("2006-01-02 15:04:05")
}
