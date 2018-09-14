package actions

import (
	"fmt"

	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// List prints your todos
func List(c *cli.Context) error {
	// fmt.Println("list!!")
	todo, err := dataAccess.NewTodo()
	if err != nil {
		return err
	}
	todos, err := todo.List()
	if err != nil {
		return err
	}
	fmt.Println(todos)
	return nil
}
