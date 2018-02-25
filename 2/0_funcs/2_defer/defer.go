package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	cnt, err := ReadFile("test_file")
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("Content of test_file:\n%s", string(cnt))

}

func ReadFile(f string) ([]byte, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, errors.New("Can not open file")
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {

		return nil, errors.New("Can not get file stat")
	}
	bytes := make([]byte, stat.Size(), 20)
	_, err = file.Read(bytes)
	if err != nil {
		return nil, errors.New("Can not read from file")
	}
	return bytes, nil
}
