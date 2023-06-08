package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	ok       int    = 0
	warning  int    = 1
	critical int    = 2
	unknow   int    = 3
	path     string = "borg-agent.log"
)

func GetAllVmsList() (arg []string, status int) {

	var cmd = "virsh -c qemu:///system list --all | grep one | awk '{print $1}'"
	out, err := exec.Command("bash", "-c", cmd).Output()

	if err != nil {
		fmt.Println("Error")
		os.Exit(101)
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

func GetBackupVmList() (arg []string, status int) {
	var cmd = "cat " + path + " | grep \"Created tasks for backup Node\" "
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		os.Exit(102)
		return []string{err.Error()}, unknow
	} else {
		fmt.Println(string(out))
		return []string{" "}, ok
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
