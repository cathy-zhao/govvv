package main

import "fmt"

var (
	// These fields are populated by govvv
	Version    = "untouched"
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
)

func main() {
	fmt.Printf("Version=%s\n", Version)
	fmt.Printf("BuildDate=%s\n", BuildDate)
	fmt.Printf("GitCommit=%s\n", GitCommit)
	fmt.Printf("GitBranch=%s\n", GitBranch)
	fmt.Printf("GitState=%s\n", GitState)
	fmt.Printf("GitSummary=%s\n", GitSummary)
}
