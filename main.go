package main

import (
	"fmt"
	"os"

	"github.com/jmlattanzi/interp/repl"
)

func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("----\nStop right there, criminal scum! You have violated the law!")
	fmt.Println("Pay the court a fine or serve your sentence.")
	fmt.Println("Your stolen goods are now forfeit.")
	fmt.Println("----")
	fmt.Println()
	repl.Start(os.Stdin, os.Stdout)
}
