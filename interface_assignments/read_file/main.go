package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	filename, err := getFileName(os.Args)

	if err != nil {
		printError(err)
		os.Exit(1)
	}

	file := handleOpenFile(filename)

	handlePrintFileContent(file)
}

func handleOpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		printError(err)
		os.Exit(2)
	}

	return file
}

func handlePrintFileContent(file *os.File) {
	nr, err := io.Copy(os.Stdout, file)
	if err != nil {
		printError(err)
		os.Exit(3)
	}

	fmt.Printf("\nRead %+v bytes", nr)
}

func printError(e error) {
	fmt.Println("Error: ", e)
}

func getFileName(argsList []string) (string, error) {

	if len(os.Args) < 2 {
		return "", errors.New("No filename was given")
	}

	return os.Args[1], nil
}
