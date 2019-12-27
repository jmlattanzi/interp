package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jmlattanzi/interp/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Stop right there criminal scum!")
	fmt.Println("sup, " + user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
