package actions

import (
	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// Create registers your toto.
func Create(c *cli.Context) error {
	todo, err := dataAccess.NewTodo()
	if err != nil {
		return err
	}

	err = todo.Create(c.Args().First())
	if err != nil {
		return err
	}
	return nil
}
