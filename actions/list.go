package actions

import (
	"encoding/json"
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
	for _, todo := range todos {
		bytes, err := json.Marshal(&todo)
		if err != nil {
			return err
		}
		fmt.Println(string(bytes))
	}
	return nil
}
