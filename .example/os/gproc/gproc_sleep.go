package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gproc"
)

func main() {
	fmt.Println(gproc.Pid())
	err := gproc.ShellRun("sleep 99999s")
	fmt.Println(err)
}
