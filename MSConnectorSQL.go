package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	mssql "github.com/microsoft/go-mssqldb"
)

func DBinstance() *sql.DB {
	test1 := fmt.Sprintf("server=%s;",
		mssql.ErrorTypeSliceIsEmpty)
	connString := test1
	connString = fmt.Sprintf("sqlserver://sa:sqlserver@192.168.137.125:49173?database=OMSDB230105EnlargeGuestBook&connection+timeout=30")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	ctx = ctx
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
