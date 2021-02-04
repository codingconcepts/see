// +build !windows

package pkg

import (
	"log"
	"os"
	"os/exec"
)

// Clear a nix terminal.
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Stdout = os.Stderr

	if err := c.Run(); err != nil {
		log.Fatalf("error clearing screen: %v", err)
	}
}
