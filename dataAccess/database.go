package dataAccess

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
)

const dbFile = "database.db"

// Data contains data in bucket
type Data struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// Create registers new data to bucket.
func Create(bucket, data string) error {
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
		err = b.Put(keybytes(id), []byte(data))
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
			datas = append(datas, Data{keystr(k), string(v)})
		}
		return nil
	})
	return datas, err
}

// Search finds data by bucket, id.
func Search(bucket, id string) (Data, error) {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return Data{}, err
	}
	defer db.Close()

	var data Data
	err = db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(bucket)).Cursor()
		prefix, err := idToBytes(id)
		if err != nil {
			return err
		}
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			data = Data{keystr(k), string(v)}
		}

		return nil
	})
	return data, err
}

// Update updates data.
func Update(bucket, id, data string) error {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	key, err := idToBytes(id)
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put(key, []byte(data))
		return err
	})
}

// Delete deletes data.
func Delete(bucket, id string) error {
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	key, err := idToBytes(id)
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Delete(key)
		return err
	})
}

func keystr(k []byte) string {
	var s string
	for _, x := range k {
		s += fmt.Sprintf("%v", x)
	}
	return s
}

func idToBytes(id string) ([]byte, error) {
	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}
	return keybytes(u), nil
}

func keybytes(u uint64) []byte {
	return []byte(string(u))
}
