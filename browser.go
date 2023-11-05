// Package browser provides helpers to open files, readers, and urls in a browser window.
//
// The choice of which browser is started is entirely client dependant.
package browser

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Stdout is the io.Writer to which executed commands write standard output.
var Stdout io.Writer = os.Stdout

// Stderr is the io.Writer to which executed commands write standard error.
var Stderr io.Writer = os.Stderr

// Open opens a new browser window pointing to url.
func Open(url string) error {
	return openBrowser(fmt.Sprintf("file://%s", url))
}

func runCmd(prog string, args ...string) error {
	cmd := exec.Command(prog, args...)
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	return cmd.Run()
}
