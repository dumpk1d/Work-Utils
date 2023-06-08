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

type listvms struct {
	time string
	list []string
}

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
			return []string{"lol"}
		} else {
			arr := strings.Split(output, "\n")
			return arr
		}
	}
}

func GetBackupVmList() (arg []string) {
	var storeVar jstruct
	var cmd = "cat " + path + " | grep \"Created tasks for backup Node\" "
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		println("File doesn't exist	")
		os.Exit(unknow)
		return []string{" "}
	} else {
		json.Unmarshal([]byte(out), &storeVar)
		fmt.Println(storeVar)
		return []string{" "}
	}
}

func GetBaclkilstVmList() {

}

func NagiosResult(status int, errorCode uint8) {
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

func ErroCheck() {

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
	GetBackupVmList()
}
