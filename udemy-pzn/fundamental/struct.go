package main

import "fmt"

type Mahasiswa struct {
	Nama, Email string
	NIM         uint64
}

type Dosen struct {
	Nama, Email string
	NIDN        uint32
}

func main() {
	var firstMhs Mahasiswa
	fmt.Println(firstMhs)
	firstMhs.Nama = "Trafalgar D. Water Law"
	firstMhs.Email = "lawonepiece@gmail.com"
	firstMhs.NIM = 201822004311

	fmt.Println(firstMhs)
	fmt.Println(firstMhs.Nama)
	fmt.Println(firstMhs.Email)
	fmt.Println(firstMhs.NIM)

	dosen := Dosen{
		Nama:  "Aditya Prabowo",
		Email: "dosenprabowo@gmail.com",
		NIDN:  117891,
	}

	fmt.Println(dosen)
	fmt.Println(dosen.Nama)
	fmt.Println(dosen.Email)
	fmt.Println(dosen.NIDN)
}
