/* 
Emulate Perl's "die" function (http://perldoc.perl.org/functions/die.html)

Example:

	import (
		. "die"
		"os"
	)

	func main() {
		filename := "/path/to/file"
		file, err := os.Open(filename)
		if err != nil {
			Die("can't open %s - %s", filename, err)
		}
	}
*/
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
