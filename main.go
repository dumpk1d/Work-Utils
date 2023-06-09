package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	ok       int    = 0
	warning  int    = 1
	critical int    = 2
	unknow   int    = 3
	path     string = "borg-agent.log"
)

type jstruct struct {
	Level   string    `json:"level"`
	Time    time.Time `json:"time"`
	Message string    `json:"message"`
}

func GetAllVmsList() (arg []string) {

	var cmd = "virsh -c qemu:///system list --all | grep one | awk '{print $1}'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Println("Error qemu")
		os.Exit(unknow)
		return []string{" "}
	} else {
		var output = string(out)
		if output == " " {
			fmt.Println("No virtual machines")
			os.Exit(ok)
			return []string{" "}
		} else {
			arr := strings.Split(output, "\n")
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

//func GetVmList(array jstruct) []string {
//}

func NagiosResult(status int) {
	switch status {
	case ok:
		fmt.Printf("OK")
		os.Exit(ok)
	case warning:
		fmt.Printf("Warning")
		os.Exit(warning)
	case critical:
		fmt.Printf("Critical")
		os.Exit(critical)
	}
}

func main() {

	var (
		ftime int
		stime int
	)

	flag.IntVar(&ftime, "f", 0, "The first date")
	flag.IntVar(&stime, "s", 0, "The second date")
	flag.Parse()

	fmt.Println("First arg:", ftime, "\n", "Second arg", stime)
	vms := GetAllVmsList()
	fmt.Println("VM'S:", vms, "\n", "status")
	taskVms := GetParseJson("Created tasks for backup Node")
	fmt.Println(taskVms)
}
