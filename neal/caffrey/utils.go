package caffrey

import (
	"sort"
)

//检测切片中是否存在某个元素
func CheckExist(str string,slice []string) bool {
	for _,v := range slice {
		if v == str {
			return true
		}	
	}
	return false
}

//将map转换为[]string
func MapToSlice(m map[string]interface{}) []string {
	slice := make([]string,0)
	sortSlice := make([]string,0)
	for k, _ := range m {
		sortSlice = append(sortSlice, k)
	}
	/*对map的键进行排序 保证 接收到的json反序列化为map是不会乱序，
	以保证返回的[]bool值一一对应输入的[]string
	*/
	sort.Strings(sortSlice)
	for _,val := range sortSlice {
		slice = append(slice, m[val].(string))
	}
	return slice
}