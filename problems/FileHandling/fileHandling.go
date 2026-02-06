package main

import (
	"errors"
	"fmt"
	"os"
)

func fileExists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func createFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func writeFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	file := "sample.txt"
	exists, _ := fileExists(file)
	fmt.Println("File exists? : ", exists)

	if !exists {
		createFile(file)
		fmt.Println("File created successfully..")
	}

	fmt.Println("Write File", writeFile(file, []byte("Hello this is my very first file being created..")))

	data, _ := readFile(file)
	fmt.Println("Read File: ", string(data))
}
