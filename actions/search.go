package actions

import (
	"encoding/json"
	"fmt"

	"github.com/SrcHndWng/go-learning-boltdb-todo/dataAccess"
	"github.com/urfave/cli"
)

// Search prints your toto by id.
func Search(c *cli.Context) error {
	id := c.Args().First()
	todo, err := dataAccess.Search("todos", id)
	if todo.ID == "" {
		fmt.Printf("id %s is not found.\n", id)
		return nil
	}
	bytes, err := json.Marshal(&todo)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))
	return nil
}
