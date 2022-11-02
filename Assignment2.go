package main

import (
	"encoding/json"
	"fmt"
)

type Mahasiswa struct {
	Idmhs        int
	Nama         string
	NilaiFisika  int
	NilaiKimia   int
	NilaiBiologi int
	Status       string
}

func main() {
	var mahasiswa []Mahasiswa
	for {
		var pil int

		fmt.Println("Menu")
		fmt.Println("1.Masukkan Data Mahasiswa")
		fmt.Println("2.Tampil Data Mahasiswa")
		fmt.Println("3.Seleksi Mahasiswa")
		fmt.Println("4.EXIT")

		fmt.Println("====================")
		fmt.Print("Silahkan pilih menu: ")
		fmt.Scanf("%d\n", &pil)
		fmt.Println("====================")

		if pil >= 4 {
			fmt.Println("----------------------------")
			fmt.Println("------- EXIT PROGRAM -------")
			fmt.Println("----------------------------")
			break
		}
		switch pil {
		case 1:
			inputMhs(&mahasiswa)
		case 2:
			for _, el := range mahasiswa {
				fmt.Println(stringify(el))
			}
		case 3:
			seleksiMhs(&mahasiswa)
		}
	}
}
func seleksiMhs(mahasiswa *[]Mahasiswa) {
	for i, el := range *mahasiswa {
		if (el.NilaiBiologi+el.NilaiFisika+el.NilaiKimia)/3 > 70 {
			(*mahasiswa)[i].Status = "Lulus"
			fmt.Println(stringify((*mahasiswa)[i]))
		} else {
			(*mahasiswa)[i].Status = "Tidak Lulus"
		}
	}
}
func inputMhs(mahasiswa *[]Mahasiswa) {
	var idmhs int
	var nama string
	var nilaiFisika int
	var nilaiKimia int
	var nilaiBiologi int

	fmt.Print("Enter id: ")
	fmt.Scan(&idmhs)

	fmt.Print("Enter name: ")
	fmt.Scan(&nama)

	fmt.Print("Masukkan Nilai Fisika: ")
	fmt.Scan(&nilaiFisika)

	fmt.Print("Masukkan Nilai Kimia: ")
	fmt.Scan(&nilaiKimia)

	fmt.Print("Enter your Biologi: ")
	fmt.Scan(&nilaiBiologi)

	mahasiswaNew := Mahasiswa{}
	mahasiswaNew.Idmhs = idmhs
	mahasiswaNew.Nama = nama
	mahasiswaNew.NilaiBiologi = nilaiBiologi
	mahasiswaNew.NilaiFisika = nilaiFisika
	mahasiswaNew.NilaiKimia = nilaiKimia

	*mahasiswa = append(*mahasiswa, mahasiswaNew)
	tampilMhs(mahasiswaNew)
}

func stringify(data interface{}) string {
	b, _ := json.MarshalIndent(data, " ", " ")
	return string(b)
}
func tampilMhs(mahasiswa Mahasiswa) {
	fmt.Println(stringify(mahasiswa))
}
