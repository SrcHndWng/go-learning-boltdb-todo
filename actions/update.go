package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

// Update updates your toto.
func Update(c *cli.Context) error {
	fmt.Printf("update!! id = %s, todo = %s\n", c.Args()[0], c.Args()[1])
	return nil
}
