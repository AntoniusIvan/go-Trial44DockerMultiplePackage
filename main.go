package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	mux.HandleFunc("/TryMainPkg", TryMainPkg)
	//mux.HandleFunc("/TryContrPkg", TryContrPkg)

	err := http.ListenAndServe(":49318", mux)
	if err != nil {
		log.Println("There was an error listening on port :49318", err)
	}

}

func TryMainPkg(w http.ResponseWriter, r *http.Request) {

	jsonData, err := json.Marshal("HellowIVaunMultiple")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	//w.Write([]byte("oioi"))
	w.Write([]byte(string(jsonData)))
}
