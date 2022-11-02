package main

import "fmt"

//
type Person interface {
	MyHobby() string
}

type Worker struct {
	IdWorker int
	Nama     string
	Hobby    string
}

func (worker Worker) MyHobby() string {
	return "I am just a worker"
}

type Staff struct {
	Worker
}

func (staff Staff) MyHobby() string {
	return staff.Hobby
}

type Manager struct {
	Worker
}

func (manager Manager) MyHobby() string {
	return manager.Hobby
}

// this function Area works on the type Rectangle and has the same function signature
// defined in the interface Shaper. Therefore, Rectangle now implements the interface Shaper.

func main() {
	var id int
	var nama string
	var hobby string

	for {
		var pil int

		fmt.Println("Menu")
		fmt.Println("1.Role Worker")
		fmt.Println("2.Role Staff")
		fmt.Println("3.Role Manager")
		fmt.Println("4.EXIT")
		fmt.Println("")

		fmt.Println("====================")
		fmt.Print("Silahkan pilih menu: ")
		fmt.Scanf("%d\n", &pil)
		fmt.Println("====================")
		fmt.Println("")

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
			worker := Worker{
				IdWorker: id,
				Nama:     nama,
			}
			fmt.Println("Id Worker: ", worker.IdWorker)
			fmt.Println("Nama Worker: ", worker.Nama)
			fmt.Println("======= Hobby =======")
			fmt.Printf("Well, %s \n", worker.MyHobby())

			break
		case 2:
			fmt.Print("Enter id: ")
			fmt.Scan(&id)

			fmt.Print("Enter your name: ")
			fmt.Scan(&nama)

			fmt.Print("Enter your hobby: ")
			fmt.Scan(&hobby)
			staff := Staff{
				Worker{
					IdWorker: id,
					Nama:     nama,
					Hobby:    hobby,
				},
			}
			fmt.Println("Id Worker: ", staff.IdWorker)
			fmt.Println("Nama Worker: ", staff.Nama)
			fmt.Println("======= Hobby =======")
			fmt.Printf("My Hobby is %s \n", staff.MyHobby())
			break
		case 3:
			fmt.Print("Enter id: ")
			fmt.Scan(&id)

			fmt.Print("Enter your name: ")
			fmt.Scan(&nama)

			fmt.Print("Enter your hobby: ")
			fmt.Scan(&hobby)
			manager := Manager{
				Worker{
					IdWorker: id,
					Nama:     nama,
					Hobby:    hobby,
				},
			}

			fmt.Printf("My Hobby is %s \n", manager.MyHobby())

			break
		default:
			fmt.Println("====Exit Program====")
		}
	}

}
