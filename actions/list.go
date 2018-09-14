package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

// List prints your todos
func List(c *cli.Context) error {
	fmt.Println("list!!")
	return nil
}
