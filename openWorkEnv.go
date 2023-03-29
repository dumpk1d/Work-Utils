// Да это гига-кринж, и что ?(Да-да я клоун, ещё не прошёл gotour до конца)
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
	//Создание файла
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

	//Проверка ответа
	if resp.StatusCode != http.StatusOK {
		os.Remove(cfgFileName)
		return fmt.Errorf("Status: %s", resp.Status)
	}

	//Запись
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// func readCfg() (err error) {

func main() {

	switch checkos := runtime.GOOS; checkos {
	case "linux":
		if _, err := os.Stat(cfgFileName); err == nil {
			fmt.Println("Ok")
		} else if errors.Is(err, os.ErrNotExist) {
			if err := downloadCfgFile(); err == nil {
				fmt.Println("Done")
			} else {
				fmt.Println(err)
			}
		}

	case "windows":
		fmt.Println("Nope:)")

	default:
		fmt.Println("Error stop 00000000")
	}
}
