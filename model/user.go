package model

type User struct {
	Username string
	Password string // Password yang sudah terenkripsi
}

var Users = []User{} // Menyimpan data user di memori
