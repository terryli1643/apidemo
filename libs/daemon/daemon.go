package daemon

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var isDaemon = false

func IsDaemonMode() bool {
	return isDaemon
}

func init() {
	argc := len(os.Args)
	if argc == 1 {
		// redirectStderr()
	} else if argc >= 2 {
		if os.Args[1] == "--daemon=true" {
			fmt.Println(os.Args)
			os.Args[1] = "--daemon"
			cmd := exec.Command(os.Args[0], os.Args[1:]...)
			cmd.Start()
			fmt.Println("Server running in daemon . [PID]", cmd.Process.Pid)
			os.Exit(0)
		} else if os.Args[1] == "--daemon" {
			isDaemon = true
			redirectStderr()
			i := 1
			os.Args = append(os.Args[:i], os.Args[i+1:]...) // 恢复参数位置
			log.Println("Daemon Server Initializing...")
		} else {
		}
	} else {
		fmt.Println("The number of arguments incorrect .")
		os.Exit(-1)
	}
}
