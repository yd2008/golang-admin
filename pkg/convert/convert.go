package convert

import (
	"encoding/json"
	"strconv"
)

func Str2uint(s string) uint {
	int, _ := strconv.Atoi(s)
	return uint(int)
}

// 把结构体转换为map，同时key改为下划线形式
// 结构体json标签要为下划线形式
func Struct2UnderlineMap(s interface{}) (map[string]interface{}, error) {
	marshal, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(marshal), &mapResult)
	if err != nil {
		return nil, err
	}
	return mapResult, nil
}
