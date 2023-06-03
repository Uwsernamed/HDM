package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var dictionary string
var currentDir string

func main() {
	// Welcome!
	fmt.Println("HDF Manager\nversion: 1.0.0\n ")

	// Getting the current working directory
	var err error
	currentDir, err = os.Getwd()
	if err != nil {
		fmt.Println("Failed to retrieve the current working directory:", err)
		return
	}

	// Time to read the user's command then send it
	var command string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("HDM %s | ", currentDir)
		if !scanner.Scan() {
			err = scanner.Err()
			if err != nil {
				fmt.Println("Failed to read input:", err)
			}
			break
		}
		command = scanner.Text()
		if command == "exit" {
			break
		}
		parsedp := strings.Split(strings.TrimSpace(command), " ")

		interpreter := &Interpreter{
			parsed: parsedp,
		}
		switch len(interpreter.parsed) {
		case 1:
			(interpreter).One()
		case 2:
			(interpreter).Two()
		default:
			fmt.Println("No function found for the given word count.\n ")
		}

	}
}

type Interpreter struct {
	parsed []string
}

func (r *Interpreter) One() {
	switch r.parsed[0] {
	case "help":
		O_help()
	case "create":
		O_create()
	default:
		fmt.Println("Unknown command. Type 'help' for a list of available commands.\n ")
	}
}

func (r *Interpreter) Two() {
	switch r.parsed[0] {
	case "cd":
		T_cd(r.parsed[1])
	case "changedirectory":
		T_cd(r.parsed[1])
	case "open":
		T_open(r.parsed[1])
	default:
		fmt.Println("Unknown command. Type 'help' for a list of available commands.\n ")
	}
}

func O_help() {
	fmt.Println("Available commands:")
	fmt.Println("- help:   display a list of available commands.")
	fmt.Println("- exit:   exit the HDF Manager.")
	fmt.Println("- create: creates a folder hierarchy map, .HDF.\n ")

	fmt.Println("- cd              [full-directory]: changes the directory.")
	fmt.Println("- changedirectory [full-directory]: changes the directory.")
	fmt.Println("- open            [text-file-path]: displays a text file, intended for folder hierarchy maps.\n ")
	// Add more commands and descriptions as needed
}

func O_create() {
	// Specify the directory path
	directory := currentDir // Replace with your desired directory path

	// Get the base folder name
	baseFolder := filepath.Base(directory)

	// Get all subfolders within the directory
	subfolders, err := getSubfolders(directory)
	if err != nil {
		log.Fatalf("Failed to get subfolders: %s", err)
	}

	// Format the subfolders as required
	text := formatSubfolders(subfolders)

	// Create the file
	filename := fmt.Sprintf("%s/%s.HDF", directory, baseFolder)
	err = os.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}

	fmt.Printf("File '%s' created successfully.\n\n", filename)
}

func T_cd(T_arg string) {
	err := os.Chdir(T_arg)
	if err != nil {
		fmt.Println("Failed to change directory.\n ")
	} else {
		fmt.Println("Successfully changed directory.\n ")
		currentDir = T_arg
	}
}

func T_open(T_arg string) {
	if hdfOpen(T_arg) == 1 {
		dictionary = T_arg
		file, err := os.Open(dictionary)
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close() // Make sure to close the file when done

		// Read the file content
		content, err := os.ReadFile(file.Name())
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		// Print the content
		fmt.Printf("%s\n\n", string(content))
	} else {
		fmt.Println("unable to select, in order to open.\n ")
	}
}

func hdfOpen(finderArg string) int {

	_, err := os.Stat(finderArg)
	if err != nil {
		if os.IsNotExist(err) {

		} else {
			fmt.Println("Error:", err)
		}
	} else {
		return 1
	}
	return 0
}

func getSubfolders(directory string) ([]string, error) {
	var subfolders []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != directory {
			subfolders = append(subfolders, path)
		}
		return nil
	})

	return subfolders, err
}

func formatSubfolders(subfolders []string) string {
	var builder strings.Builder
	var fsfarr []int

	for _, afolder := range subfolders {
		fsfarr = append(fsfarr, strings.Count(afolder, string(filepath.Separator)))
	}
	fsfleastNumber := findLeastNumber(fsfarr)

	for _, folder := range subfolders {
		indent := strings.Repeat("    ", strings.Count(folder, string(filepath.Separator))-fsfleastNumber)
		builder.WriteString(fmt.Sprintf("%s[%s]\n", indent, filepath.Base(folder)))
	}

	return builder.String()
}

func findLeastNumber(arr []int) int {
	if len(arr) == 0 {
		return -1 // Return -1 if the array is empty
	}

	minNumber := arr[0] // Assume the first element is the minimum

	for _, num := range arr {
		if num < minNumber {
			minNumber = num
		}
	}

	return minNumber
}
