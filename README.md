<h1 align="center">
    <img width="194" src=".github/cvehound-logo.svg#gh-light-mode-only" alt="cvehound">
    <img width="194" src=".github/cvehound-logo.svg#gh-dark-mode-only" alt="cvehound">
</h1>

# CVEHound

A Go tool for searching the [CVE Project](https://github.com/CVEProject/cvelist) repository for a particular keyword.

## Requirements

- Go
- Git
- [ripgrep](https://github.com/BurntSushi/ripgrep)

## Usage

```go

go build
./cvehound keyword

```

## How it works

- The cvehound script first checks if the cvelist repository exists in the user's home directory.
- If the repository doesn't exists, the repository is cloned.
- If the repository does exists, then a git pull is performed to update the repository.

Once the repository is up-to-date, ripgrep is used to search for the keyword specified by the user. The results are displayed in the command line.

**Notes**
This script uses the `os` and `os/exec` packages to interact with the file system and run external commands, respectively. The filepath package is used to construct the path to the local repository.

Before using this script, make sure you have `ripgrep` (`rg`) installed on your system. You can then build and run the script by running `go build` followed by `./cvehound keyword`.

## Contributing

If you find a bug or have a feature request, please open an issue or a pull request on GitHub.
