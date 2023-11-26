// SPDX-License-Identifier: BSD-3-Clause

package registryd

import (
	"bytes"
	"encoding/gob"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

func writeData(db *bolt.DB, bucket, key string, value any) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		if err := gob.NewEncoder(&buf).Encode(value); err != nil {
			return err
		}

		return b.Put([]byte(key), buf.Bytes())
	})
}

func readData(db *bolt.DB, bucket, key string, value any) error {
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s does not exist", bucket)
		}

		v := b.Get([]byte(key))
		if v == nil {
			return fmt.Errorf("key %s does not exist", key)
		}

		return gob.NewDecoder(bytes.NewReader(v)).Decode(value)
	})
}

func deleteData(db *bolt.DB, bucket, key string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s does not exist", bucket)
		}

		return b.Delete([]byte(key))
	})
}

func getKeys(db *bolt.DB, bucket string) ([]string, error) {
	var keys []string

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s does not exist", bucket)
		}

		return b.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
	})

	return keys, err
}
