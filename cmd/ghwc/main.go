//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package main

import (
	"fmt"

	"github.com/alihassan4198-tech/ghw"
	"github.com/alihassan4198-tech/ghw/cmd/ghwc/commands"
)

var (
	// version of application at compile time (-X 'main.version=$(VERSION)').
	version = "(Unknown Version)"
	// buildHash GIT hash of application at compile time (-X 'main.buildHash=$(GITCOMMIT)').
	buildHash = "No Git-hash Provided."
	// buildDate of application at compile time (-X 'main.buildDate=$(BUILDDATE)').
	buildDate = "No Build Date Provided."
)

func main() {
	commands.Execute(version, buildHash, buildDate)
	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting CPU info: %v", err)
	}

	fmt.Printf("%v\n", cpu)
}
