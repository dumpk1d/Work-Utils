package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

const (
	unknow   uint8 = 0
	ok       uint8 = 1
	warning  uint8 = 2
	critical uint8 = 3
)

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
}

func GetAllVmsList() (mass []string, status uint8) {

	var cmd = "virsh -c qemu:///system list --all | grep one | awk '{print $2}'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		return []string{" "}, unknow
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

func GetBackupVmList() {

}

func GetBaclkilstVmList() {

}
