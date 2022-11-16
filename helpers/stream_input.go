package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
Streams input from a file
*/
func StreamInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not open file:\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Println(scanner.Text())
	if scanner.Scan() {

	}
	fmt.Println(scanner.Text())
}
