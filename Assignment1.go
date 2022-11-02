package main

import (
	"fmt"
)

type Staff struct {
	idStaff    int    `json:"id"`
	namaStaff  string `json:"nama"`
	gapokStaff int    `json:"gapok"`
	tunjStaff  int    `json:"tunjangan"`
	totalStaff int    `json:"totalgaji"`
}

func main() {

	fmt.Println("========== 1 Dimension =============")

	var id int
	var nama string
	var gapok int
	var tunjangan int
	var totgaji int
	staff := Staff{}

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
			fmt.Print("Enter id: ")
			fmt.Scan(&id)

			fmt.Print("Enter your name: ")
			fmt.Scan(&nama)

			fmt.Print("Enter your salary: ")
			fmt.Scan(&gapok)

			fmt.Print("Enter your allowance: ")
			fmt.Scan(&tunjangan)

			staff.idStaff = id
			staff.namaStaff = nama
			staff.gapokStaff = gapok
			staff.tunjStaff = tunjangan

			fmt.Printf("Id Staff : %d\n", staff.idStaff)
			fmt.Printf("Nama Staff : %s\n", staff.namaStaff)
			fmt.Printf("Gaji Pokok Staff : %d\n", staff.gapokStaff)
			fmt.Printf("Tunjangan Staff : %d\n", staff.tunjStaff)

			fmt.Scanf("%d\n", &pil)

			break
		case 2:
			totgaji = gapok + tunjangan
			staff.totalStaff = totgaji

			fmt.Printf("Perhitungan Total Gaji : %d atas nama staff %s\n", totgaji, staff.namaStaff)
			break
		case 3:
			fmt.Println("===========Data Staff===========")
			fmt.Printf("Id Staff : %d\n", staff.idStaff)
			fmt.Printf("Nama Staff : %s\n", staff.namaStaff)
			fmt.Printf("Gaji Pokok Staff : %d\n", staff.gapokStaff)
			fmt.Printf("Tunjangan Staff : %d\n", staff.tunjStaff)
			fmt.Printf("Total Gaji Staff : %d\n", staff.totalStaff)

			break
		default:
			fmt.Println("====Exit Program====")
		}
	}

}
