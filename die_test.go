package die

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestDie(t *testing.T) {
	gofile := "/tmp/die-testprog.go"
	if err := ioutil.WriteFile(gofile, testprog, 0666); err != nil {
		t.Fatalf("can't create go file")
	}

	arg := time.Now().UTC().String()
	cmd := exec.Command("go", "run", gofile, arg)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err == nil {
		t.Fatalf("completed normally, should have failed")
	}

	expected := fmt.Sprintf("error: %s", arg)
	lines := strings.Split(stderr.String(), "\n")
	out := lines[0] // go run add it's own junk at the end
	if out != expected {
		t.Fatalf("bad data, got:\n%s\nwanted:\n%s", out, expected)
	}
}

var testprog = []byte(`
// Test program for die
package main

import (
	"flag"

	. "github.com/tebeka/die"
)

var data = ""

func main() {
	flag.Parse()
	data = flag.Arg(0)

	Die("%s", data)
}
`)
