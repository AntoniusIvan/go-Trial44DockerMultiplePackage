package main

import (
	"log"
	"net/http"

	"github.com/AntoniusIvan/go-Trial44DockerMultiplePackage/Controllers"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	//NATIVE SECTION
	mux := http.NewServeMux()
	mux.HandleFunc("/TryMainPkg", Controllers.TryMainPkg)
	//mux.HandleFunc("/TryContrPkg", TryContrPkg)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println("There was an error listening on port :8080", err)
	}

}
