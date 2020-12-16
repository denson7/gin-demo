package utils

import (
	"strconv"
	"strings"
)

func BuildJson(kvs map[string]string) string{
	jsonStr := "{"
	for k,v := range kvs{
		jsonStr += "\""+k+"\":"+"\""+v+"\","
	}
	jsonStr = strings.TrimRight(jsonStr,",")
	jsonStr += "}"
	return jsonStr
}

func BuildIntValueJson(kvs map[string]int) string{
	jsonStr := "{"
	for k,v := range kvs{
		jsonStr += "\""+k+"\":"+strconv.Itoa(v)+","
	}
	jsonStr = strings.TrimRight(jsonStr,",")
	jsonStr += "}"
	return jsonStr
}