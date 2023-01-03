package models

type Mahasiswa struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Prodi    string `json:"prodi"`
	Semester int    `json:"semester"`
}

type AddMahasiswa struct {
	Nama     string `json:"nama"`
	Prodi    string `json:"prodi"`
	Semester int    `json:"semester"`
}
