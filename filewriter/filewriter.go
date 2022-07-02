package fileWriter

import (
	"fmt"
	"os"
)

func fileWriter(message string) {

	f, err := os.Create("./hello.txt")
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(message)
	if err != nil {
		fmt.Printf("failed to write message to file \n: %v", err)
		return
	}

	fmt.Println("Done!")
}
