package dataAccess

import (
	"github.com/SrcHndWng/go-learning-boltdb-todo/utils"
	"github.com/boltdb/bolt"
)

const dbFile = "todos.db"
const bucket = "todos"

// Todo implements crud for todo.
type Todo struct {
	db *bolt.DB
}

// NewTodo creates Todo instance and returns.
func NewTodo() (*Todo, error) {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, err
	}

	todo := Todo{db}
	return &todo, nil
}

// Create registers data.
func (todo *Todo) Create(data string) error {
	return todo.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		err = b.Put(utils.Itob(id), []byte(data))
		if err != nil {
			return err
		}
		return nil
	})
}

// List searchs todo and returns.
func (todo *Todo) List() ([]string, error) {
	var todos []string
	err := todo.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			todos = append(todos, string(v))
		}
		return nil
	})
	return todos, err
}
