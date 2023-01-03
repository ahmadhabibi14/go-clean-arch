package routes

import (
	"go-rest-api/controllers"
	"net/http"
)

func Route() {
	// curl -X POST -d '{"nama":"Ahmad Rizky Nusantara Habibi", "prodi":"Ilmu Komputer", "semester":1}' localhost:8080/addMahasiswa
	http.HandleFunc("/addMahasiswa", controllers.AddMahasiswa) // Method POST

	// curl -X GET localhost:8080/mahasiswa
	http.HandleFunc("/mahasiswa", controllers.GetAllMahasiswa) // Method GET

	// curl -X PUT -d '{"id":1, "nama":"Ahmad Rizky Nusantara Habibi", "prodi":"Ilmu Komputer", "semester":1}' localhost:8080/updateMahasiswa
	http.HandleFunc("/updateMahasiswa", controllers.UpdateMahasiswa) // Method PUT

	// curl -X DELETE -d '{"id":1}' localhost:8080/deleteMahasiswa
	http.HandleFunc("/deleteMahasiswa", controllers.DeleteMahasiswa) // Method DELETE
}
