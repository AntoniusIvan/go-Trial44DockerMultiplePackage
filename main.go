package main

import (
	"log"
	"net/http"

	"github.com/AntoniusIvan/go-Trial44DockerMultiplePackage/Controllers"
	"github.com/jmoiron/sqlx"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }
	sqlx := sqlx.NewDb()
	sqlx = sqlx
	//NATIVE SECTION
	mux := http.NewServeMux()
	mux.HandleFunc("/TryMainPkg", Controllers.TryMainPkg)
	//mux.HandleFunc("/TryContrPkg", TryContrPkg)

	err := http.ListenAndServe(":49318", mux)
	if err != nil {
		log.Println("There was an error listening on port :49318", err)
	}

}
