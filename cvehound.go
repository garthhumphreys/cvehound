package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const repo = "https://github.com/CVEProject/cvelist.git"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cvehound keyword")
		os.Exit(1)
	}

	keyword := os.Args[1]

	home := os.Getenv("HOME")
	repoPath := filepath.Join(home, "cvelist")

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		fmt.Println("Cloning cvelist repository...")
		err := exec.Command("git", "clone", repo, repoPath).Run()
		if err != nil {
			fmt.Println("Error cloning repository:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Updating cvelist repository...")
		err := exec.Command("git", "-C", repoPath, "pull", "origin", "master").Run()
		if err != nil {
			fmt.Println("Error updating repository:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Searching for keyword...")
	regex_string := fmt.Sprintf("(\"TITLE\".*%s\")", keyword)
	output, err := exec.Command("rg", regex_string, repoPath).CombinedOutput()
	if err != nil {
		fmt.Println("Error searching for keyword:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
