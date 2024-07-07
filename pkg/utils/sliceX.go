package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func SliceInt642String(data []int64, split string) string {
	var res string
	for i := 0; i < len(data); i++ {
		res = fmt.Sprintf("%s%s%d", res, split, data[i])
	}

	return res
}

func String2SliceInt64(data string, split string) []int64 {
	var res []int64
	splitStr := strings.Split(data, split)
	for i := 0; i < len(splitStr); i++ {
		item, err := strconv.ParseInt(splitStr[i], 10, 64)
		if err != nil {
			continue
		}

		res = append(res, item)
	}

	return res
}

func SliceStringHash256(str ...string) string {
	strJoin := strings.Join(str, "_")

	hasher := sha256.New()
	hasher.Write([]byte(strJoin))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
