package actions

import (
	"fmt"

	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// List prints your todos
func List(c *cli.Context) error {
	todos, err := dataAccess.List("todos")
	if err != nil {
		return err
	}
	fmt.Println(todos)
	return nil
}
