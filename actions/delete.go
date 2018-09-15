package actions

import (
	"fmt"

	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// Delete deletes your toto by id.
func Delete(c *cli.Context) error {
	if c.Args().First() == "" {
		fmt.Println(invalidArgsMessage)
		return nil
	}
	err := dataAccess.Delete("todos", c.Args().First())
	if err != nil {
		return err
	}
	return nil
}
