package caffrey
import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func SliceToMap(slice []string) map[string]string{
	m := make(map[string]string)
	for k,v := range slice {
		k1 := strconv.Itoa(k)
		m[k1] = v
	}
	return m
}

func ToSlice(s interface{}) []bool {
	slice := make([]bool,0)
	ss := s.([]interface{})
	for _,val := range ss {
		slice = append(slice, val.(bool))
	}
	return slice
}

func BcjClient(inData []string) ([]bool,error){
	inMap := SliceToMap(inData)
	jsonStr,err := json.Marshal(inMap)
	if err != nil {
		fmt.Println("BcjClient.BcjClient() json.Marshal falied!")
		panic(err.Error())
	}
	url := "https://localhost:8077"
	req,err := http.NewRequest("POST",url,bytes.NewBuffer(jsonStr)) 
	if err != nil {
		fmt.Println("BcjClient.BcjClient() http.NewRequest falied!")
		panic(err.Error())
	}
	req.Header.Add("Content-Type","application/json")
	defer req.Body.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println("request fail")
		panic(err.Error())
	}
	defer resp.Body.Close()

	res,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("BcjClient.BcjClient() ioutil.ReadAll falied!")
		panic(err.Error())
	}
	var data map[string]interface{}
	json.Unmarshal(res,&data)
	// fmt.Printf("%T  %v",data["data"],data["data"])
	checkSlice := ToSlice(data["data"])
	return checkSlice,err

}