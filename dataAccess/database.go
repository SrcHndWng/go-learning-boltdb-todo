package dataAccess

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const dbFile = "database.db"

// Data contains data in bucket
type Data struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// Create registers new data to bucket.
func Create(bucket string, data string) error {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		err = b.Put([]byte(string(id)), []byte(data))
		if err != nil {
			return err
		}
		return nil
	})
}

// List searches data from bucket and returns.
func List(bucket string) ([]Data, error) {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var datas []Data
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var s string
			for _, x := range k {
				s += fmt.Sprintf("%v", x)
			}
			datas = append(datas, Data{s, string(v)})
		}
		return nil
	})
	return datas, err
}
