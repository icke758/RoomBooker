package menu

import (
	"fmt"
	"room/utils"
)

func ShowMenuOptions() {
	menuMap := map[string]string{
		"1": "Mostrar quartos",
		"2": "Reservar quarto",
		"3": "Disponibilizar quarto",
		"4": "Criar quarto",
		"5": "Deletar quarto",
	}

	keys := utils.SortKeys(menuMap)

	for _, key := range keys {
		fmt.Printf("[%s] -> [%s] \n", key, menuMap[key])
	}
	fmt.Println("-----------------------------")
}

func ShowRoomsMenu(rooms string) {
	fmt.Println("Mostrar quartos")
	fmt.Println(rooms)
}

func ShowRoomsToBook() {
	fmt.Println("Reservar quarto")
	fmt.Println("Qual quarto você vai reservar?")
}
