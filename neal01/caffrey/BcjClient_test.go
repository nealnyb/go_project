package caffrey

import (
	"fmt"
	"testing"
	"time"
)

func startService() {
	err := AppMain()
	if err != nil {
		fmt.Println("start service fail!")
		return
	}
}
func TestBcjClient(t *testing.T) {
	//启动service
	go startService()
	time.Sleep(time.Second)

	inData01 := []string{"a","b","c","d","e"}
	res01,err := BcjClient(inData01)
	expectRes01 := []bool{false,false,false,false,false}
	flag := 0
	for i := 0; i < 5;i++ {
		if res01[i] != expectRes01[i] || err != nil {
			flag++
		}
	}

	inData02 := []string{"a","b","c","g","h"}
	res02,err := BcjClient(inData02)
	expectRes02 := []bool{true,true,true,false,false}
	for i := 0; i < 5;i++ {
		if res02[i] != expectRes02[i] || err != nil {
			flag++
		}
	}

	if flag > 0 {
		t.Fatal("BcjClient() test failed!")
	}

}