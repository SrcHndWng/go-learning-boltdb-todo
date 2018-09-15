package actions

import (
	"fmt"

	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// Create registers your toto.
func Create(c *cli.Context) error {
	if c.Args().First() == "" {
		fmt.Println(invalidArgsMessage)
		return nil
	}
	err := dataAccess.Create("todos", c.Args().First())
	if err != nil {
		return err
	}
	return nil
}
