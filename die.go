// Emulate Perl's "die" function
package die

import (
	"fmt"
	"os"
)

// Die prints error message and aborts the program
func Die(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "error: %s\n", msg)
	os.Exit(1)
}
