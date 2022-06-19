package domain

import "time"

type Anggota struct {
	IdAgt     int       `json:"id_agt"`
	Nim       string    `json:"nim"`
	Nama      string    `json:"nama"`
	Ttl       string    `json:"ttl"`
	JK        string    `json:"jk"`
	Alamat    string    `json:"alamat"`
	Prodi     string    `json:"prodi"`
	Angkatan  string    `json:"angkatan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type AnggotaRepository interface {
	Add(cursor string, num int)
	FetchAll()
	FetchById()
	UpdateById()
	DeleteById()
}
