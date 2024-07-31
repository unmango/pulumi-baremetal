package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/Masterminds/semver/v3"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	out, err := exec.Command("pulumictl", "get", "version").CombinedOutput()
	if err != nil {
		panic(err)
	}

	version := strings.TrimSpace(string(out))
	ver, err := semver.NewVersion(version)
	if err != nil {
		panic(err)
	}

	next := ver.IncPatch().String()
	tag := fmt.Sprintf("v%s", next)
	sdkTag := fmt.Sprintf("sdk/v%s", next)

	fmt.Println("About to create tags:")
	fmt.Println(tag)
	fmt.Println(sdkTag)
	fmt.Println()

	fmt.Println("Are you sure?")
	var response string
	_, err = fmt.Scanln(&response)
	if err != nil {
		return fmt.Errorf("user rejected: %w", err)
	}

	if strings.ToUpper(response) != "YES" {
		panic("Aborting")
	}

	createTag(tag)
	createTag(sdkTag)

	return nil
}

func createTag(tag string) error {
	out, err := exec.Command("git", "tag", tag).CombinedOutput()
	fmt.Print(string(out))
	if err != nil {
		return fmt.Errorf("git tag: %w", err)
	}

	out, err = exec.Command("git", "push", "origin", tag).CombinedOutput()
	fmt.Print(string(out))
	if err != nil {
		return fmt.Errorf("git tag push: %w", err)
	}

	return nil
}
