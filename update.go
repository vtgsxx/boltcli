package main

import (
	// "fmt"
	"log"
	bolt "go.etcd.io/bbolt"
)


func UpdateKey(db *bolt.DB, bucket string, k string, v string){
	if err := db.Update(func (tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(k), []byte(v))
		
		return err
	}); err != nil{
		log.Fatal(err)
		return
	}
}