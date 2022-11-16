package helpers

import (
	"log"
	"os"
	"testing"
)

var _ = (func() interface{} {
	_testing = true
	return nil
}())

func TestReadAllInt(t *testing.T) {
	input := ReadAllInputInt("/tmp/testfile")
	inputTest := []int{1011, 1022, 1033, 1044, 1055}
	for i := range input {
		if input[i] != inputTest[i] {
			t.Error("Input did not match expected", input, inputTest)
			break
		}
	}
}

func TestReadAllString(t *testing.T) {

}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	cleanup()
}

func setup() {
	inputFile := `1011
1022
1033
1044
1055
`

	file, err := os.Create("/tmp/testfile")
	if err != nil {
		log.Fatal("Error during test setup", err)
	}
	defer file.Close()
	file.Write([]byte(inputFile))

}

func cleanup() {
	os.Remove("/tmp/testfile")
}
