package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

// Create registers your toto.
func Create(c *cli.Context) error {
	fmt.Println("create!! todo = " + c.Args().First())
	return nil
}
