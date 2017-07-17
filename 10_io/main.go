package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadIn() {
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}

func WriteOut() {
	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()

	file.WriteString("test")
}

func ReadInDir() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}

func WalkDir() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, " : is Dir? ", info.IsDir())
		return nil
	})
}

func main() {
	WriteOut()
	ReadIn()

	ReadInDir()

	WalkDir()
}
