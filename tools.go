package go_vesdk

import (
	"github.com/go-dragon/util"
	"strconv"
)

func toJson(data interface{}) string {
	js, _ := util.FastJson.Marshal(data)
	return string(js)
}

func intSliceToStringSlice(intS []int) []string {
	length := len(intS)
	strSlice := make([]string, length)
	for i := 0; i < length; i++ {
		strSlice = append(strSlice, strconv.FormatInt(int64(intS[i]), 10))
	}
	return strSlice
}
