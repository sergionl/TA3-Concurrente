package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	//"time"
	"sync"
)

type Player struct {
	pos_objetivo int  //Meta del jugador
	pos_inicial  int  //posicion inicial
	pos          int  //posicion actual
	continuar    bool //si sigue avanzando o no
	equipo       int  //A cuál equipo pertenece: 1 o 2
	id           int  //ID del jugador
	meta         bool //Si llegó a la meta o no
}

func duelo(n int, players []Player, current1 int, current2 int, player1_pos chan int, player2_pos chan int) int {
	fmt.Printf("Duelo ente jugador %d y jugador %d en la posición %d\n", current1, current2, players[current1].pos)
	//0 = piedra, 1 = papel, 2 = tijera
	jugada1 := rand.Intn(3)
	jugada2 := rand.Intn(3)
	for {
		//empate
		if jugada1 != jugada2 {
			//gana jugador 1
			if (jugada1 == 0 && jugada2 == 1) || (jugada1 == 1 && jugada2 == 2) || (jugada1 == 2 && jugada2 == 0) {
				fmt.Printf("Gano jugador %d\n", current1)
				//fmt.Printf("pos Inicial %d\n", players[current2].pos_inicial)
				//player2_pos <- players[current2].pos_inicial //jugador 2 vuelve a empezar
				players[current2].pos = players[current2].pos_inicial
				current2--
				if current2 <= (n/2)-1 {
					current2 = n - 1
				}

				current2 = cambiarE2(n, players, current2)
				fmt.Printf("Ahora le tocara jugar a jugador %d\n por que el otro perdio\n", current2)
				return current2

			} else { //gana jugador 2
				fmt.Printf("Gano jugador %d\n", current2)
				//player1_pos <- players[current1].pos_inicial //jugador 1 vuewlve a empezar
				//fmt.Printf("pos Inicial %d\n", players[current1].pos_inicial)
				players[current1].pos = players[current1].pos_inicial
				current1++
				if current1 == (n / 2) {
					current1 = 0
				}
				current1 = cambiarE1(n, players, current1)
				fmt.Printf("Ahora le tocara jugar a jugador %d\n por que el otro perdio\n", current1)
				return current1

			}

		} else {
			jugada1 = rand.Intn(3)
			jugada2 = rand.Intn(3)
		}
	}

}

func avanzarE1(players []Player, i int, player1_pos chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	players[i].pos += 1
	player1_pos <- players[i].pos
	if players[i].pos == players[1].pos_objetivo {
		fmt.Printf("El jugador %d del equipo %d llego al final\n", i, players[i].equipo)
		players[i].meta = true
		//player1_pos <- 1
	}
	//fmt.Printf("El jugador %d del equipo %d avanza a la posición %d\n", i, players[i].equipo, players[i].pos)
}
func avanzarE2(players []Player, i int, player2_pos chan int, wg *sync.WaitGroup) {
	// validar
	defer wg.Done()
	players[i].pos -= 1
	player2_pos <- players[i].pos
	if players[i].pos == players[i].pos_objetivo {
		fmt.Printf("El jugador %d del equipo %d llego al final\n", i, players[i].equipo)
		players[i].meta = true
		//player2_pos <- 20
	}
	//fmt.Printf("El jugador %d del equipo %d avanza a la posición %d\n", i, players[i].equipo, players[i].pos)

}

func ganar(players []Player, n int) bool {
	cont := 0
	//equipo 1
	for id := 0; id < n/2+1; id++ {
		if players[id].meta == true {
			cont++
		}
	}
	if cont == n/2 {
		fmt.Printf("Gano Equipo 1")
		return true
	}

	cont = 0
	//equipo 2
	for id := n / 2; id < n; id++ {
		if players[id].meta == true {
			cont++
		}
	}
	if cont == n/2 {
		fmt.Printf("Gano Equipo 2")
		return true
	}
	return false
}

