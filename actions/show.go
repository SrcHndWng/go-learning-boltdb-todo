package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

// Show prints your toto by id.
func Show(c *cli.Context) error {
	fmt.Printf("show!! id = %s\n", c.Args().First())
	return nil
}
