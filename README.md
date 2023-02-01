# cvehound

A Go tool for searching the [CVE Project](https://github.com/CVEProject/cvelist) repository for a particular keyword.

## Requirements

- Go
- [ripgrep](https://github.com/BurntSushi/ripgrep)

## Usage
```go

go build
./cvehound keyword

```

## How it works
The cvehound script first checks if the cvelist repository exists in the user's home directory. If it doesn't, the repository is cloned. If it does, a git pull is performed to update the repository.

Once the repository is up-to-date, ripgrep is used to search for the keyword specified by the user. The results are displayed in the command line.

**Notes**
This script uses the `os` and `os/exec` packages to interact with the file system and run external commands, respectively. The filepath package is used to construct the path to the local repository.

Before using this script, make sure you have `ripgrep` (`rg`) installed on your system. You can then build and run the script by running `go build` followed by `./cvehound keyword`.

## Contributing
If you find a bug or have a feature request, please open an issue or a pull request on GitHub.
