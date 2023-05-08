package logging_test

import (
	"github.com/TeamMomentum/go-pkg/logging"
	"os"
)

func ExamplePrintln_empty() {
	logging.SetOutput(os.Stdout)

	logging.Println()
	// Output: {"severity":"INFO"}
}

func ExamplePrintln_hello_world() {
	logging.SetOutput(os.Stdout)

	logging.Println("hello", "world")
	// Output: {"severity":"INFO","message":"hello world"}
}
