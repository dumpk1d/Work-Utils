package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	ok       uint8  = 0
	warning  uint8  = 1
	critical uint8  = 2
	unknow   uint8  = 3
	path     string = "borg-agent.log"
)

func GetAllVmsList() (arg []string, status uint8) {

	var cmd = "virsh -c qemu:///system list --all | grep one | awk '{print $2}'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Println("Error")
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
	var cmd = "cat " + path + " | grep \"Created tasks for backup Node\" | tr -d '.}\"' | sed -e 's|.*VMs:||'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {

		return []string{err.Error()}, unknow
	} else {
		fmt.Println(string(out))
		return []string{" "}, ok
	}
}

func GetBaclkilstVmList() {

}

func NagiosResult(status uint8) {
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
	case	
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
