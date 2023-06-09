package main

// Да простит меня бог
import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	OK       int    = 0
	warning  int    = 1
	critical int    = 2
	unknow   int    = 3
	path     string = "file.log"
)

// Структура для хранения распаршенной JSON
type jstruct struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

// Пока не понял как и зачем
//func CheckTime(stime string,ftime string,arg jstruct) (ArgsNum int){

func GetAllVmsList() (arg []string) {

	var cmd = "virsh -c qemu:///system list --all | grep one | awk '{print $2}' | tr -d one- | tr -d \" \""
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("Error qemu")
		os.Exit(unknow)
		return []string{" "}
	} else {
		if string(out) == " " {
			//НЕ ЗАБЫТЬ УТОЧНИТЬ код возврата
			fmt.Println("No virtual machines")
			os.Exit(OK)
			return []string{" "}
		} else {
			arr := strings.Split(string(out), "\n") // *
			return arr
		}
	}
}

func GetParseJson(str string) (arg []jstruct) {
	var (
		storeVar []jstruct
		tmpVar   jstruct
	)
	var cmd = "cat " + path + " | grep \"" + str + "\" "
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("File doesn't exist")
		os.Exit(unknow)
		return storeVar
	} else {
		arr := strings.Split(string(out), "\n")
		for _, b := range arr {
			if b != "" {
				json.Unmarshal([]byte(b), &tmpVar)
				storeVar = append(storeVar, tmpVar)
			}
		}
		return storeVar
	}
}

func GetVmList(array jstruct) []string {
	// Да-да, я тут подустал. Но какой же красивый костыль!
	var cmd = "echo " + string(array.Message) + "| sed 's|.*:||'" + "| tr -d \".\n \""
	out, _ := exec.Command("bash", "-c", cmd).Output()
	if string(out) == " \n " {
		return []string{" "}
	} else {
		arr := strings.Split(string(out), ",")
		return arr
	}
}

func GetArrayDiffs(blacklist []string, taskVms []string, AllVmsList []string) []string {
	var (
		diff_list   []string
		result_list []string
	)
	if blacklist[0] == "" && taskVms[0] == "" {
		//НЕ ЗАБЫТЬ УТОЧНИТЬ код возврата
		fmt.Println("Tasks doesn't exist")
		os.Exit(critical)
		return []string{" "}
	} else {
		for _, b := range blacklist {
			diff_list = append(diff_list, b)
		}
		for _, b := range taskVms {
			diff_list = append(diff_list, b)
		}
		// Магия которая сравнивает 2 массива
		m := make(map[string]bool)
		for _, item := range AllVmsList {
			m[item] = true
		}
		for _, item := range diff_list {
			if _, ok := m[item]; !ok {
				result_list = append(result_list, item)
			}
		}
		return result_list
	}
}

func main() {
	// До лучших времён :D
	//var (
	//      ftime       string
	//      stime       string
	//)
	//flag.StringVar(&ftime, "f", "23", "The first date")
	//flag.StringVar(&stime, "s", "09", "The second date")
	//flag.Parse()
	yesterday := time.Now().AddDate(0, 0, -1)
	taskVms := GetParseJson("Created tasks for backup Node")
	blacklistVms := GetParseJson("Blacklisted vm")

	if taskVms[0].Time.Day() == yesterday.Day() {
		var allvms = GetAllVmsList()
		var vmlist = GetVmList(taskVms[0])
		var blackvm = GetVmList(blacklistVms[0])
		var resultArg = GetArrayDiffs(blackvm, vmlist, allvms)
		if len(resultArg) != 0 {
			for _, b := range resultArg {
				fmt.Println(b, "- No Job")
			}
			os.Exit(critical)
		} else {
			fmt.Println("Ok")
			os.Exit(OK)
		}
	} else {
		fmt.Println("Log line for", yesterday.Day(), yesterday.Month(), "not found")
		os.Exit(3)
	}
}
