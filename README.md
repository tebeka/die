# Perl's "die" for Go

Emulate Perl's [die](http://perldoc.perl.org/functions/die.html) function for
Go.

## Example:

```go
package main

import (
	"os"

	. "github.com/tebeka/die"
)

func main() {
	filename := "/path/to/file"
	file, err := os.Open(filename)
	if err != nil {
		Die("can't open %s - %s", filename, err)
	}
}
```

## Project Info
See https://github.com/tebeka/die
