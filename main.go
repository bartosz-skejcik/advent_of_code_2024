package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func scanDir(ignoredFolders []string) ([]string, error) {
	// scan the current working directory and return all the folders
	var folders []string
	root := "."
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		isIgnored := false
		for _, ignoredFolder := range ignoredFolders {
			if strings.Contains(path, ignoredFolder) || path == "." {
				isIgnored = true
				break
			}
		}
		if info.IsDir() && !isIgnored {
			folders = append(folders, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return folders, nil
}

func displayFolders() []string {
	ignoredFolders := []string{".git", ".idea", "internal"}
	// Assume scanDir scans the directory and filters out ignored folders.
	folders, err := scanDir(ignoredFolders)
	if err != nil {
		fmt.Println("Error scanning directories:", err)
		return nil
	}
	var dayFolders []string
	for _, folder := range folders {
		if strings.Contains(folder, "day") {
			// Extract the numeric part from the folder name
			parts := strings.Split(folder, "_")
			if len(parts) > 1 {
				fmt.Printf("%s. %s\n", parts[1], folder)
				dayFolders = append(dayFolders, folder)
			} else {
				fmt.Printf("?. %s\n", folder) // Fallback for unexpected cases
				dayFolders = append(dayFolders, folder)
			}
		} else {
			// Display non-"day" folders with a placeholder or custom logic
			fmt.Printf("?. %s\n", folder)
		}
	}
	return dayFolders
}

func getDayFolder(day int, folders []string) string {
	for _, folder := range folders {
		// Check if the folder contains the day number
		if strings.Contains(folder, fmt.Sprintf("day_%d", day)) ||
			strings.Contains(folder, fmt.Sprintf("day%d", day)) {
			return folder
		}
	}
	return ""
}

func runDay(day int) {
	// Display folders to help user identify correct day folder
	folders := displayFolders()

	// Get the folder for the day
	folder := getDayFolder(day, folders)

	if folder == "" {
		fmt.Printf("No folder found for day %d\n", day)
		return
	}

	// Potential implementation could include:
	err := os.Chdir(folder)
	if err != nil {
		log.Fatalf("Error changing directory: %v", err)
	}

	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Error running command: %v", err)
	}

	err = os.Chdir("..")
}

func main() {
	// Print a welcome message
	fmt.Printf("Welcome to Advent of Code 2024\n\n")

	// Check if a day is provided as a command-line argument
	var day int
	var err error

	if len(os.Args) > 1 {
		// Try to parse the first command-line argument as an integer
		day, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Invalid day provided. Please enter a valid day number.")
			return
		}
		runDay(day)
	} else {
		// Display folders and prompt for input
		displayFolders()

		fmt.Println("\nPlease enter the day you want to run:")
		_, err := fmt.Scan(&day)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		runDay(day)
	}
}
