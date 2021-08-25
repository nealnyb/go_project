package caffrey

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var safeData *SafeData = &SafeData{}

func hander(w http.ResponseWriter,r *http.Request) {
	//接收客户端请求json数据并反序列化
	var data map[string]interface{}
	res,_ := ioutil.ReadAll(r.Body)
	json.Unmarshal(res,&data)
	//将序列化后的map转换为切片类型
	inSlice := MapToSlice(data)
	//将输入的切片和存储的数据进行对比，存在返回true,不存在返回false,返回一个[]bool
	srcSlice := safeData.Get()
	checkSlice := make([]bool,0,len(inSlice))
	for _,inVal := range inSlice {
		check := CheckExist(inVal,srcSlice)
		checkSlice = append(checkSlice, check)
	}
	//将输入的切片数据增加到存储数据中
	safeData.Add(inSlice)
	// checkSlice := InterviewServer(inSlice,safeData)
	//向客户端返回一个json的数据
	w.Header().Set("Content-Type","application/json")
	post := &ResponeData{
		Data: checkSlice,
	}
	jsonData,_ := json.Marshal(post)
	w.Write(jsonData)
	// fmt.Printf("%T %v\n",checkSlice,checkSlice)
}

func AppMain() error {
	http.HandleFunc("/",hander)
	// http.ListenAndServe("localhost:8099",nil)
	err := http.ListenAndServeTLS("localhost:8077","cert.pem","key.pem",nil)
	return err
}