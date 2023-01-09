package controllers

import (
	"encoding/json"
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"log"
	"net/http"
	"os"
)

// GET semua Mahasiswa
func GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	var allMahasiswa []models.Mahasiswa
	db := database.Connect()
	var rows, err = db.Query(
		"SELECT * FROM mahasiswa",
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer rows.Close()
	for rows.Next() {
		var each = models.Mahasiswa{}
		var err = rows.Scan(
			&each.Id,
			&each.Nama,
			&each.Prodi,
			&each.Semester,
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		allMahasiswa = append(allMahasiswa, each)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.WriteHeader(http.StatusOK)
	// Ini untuk mengonversi object struct ke json
	json.NewEncoder(w).Encode(allMahasiswa)
}

// Tambah Mahasiswa
func AddMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
	addMahasiswa := models.AddMahasiswa{}
	// Ini untuk mengonversi data json ke bentuk object struct
	err := json.NewDecoder(r.Body).Decode(&addMahasiswa)
	if err != nil {
		fmt.Println("error occurred while decoding mahasiswa data :: ")
		return
	}
	db := database.Connect()
	_, err = db.Exec(
		"INSERT INTO `mahasiswa`(`nama`, `prodi`, `semester`) VALUES (?, ?, ?)",
		addMahasiswa.Nama,
		addMahasiswa.Prodi,
		addMahasiswa.Semester,
	)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		log.Printf("Berhasil menambahkan %s sebagai Mahasiswa", addMahasiswa.Nama)
	}
	defer db.Close()
}

// Hapus Mahasiswa
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.WriteHeader(http.StatusOK)
	deleteMahasiswa := models.Mahasiswa{}
	// Ini untuk mengonversi data json ke bentuk object struct
	err := json.NewDecoder(r.Body).Decode(&deleteMahasiswa)
	if err != nil {
		fmt.Println("error occurred while decoding mahasiswa data :: ")
		return
	}

	db := database.Connect()
	_, err = db.Exec(
		"DELETE FROM mahasiswa WHERE id = ?",
		deleteMahasiswa.Id,
	)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		log.Printf(
			"Berhasil menghapus %s sebagai Mahasiswa",
			deleteMahasiswa.Nama,
		)
	}
	defer db.Close()
}

// Update Mahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.WriteHeader(http.StatusOK)
	updateMahasiswa := models.Mahasiswa{}
	// Ini untuk mengonversi data json ke bentuk object struct
	err := json.NewDecoder(r.Body).Decode(&updateMahasiswa)
	if err != nil {
		fmt.Println("error occurred while decoding mahasiswa data :: ")
		return
	}
	db := database.Connect()
	_, err = db.Exec(
		"UPDATE mahasiswa SET id = ?, nama = ?, prodi = ?, semester = ? WHERE id = ?",
		updateMahasiswa.Id,
		updateMahasiswa.Nama,
		updateMahasiswa.Prodi,
		updateMahasiswa.Semester,
		updateMahasiswa.Id,
	)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		log.Printf(
			"Berhasil mengupdate %s sebagai Mahasiswa",
			updateMahasiswa.Nama,
		)
	}
	defer db.Close()
}
