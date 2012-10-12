package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func get_history() {
	cmd, err := exec.Command("ls").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	in := bufio.NewReader(string(cmd))

	s, ok := in.ReadString('\n')
	if ok != nil {
		fmt.Println(s)
	}

}
func main() {
	get_history()
}
