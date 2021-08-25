package main

import (
	"neal/caffrey"
	"fmt"
)

func main(){
	err := caffrey.AppMain()
	if err != nil {
		fmt.Println("service start fail!")
		return
	}
}