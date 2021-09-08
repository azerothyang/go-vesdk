package go_vesdk

import "github.com/go-dragon/util"

func toJson(data interface{}) string {
	js, _ := util.FastJson.Marshal(data)
	return string(js)
}
