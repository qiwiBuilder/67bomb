package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	var acceptbool bool
	fmt.Println("WARNING! This bomb is very simple, but dangerous. If you agree to run this bomber type True if you agree.")
	fmt.Println("If you don't agree quit the executive RIGHT NOW.")
	fmt.Println("Note: made by PechenieSMushyakom")
	fmt.Scan(&acceptbool)
	if acceptbool {
		for i := 0; i < 100; i++ {
			fmt.Println("67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 67 ")
			open("https://youtu.be/v0NDDoNRtQ8?t=24")
			open("https://youtu.be/jq4femgTxBE?si=Asr24nfNlWO9Ze9G")
			time.Sleep(67 * time.Millisecond)
		}
	}
}
