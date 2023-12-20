package main

import (
	"fmt"
	"log"
	bolt "go.etcd.io/bbolt"
)

func ViewAllBucket(db *bolt.DB){
	if err := db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
				fmt.Println(string(name))
				return nil
		})
	}); err != nil {
		fmt.Println(err)
		return
	}
}

func ViewBucket(db *bolt.DB, bucket string){
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key: %q\nvalue: %q\n", k, v)
			return nil
		})
		return nil
	}); err != nil{
		log.Fatal(err)
	}
}

func ViewKey(db *bolt.DB, bucket string, key string){
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v := b.Get([]byte(key))
		if v == nil {
			log.Fatal("key not exists")
			return nil
		}
		fmt.Printf("%q", v)
		return nil
	}); err != nil{
		log.Fatal(err)
	}
}