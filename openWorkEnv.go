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

const GOOS = runtime.GOOS

func downloadCfgFile() (err error) {
	//Create
	out, err := os.Create("wenv.json")
	if err != nil {
		return err
	}
	defer out.Close()

	//Get
	resp, err := http.Get("https://raw.githubusercontent.com/dumpk1d/Work-Utils/main/wenv.json")
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
		if _, err := os.Stat("wenv.json"); err == nil {
			fmt.Println("Ok")
		} else if errors.Is(err, os.ErrNotExist) {
			fmt.Println(downloadCfgFile())
		}
	case "windows":
		fmt.Println("404 Error")
	default:
		fmt.Println("Error stop 00000000")
	}
}
