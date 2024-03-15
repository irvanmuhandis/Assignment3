package route

import (
	"assignment3/controller"
	"fmt"
	"net/http"
)

var PORT = ":8080"

func StartServer() {
	http.HandleFunc("/status", controller.GetStatus)
	http.HandleFunc("/statusjson", controller.GetStatusJSON)
	fmt.Println("Application is listening on PORT ", PORT)
	fmt.Println("Open web http://localhost:8080/status")
	http.ListenAndServe(PORT, nil)
}
