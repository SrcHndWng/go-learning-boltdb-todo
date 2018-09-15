package actions

import (
	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// Update updates your toto.
func Update(c *cli.Context) error {
	err := dataAccess.Update("todos", c.Args()[0], c.Args()[1])
	if err != nil {
		return err
	}
	return nil
}
