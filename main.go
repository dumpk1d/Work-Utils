package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const (
	unknow   uint8 = 0
	ok       uint8 = 1
	warning  uint8 = 2
	critical uint8 = 3
)

type logstruct struct {
	Level    string    `json:"level"`
	Func     string    `json:"func"`
	Point    string    `json:"point"`
	Task     string    `json:"task"`
	Command  string    `json:"command"`
	IsSystem bool      `json:"is_system"`
	Vmid     int       `json:"vmid"`
	Time     time.Time `json:"time"`
	Message  string    `json:"message"`
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
	vms, status := GetAllVmsList()
	fmt.Println("VM'S:", vms, "\n", "status", status)
	GetBackupVmList()
}

func GetAllVmsList() (arg []string, status uint8) {

	var cmd = "virsh -c qemu:///system list --all | grep one | awk '{print $2}'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		return []string{err.Error()}, unknow
	} else {
		var output = string(out)
		if output == " " {
			return []string{"lol"}, ok
		} else {
			arr := strings.Split(output, "\n")
			return arr, ok
		}
	}
}

func GetBackupVmList() (arg []string, status uint8) {
	var cmd = "cat borg-agent.log | grep \"Created tasks for backup Node\""
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return []string{err.Error()}, unknow
	} else {
		println(string(out))
		return []string{" "}, ok
	}
}

func GetBaclkilstVmList() {

}
