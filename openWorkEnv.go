// Да это кринж, и что ?
package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

const GOOS = runtime.GOOS

func createLinuxCfg() {

}

func main() {
	switch checkos := runtime.GOOS; checkos {
	case "linux":
		if _, err := os.Stat("wenv.json"); err == nil {
			fmt.Println("Ok")
		} else if errors.Is(err, os.ErrNotExist) {
			os.Create("wenv.json")
		}
	case "windows":
		fmt.Println("404 Error")
	default:
		fmt.Println("Error stop 00000000")
	}
}
