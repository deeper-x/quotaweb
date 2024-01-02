package main

import (
	"log"

	"github.com/deeper-x/quotaweb/memdb"
)

var q memdb.Quoter

func main() {
	q := memdb.NewQuota()

	ok, err := q.Set()
	if err != nil {
		panic(err)
	}

	if !ok {
		log.Println("client not allowed")
		return
	}

	res, err := q.Get()
	if err != nil {
		panic(err)
	}

	log.Println(res)
}
