package main

import (
	"fmt"
	"runtime"
)

func main () {
	fmt.Println("Go run Good")
	switch os := runtime.GOOS; os {
	case "darwin" :
	fmt.Println("osx")
	case "linux" :
	fmt.Println("Linux")
	default :
	fmt.Printf("%s",os)
	}
}
