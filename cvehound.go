package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const cveListRepo = "https://github.com/CVEProject/cvelist.git"
const cveListLocalPath = "cvelist"

var home string
var cveListPath string

func initCVERepo() {
	if _, err := os.Stat(cveListPath); os.IsNotExist(err) {
		fmt.Println("Cloning cvelist repository...")
		err := exec.Command("git", "clone", cveListRepo, cveListPath).Run()
		if err != nil {
			fmt.Println("Error cloning repository:", err)
			os.Exit(1)
		}
	} else {
		fmt.Println("Updating cvelist repository...")
		err := exec.Command("git", "-C", cveListPath, "pull", "origin", "master").Run()
		if err != nil {
			fmt.Println("Error updating repository:", err)
			os.Exit(1)
		}
	}
}

func searchCVEList(searchTerm string) {
	fmt.Println("Searching for Search Term...")
	regex_string := fmt.Sprintf("(\"TITLE\".*%s\")", searchTerm)
	output, err := exec.Command("rg", regex_string, cveListPath).CombinedOutput()
	if err != nil {
		fmt.Println("Error searching for cvelist searchTerm:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}

func init() {
	home = os.Getenv("HOME")
	cveListPath = filepath.Join(home, cveListLocalPath)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: cvehound searchTerm")
		os.Exit(1)
	}

	searchTerm := os.Args[1]

	initCVERepo()
	searchCVEList(searchTerm)
}
