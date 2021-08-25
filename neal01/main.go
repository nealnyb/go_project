package main
import (
	"neal01/caffrey"
	"fmt"
)

func main(){
	err := caffrey.AppMain()
	if err != nil {
		fmt.Println("service start fail!")
		return
	}
}