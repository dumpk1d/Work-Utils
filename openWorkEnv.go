// Да это кринж, и что ?
package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

// определение ОС
const GOOS string = runtime.GOOS
const cfgUrl string = "https://raw.githubusercontent.com/dumpk1d/Work-Utils/main/wenv.json"
const cfgFileName string = "wenv.json"

// Скачивает файл конфига из репозитория
func downloadCfgFile() (err error) {
	//Create
	out, err := os.Create(cfgFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	//Get
	resp, err := http.Get(cfgUrl)
	if err != nil {
		return err
	}
	defer out.Close()

	//Check
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status: %s", resp.Status)
	}

	//Write
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	switch checkos := runtime.GOOS; checkos {
	case "linux":
		if _, err := os.Stat(cfgFileName); err == nil {
			fmt.Println("Ok")
		} else if errors.Is(err, os.ErrNotExist) {
			var code = downloadCfgFile()
			if code == nil {
				fmt.Println("Ok")
			} else {
				os.Remove(cfgFileName)
				fmt.Println(code)
			}
		}
	case "windows":
		fmt.Println("404 Error")
	default:
		fmt.Println("Error stop 00000000")
	}
}
