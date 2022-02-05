package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	must := func(err error) {
		if err != nil {
			log.Fatalln(err)
		}
	}

	switch os.Args[1] {
	case "run":
		cmd := exec.Command("/proc/self/exe", append([]string{"spawn"}, os.Args[2:]...)...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS}
		must(cmd.Run())

	case "spawn":
		cmd := exec.Command(os.Args[2], os.Args[3:]...)
		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		must(syscall.Chroot("/home/root"))
		must(os.Chdir("/"))
		must(syscall.Mount("proc", "proc", "proc", 0, ""))
		must(cmd.Run())

	default:
		log.Fatalln("not supported command")
	}
}
