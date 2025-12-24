package ascii

import (
	"fmt"
)

func PrintNPErr(indexes []int) {
	if len(indexes) > 0 {
		fmt.Print("Non-Printable characters in position: \n")
		// For each index, print the index
		for i, v := range indexes {
			fmt.Print(v)

			// If the index is not the last index, print a comma
			if i != len(indexes)-1 {
				fmt.Print(", ")
				// If the index is the last index, print a newline
			} else {
				fmt.Print("\n")
			}
		}
	}
}
