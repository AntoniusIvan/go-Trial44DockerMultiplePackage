package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
