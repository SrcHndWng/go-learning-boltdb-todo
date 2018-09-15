package main

import (
	"log"
	"os"

	"github.com/SrcHndWng/go-learning-boltdb-todo/actions"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-learning-boltdb-todo"
	app.Usage = "register your todos"
	app.Commands = []cli.Command{
		{
			Name:      "list",
			ShortName: "l",
			Usage:     "Shows your todo list",
			UsageText: "binary-path list",
			Action:    actions.List,
		},
		{
			Name:      "create",
			ShortName: "c",
			Usage:     "Creates your todo",
			UsageText: "binary-path create \"your todo\"",
			Action:    actions.Create,
		},
		{
			Name:      "search",
			ShortName: "s",
			Usage:     "Searchs your todo by id",
			UsageText: "binary-path show your-todo-id(ex.1)",
			Action:    actions.Search,
		},
		{
			Name:      "update",
			ShortName: "u",
			Usage:     "Updates your todo by id",
			UsageText: "binary-path update your-todo-id(ex.1) \"your new todo\"",
			Action:    actions.Update,
		},
		{
			Name:      "delete",
			ShortName: "d",
			Usage:     "Deletes your todo by id",
			UsageText: "binary-path delete your-todo-id(ex.1)",
			Action:    actions.Delete,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