func cambiarE1(n int, players []Player, i int) int {
	cont := 0
	for {
		if cont == n/2 {
			break
		} else {
			if players[i].meta == true {
				i++
				if i == (n / 2) {
					i = 0
				}
			} else {

				return i
			}
		}

		cont++
	}
	return i
}
func cambiarE2(n int, players []Player, i int) int {
	cont := 0
	for {
		if cont == n/2 {
			break
		} else {
			if players[i].meta == true {
				i--
				if i == (n/2)-1 {
					i = n - 1
				}
			} else {

				return i
			}
		}

		cont++
	}
	return i
}
func Controller(n int, players []Player) {

	var wg sync.WaitGroup

	//duel_winner := make([]chan int) //Ganador del enfrentamiento
	player1_pos := make(chan int)
	player2_pos := make(chan int) //Almacenan las posiciones de 1 jugador de cada equipo//

	//losser := make(chan int)

	//player1_id := make(chan int)
	//player2_id := make(chan int)

	//players_pos[0] = jugador equipo 1, players_pos[1] = jugador equipo 2

	fmt.Println("Jugadores y sus equipos: ")

	//Crear los jugadores y asignarlos a su equipo correspondiente
	for i := 0; i < n; i++ {

		players[i].meta = false //setear que aún no llegan a la meta
		players[i].id = i       //Asignar el id

		//Asignar equipos y posiciones
		if i < n/2 {
			players[i].equipo = 1
			players[i].pos_inicial = 1
			players[i].pos_objetivo = 20
			players[i].pos = 1
		} else {
			players[i].equipo = 2
			players[i].pos_inicial = 20
			players[i].pos_objetivo = 1
			players[i].pos = 20
		}
		fmt.Printf("EL jugador %d se creó y está en el equipo %d\n", i, players[i].equipo)
	}
	match_end := false
	players[0].continuar = true   //1er jugador del equipo 1 avance
	players[n-1].continuar = true //1er jugador del equipo 2 avance
	current1 := 0
	current2 := n - 1
	//player1_id <- 0
	//player2_id <- n - 1
	cont := 0
	pos1 := 1
	pos2 := 20
	//perdedor := 0
	for {
		//fmt.Printf("bucle")
		//wg.Wait()
		if pos1 == pos2 {
			perdedor := duelo(n, players, current1, current2, player1_pos, player2_pos)
			if perdedor >= n/2 {
				current2 = perdedor
			} else {
				current1 = perdedor
			}
		}
		wg.Add(1)
		go avanzarE1(players, current1, player1_pos, &wg)
		current1 = cambiarE1(n, players, current1)
		pos1 = <-player1_pos
		wg.Wait()
		if pos1 == pos2 {
			perdedor := duelo(n, players, current1, current2, player1_pos, player2_pos)
			if perdedor >= n/2 {
				current2 = perdedor
			} else {
				current1 = perdedor
			}
		}
		wg.Add(1)
		current2 = cambiarE2(n, players, current2)
		go avanzarE2(players, current2, player2_pos, &wg)

		pos2 = <-player2_pos

		//duelo(players, current1, current2, player1_pos, player2_pos)
		//match_end = true
		cont++

		fmt.Printf("El jugador %d del equipo %d avanza a la posición %d\n", current1, players[current1].equipo, pos1)
		fmt.Printf("El jugador %d del equipo %d avanza a la posición %d\n", current2, players[current2].equipo, pos2)

		match_end = ganar(players, n)
		//if cont == 20 {
		//	break
		//}
		if match_end == true {
			break
		}
	}

}

func main() {
	var (
		n int
	)

	if len(os.Args) > 1 {
		if x, err := strconv.Atoi(os.Args[1]); err != nil {
			return
		} else {
			n = x

		}
	} else {
		n = 10
	}

	if n%2 != 0{
		fmt.Println("El número debe ser par")
		return 
	}
	//n = 10 //Cantidad de jugadores
	players := make([]Player, n)

	//Controller(n_players, players)
	Controller(n, players)
}
