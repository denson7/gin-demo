package utils

import "reflect"


//该方法支持所有类型的判断，但是严重影响性能
/**
eg:
intArr := []int{1,2,3,4,5}
InArray(1,intArr)
strArr := []string{"aa","bb","cc","dd"}
InArray("aa",strArr)
 */
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				return true
			}
		}
	}
	return false
}
