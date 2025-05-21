package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	notesPath = "notes-app"
)

func main() {
	args := os.Args[1:]

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	dirPath := filepath.Join(homeDir, notesPath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}

	if len(args) == 0 {
		listAllFiles(dirPath)
		return
	}

	if len(args) == 1 {
		openFile(dirPath, args[0])
	} else {
		fmt.Println("Too many arguments. Usage:")
		fmt.Println("   notes            # list of notes")
		fmt.Println("   notes <filename> # open/create note")
		os.Exit(1)
	}
}

func listAllFiles(notesDir string) {
	files, err := os.ReadDir(notesDir)
	if err != nil {
		fmt.Println("Unable to read notes in directory:", err)
		os.Exit(1)
	}

	fmt.Println("-- Notes --")
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println("*", file.Name())
		}
	}
}

func openFile(notesDir, fileName string) {
	// Set markdown as default extension
	extension := filepath.Ext(fileName)
	if extension == "" {
		fileName += ".md"
	}

	notePath := filepath.Join(notesDir, fileName)

	if _, err := os.Stat(notePath); os.IsNotExist(err) {
		file, err := os.Create(notePath)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			os.Exit(1)
		}

		file.Close()
		fmt.Println("Created new note:", notePath)
	} else {
		fmt.Println("Opening existing note:", notePath)
	}

	// Open file in neovim
	cmd := exec.Command("nvim", notePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Failed to open file in Neovim:", err)
		os.Exit(1)
	}
}
