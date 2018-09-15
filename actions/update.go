package actions

import (
	"fmt"

	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// Update updates your toto.
func Update(c *cli.Context) error {
	if (c.Args().Get(0) == "") || (c.Args().Get(1) == "") {
		fmt.Println(invalidArgsMessage)
		return nil
	}
	err := dataAccess.Update("todos", c.Args().Get(0), c.Args().Get(1))
	if err != nil {
		return err
	}
	return nil
}
