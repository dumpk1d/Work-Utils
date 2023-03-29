// Да это гига-кринж, и что ?
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

const cfgUrl string = "https://raw.githubusercontent.com/dumpk1d/Work-Utils/main/wenv.?"
const cfgFileName string = "wenv.yml"

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

//func readCfg() { //парсинг ? конфига ( тут начинается мясо))))) )

func main() {

	switch checkos := runtime.GOOS; checkos {
	case "linux":
		if _, err := os.Stat(cfgFileName); err == nil { // Если файл на месте
			fmt.Println(string("\033[32m"), "Ok")
			fmt.Println()
		} else if errors.Is(err, os.ErrNotExist) { // Если файл не нужен
			if err = downloadCfgFile(); err == nil { // Проверка есть ли ошибка
				fmt.Println(string("\033[32m"), "Done")
			} else {
				fmt.Println(string("\033[31m"), err)
			}
		}

	case "windows":
		fmt.Println("Nope:)")

	default:
		fmt.Println("Error stop 00000000")
	}
}
