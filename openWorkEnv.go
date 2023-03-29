// Да это гига-кринж, и что ?
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func envOrDefault(key, def string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return def
	}

	return val
}

var (
	cfgUrl      = envOrDefault("CFG_URL", "https://raw.githubusercontent.com/dumpk1d/Work-Utils/main/wenv.yml")
	cfgFileName = envOrDefault("CFG_FILE_NAME", "wenv.yml")
)

// Скачивает файл конфига из репозитория
func downloadCfgFile() (err error) {
	// Создание файла
	out, err := os.Create(cfgFileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get
	resp, err := http.Get(cfgUrl)
	if err != nil {
		return err
	}
	defer out.Close()

	// Проверка ответа
	if resp.StatusCode != http.StatusOK {
		os.Remove(cfgFileName)
		return fmt.Errorf("Status: %s", resp.Status)
	}

	// Запись
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// initCfg() { //парсинг ? конфига ( тут начинается мясо))))) )

func main() {
	_, err := os.Stat(cfgFileName)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalln(err)
		}

		err := downloadCfgFile()
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println(string("\033[32m"), "Ok")
}
