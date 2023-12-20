package main

import (
	"log"
	"os"
	"time"
	"flag"
	bolt "go.etcd.io/bbolt"
)

var command string
var dbFile string
var bucket string
var key string
var keyT string
var value string
var valueT string

func init(){
	flag.StringVar(&command, "command", "", "command")
	flag.StringVar(&dbFile, "db", "", "dbFile")
	flag.StringVar(&bucket, "bucket", "", "bucket")
	flag.StringVar(&key, "key", "", "key")
	flag.StringVar(&keyT, "keyT", "", "keyT")
	flag.StringVar(&value, "value", "", "value")
	flag.StringVar(&valueT, "valueT", "", "valueT")
	flag.Parse()
}


func main(){
	requireFlagPtr := []*string{&command, &dbFile}
	for _, v := range requireFlagPtr{
		if len(*v) == 0{
			log.Fatal("not enought args")
			return
		}
	}

	_, err := os.Stat(dbFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := bolt.Open(dbFile, 0600, &bolt.Options{Timeout: 2 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	switch command{
	case "view":
		if len(bucket) == 0{
			ViewAllBucket(db)
		} else{
			if len(key) == 0{
				ViewBucket(db, bucket)
			} else{
				ViewKey(db, bucket, key)
			}
		}
	case "update":
	
	default:
		log.Fatal("not yet support command")
	}
}