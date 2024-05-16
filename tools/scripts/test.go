package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			db, err := sql.Open("mysql", "root@tcp(127.0.0.1:15306)/mysql")
			if err != nil {
				panic(err.Error())
			}
			defer db.Close()
			for j := 0; j < 1000; j++ {
				db.Query("select * from d1.t1")
				//fmt.Println("1")
			}
			wg.Done()
		}()

	}
	wg.Wait()
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:15306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	_, err = db.Query("select * from d1.t1")
	if err != nil {
		log.Printf("final: " + err.Error())
	}
}
