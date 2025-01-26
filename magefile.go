//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"
	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)


// Pull the latest image for the AudioShelf Server
func DockerPull() error {
	fmt.Println("Pulling latest Audioshelf image...")
	cmd := exec.Command("docker", "compose", "pull")
	return cmd.Run()
}

// Start the AudioShelf Server
func StartServer() error {
	fmt.Println("Starting Audioshelf server...")
	cmd := exec.Command("docker", "compose", "up", "-d")
	return cmd.Run()
}

// Stop the AudioShelf Server
func StopServer() error {
	fmt.Println("Stopping Audioshelf server...")
	cmd := exec.Command("docker", "compose", "down")
	return cmd.Run()
}

// Update the AudioShelf Server to the latest image
func UpdateServer() error {
	// Stop the server
	mg.Deps(StopServer)

	// Pull the latest image
	mg.Deps(DockerPull)

	// Start the server
	mg.Deps(StartServer)

	return nil
}
