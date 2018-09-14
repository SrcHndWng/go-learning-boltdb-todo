package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

// Delete deletes your toto by id.
func Delete(c *cli.Context) error {
	fmt.Printf("delete!! id = %s\n", c.Args().First())
	return nil
}
