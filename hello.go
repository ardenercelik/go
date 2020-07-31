package main

import (
	"fmt"
	"os"
	_ "strings"
	"time"
)

func echo() {
	//s := ""
	for index, element := range os.Args[1:] {
		fmt.Println(index, element)
	}

}

func main() {
	start := time.Now()
	// a := strings.Join(os.Args[1:], " ")
	// fmt.Println(a)
	echo()
	elapsed := time.Since(start)
	fmt.Println("%s saniye geçti", elapsed)
}
