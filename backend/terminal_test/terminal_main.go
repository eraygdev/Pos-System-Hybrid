package terminal_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var rdr = bufio.NewReader(os.Stdin)

func InitCmd() {
Loop:
	for {
		fmt.Println("\n" + strings.Repeat("=", 45))
		fmt.Println("          🚀 POS SYSTEM - MAIN MENU          ")
		fmt.Println(strings.Repeat("=", 45))
		selection := AskQ(1)

		switch selection {
		case -1:
			fmt.Print("\n⚠️  Changes will not be saved! ")
			if y_or_n() {
				break Loop
			}
		case 0:
			leaveSafe()
		case 1:
			if err := seeDetails(); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("An error occurred.")
		}

		fmt.Println("\n" + strings.Repeat("-", 45))
		fmt.Println("👉 Press [ENTER] to return to Main Menu")
		fmt.Println(strings.Repeat("-", 45))
		_, _ = rdr.ReadString('\n')
	}
	fmt.Println("Session ended.")
}
