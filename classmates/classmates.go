package classmates

import "fmt"

type Classmates struct {
	absen int 
	nama string
	alamat string
	pekerjaan string
	alasan string
}

var Mentee = []Classmates{
	{absen: 1, nama: "Fariz Dandy Lazuardi", alamat: "Bekasi", pekerjaan: "Software Engineer", alasan: "Ingin menjadi Backend Developer sejati"},
	{absen: 2, nama: "John Doe", alamat: "Jakarta", pekerjaan: "Product Designer", alasan: "Ingin Belajar Golang"},
	{absen: 3, nama: "Marie Zhang", alamat: "Surabaya", pekerjaan: "Frontend Developer", alasan: "Mau explore Backend dengan Golang"},
	{absen: 4, nama: "Lily Moron", alamat: "Melbourne", pekerjaan: "UI Designer", alasan: "Mau belajar coding dengan Golang"},
	{absen: 5, nama: "Marcus Rashford", alamat: "Manchester", pekerjaan: "Runner", alasan: "Golang is high demand!"},
}

func (mates Classmates) Result(){
	fmt.Printf("Nama : %v\n", mates.nama)
	fmt.Printf("Alamat : %v\n", mates.alamat)
	fmt.Printf("Pekerjaan : %v\n", mates.pekerjaan)
	fmt.Printf("Alasan Mengikuti kelas Golang di Hacktiv8  : %v\n", mates.alasan)

}