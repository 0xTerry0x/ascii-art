package internal

import (
	"fmt"
)

func PrintNPErr(indexes []int) {
	if len(indexes) > 0 {

		fmt.Print("Non-Printable characters in position: \n")
		for i, v := range indexes {
			fmt.Print(v)

			if i != len(indexes)-1 {
				fmt.Print(", ")
			} else {
				fmt.Println()
			}

		}
	}
}