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
}

func main() {
	var mahasiswa []Mahasiswa
	var mahasiswaLulus []Mahasiswa
	for {
		var pil int

		fmt.Println("Menu")
		fmt.Println("1.Input Data Staff")
		fmt.Println("2.Hitung Gaji Total")
		fmt.Println("3.Tampil Data Staff")
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
			for _, el := range mahasiswaLulus {
				fmt.Println(stringify(el))
			}
		}
	}
}
func seleksiMhs(mahasiswa *[]Mahasiswa) {
	for _, el := range *mahasiswa {
		fmt.Printf("hitung total gaji %s \n", el.Nama)
		if (el.NilaiBiologi + el.NilaiFisika + el.NilaiKimia/3) > 70 {
			mahasiswaNew := Mahasiswa{}
			mahasiswaNew.Idmhs = el.Idmhs
			mahasiswaNew.Nama = el.Nama
			mahasiswaNew.NilaiBiologi = el.NilaiBiologi
			mahasiswaNew.NilaiFisika = el.NilaiFisika
			mahasiswaNew.NilaiKimia = el.NilaiKimia
			*mahasiswa = append(*mahasiswa, mahasiswaNew)
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

	fmt.Print("Masukkan Nilai kimia: ")
	fmt.Scan(&nilaiKimia)

	fmt.Print("Masukkan Nilai Biologi: ")
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
