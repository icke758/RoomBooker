package main

import(
	"fmt"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func main() {
	var userAction int
	var roomNumber int
	var roomAvaliability string
	var daily float64

	for{
		fmt.Println("Bom dia! Qual ação vocẽ gostaria de realizar?")
		fmt.Println("-----------------------------")
		fmt.Println("1 -> Mostrar quartos")
		fmt.Println("2 -> Reservar quarto")
		fmt.Println("3 -> Disponibilizar quarto")
		fmt.Println("4 -> Criar quarto")
		fmt.Println("5 -> Deletar quarto")
		fmt.Println("-----------------------------")

		fmt.Scanf("%d", &userAction)
		if userAction == 1 {
			fmt.Println("Mostrar quartos")
			fmt.Println(showRooms())
		} else if userAction == 2 {
			fmt.Println("Reservar quarto")
			fmt.Println("Qual quarto você vai reservar?")
			fmt.Scanf("%d", &roomNumber)

			bookRoom(roomNumber)
		} else if userAction == 3 {
			fmt.Println("Disponibilizar quarto")
			fmt.Println("Qual quarto você quer disponibilizar?")
			fmt.Scanf("%d", &roomNumber)

			unbookRoom(roomNumber)
		} else if userAction == 4 {
			fmt.Println("Criar quarto")
			fmt.Println("Numero do quarto: ")
			fmt.Scanf("%d", &roomNumber)
			fmt.Println("Disponibilidade: ")
			fmt.Scanf("%s", &roomAvaliability)
			fmt.Println("Preço diário: ")
			fmt.Scanf("%f", &daily)

			createRoom(roomNumber, roomAvaliability, daily)
		} else if userAction == 5 {
			var id int
			fmt.Println("Apagar Quarto")
			fmt.Println("Qual quarto você quer apagar?")
			fmt.Scanf("%d", &id)

			deleteRoom(id)
		}else{
			fmt.Println("Ok, adeus! :D")
			break
		}
	}
}

func showRooms()(result string){
	db, err := sql.Open("sqlite3", "rooms.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM rooms;")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	 for rows.Next() {
		var id int
		var room int
		var avaliability string
		var price float32

		err = rows.Scan(&id, &room, &avaliability ,&price)

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Quarto Disponibilidade Preço\n", room, avaliability, price)
	}
	return
}

func bookRoom(number int) (result string) {
	db, err := sql.Open("sqlite3", "rooms.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("UPDATE rooms SET avaliability = 'alugado' WHERE number = ?;", number)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto número", number, "reservado!")

	return
}

func createRoom(number int, avaliability string, daily float64){
	db, err := sql.Open("sqlite3", "rooms.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("INSERT INTO rooms (number, avaliability, daily) VALUES (?, ?, ?);", number, avaliability, daily)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto número", number, "criado")

	return
}

func deleteRoom(id int) {
	db, err := sql.Open("sqlite3", "rooms.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM rooms WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto com o Id:", id, "deletado")
	return
}

func unbookRoom(number int) (result string) {
db, err := sql.Open("sqlite3", "rooms.db")

if err != nil {
	log.Fatal(err)
}
defer db.Close()

_, err = db.Exec("UPDATE rooms SET avaliability = 'disponivel' WHERE number = ?;", number)

if err != nil {
	log.Fatal(err)
}

fmt.Println("Quarto número", number, "disponibilizado")
return
}


//TODO 
//Add Function to create new rooms (DONE)
//add Function to unbook a room (DONE)
//add Function to delete a room(DONE)
//CLEAN THE CODE
	//Take the comments out (DONE)
	//Create Modules
//Versioning
//TODO To Future: 
//Create a good CLI Interface.
//Add the calcs to the price and a day counter.
//TODO To More Future:
//Transform this in a API. (Don't know if it's gonna be from scratch or gonna use a Framwork)
