package main

import (
	"fmt"
	"net/http"

	"go-rest-api/routes"
)

func main() {
	routes.Route()
	fmt.Println("Server siap di http://localhost:8080/")
	fmt.Print("route : \n /addMahasiswa \n /mahasiswa \n /updateMahasiswa \n /deleteMahasiswa \n")
	fmt.Println("--------------------")
	http.ListenAndServe(":8080", nil)
}
