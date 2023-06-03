package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Player struct {
	pos_objetivo int  //Meta del jugador
	pos_inicial  int  //posicion inicial
	pos          int  //posicion actual
	continuar    bool //si sigue avanzando o no
	equipo       int  //A cuál equipo pertenece: 1 o 2
	id           int  //ID del jugador
	meta		 bool //Si llegó a la meta o no
}

func NextPlayer(e int, _player [Player], n int) {
	cont := 0
	for {

		if (_player[e].equipo == 1) {
			e++
			if (e == (n / 2)) {e = 0} //Validar que se mantenga con los IDs del equipo 1
		}

		else {
			e--
			if (e == (n / 2) - 1) {e = n -1} //Validar que se mantenga con los IDs del equipo 2
		}

		if (_player[e].meta == false || cont == n / 2) {
			e = -1
			break
		}

		cont++
	}

	return e
}

func match(jugador1 int, jugador2 int, duel_winner) {
	//0 = piedra, 1 = papel, 2 = tijera
	jugada1 := rand.Intn(3)
	jugada2 := rand.Intn(3)

	//Empate
	if (jugada1 == jugada2) {
		duel_winner = -1
		Close(duel_winner)
	}

	//No es empate
	else {
		if (jugada1 == 0 && jugada2 == 1) || (jugada1 == 1 && jugada2 == 2) || (jugada1 == 2 && jugada2 == 0) {
			duel_winner = jugador2
			Close(duel_winner)
		}

		else{
			duel_winner = jugadorjugador1
			Close(duel_winner)
		}
	}

}

go func Avanzar(_player []Player, i int, j int) {
	//Avanzar
	if (_player[i].equipo == 1) {
		_player[i].pos += 1
		_player[j].pos -= 1
		players_pos[0] = _player[i].pos
		players_pos[1] = _player[j].pos
	}

	else{
		_player[i].pos -=1
		_player[j].pos += 1
		players_pos[0] = _player[j].pos
		players_pos[1] = _player[i].pos
	}

	fmt.Println("El jugador %s del equipo %s avanza a la posición %s", i, _players[i].equipo, _player[i].pos)
	fmt.Println("El jugador %s del equipo %s avanza a la posición %s", j, _players[j].equipo, _player[j].pos)

	//Verificar si hay enfrentamiento o no:
	if (players_pos[0] == players_pos[1]){
		fmt.Println("Los jugadores %s y %s se enfrentan en un duelo", i, j)
		
		for {
			match(i, j, duel_winner)
			if (duel_winner > 0) {break}

			fmt.Println("Los jugadores %s y %s se enfrentan  de nuevo en un duelo porque empataron", i, j)
		}
	}
	
	Close(players_pos) //Cerrar el canal
}

func Controller(n_players int, players []Player) {
	duel_winner := make([]chan int) //Ganador del enfrentamiento
	players_pos := make([]chan int, 2) //Almacenan las posiciones de 1 jugador de cada equipo
	//players_pos[0] = jugador equipo 1, players_pos[1] = jugador equipo 2

	fmt.Println("Jugadores y sus equipos: ")
	
	//Crear los jugadores y asignarlos a su equipo correspondiente
	for i := 0; i < n; i++ {
		players[i].meta = false //setear que aún no llegan a la meta
		players[i].id := i //Asignar el id

		//Asignar equipos y posiciones
		if i < n/2 {
			players[i].equipo = 1
			players[i].pos_inicial = 1
			players[i].pos_objetivo = 20
			players[i].pos = 1
		}
		else {
			players[i].equipo = 2
			players[i].pos_inicial = 20
			players[i].pos_objetivo = 1
			players[i].pos = 20
		}
		fmt.Println("EL jugador %s se creó y está en el equipo %s", i, players[i].equipo)
	}

	players[0].continuar = true //1er jugador del equipo 1 avance
	players[n - 1].continuar = true //1er jugador del equipo 2 avance

	equipo1 := make([]bool, n/2) //Vector que contiene el orden del equipo 1
	equipo2 := make([]bool, n) //Vector que contiene el orden del equipo 2}

	equipo1[0] = true //Indicar el orden del equipo
	equipo2[n - 1] = true //Indicar el orden del equipo

	id := 0
	e1 = 0
	e2 = n - 1
	match_end := false

	fmt.Println("Empieza la partida")
	fmt.Println("El jugador %s del equipo %s comienza a avanzar", e1, players[e1].equipo)
	fmt.Println("El jugador %s del equipo %s comienza a avanzar", e2, players[e2].equipo)

	//Jugar
	for {
		Avanzar(players, equipo1, equipo2)
		id <- duel_winner

		fmt.Println("El jugador %s ganó del equipo %s, puede continuar", id, players[id].equipo)
		if (id >= n / 2) {
			fmt.Println("El jugador %s perdió del equipo %s, regresa al inicio", e1, players[e1].equipo)
			equipo1[e1] = false //El jugador actual ya no avanza
			players[e1].continuar = false //Decir que ya no puede seguir avanzando
			players[e1].pos = players[e1].pos_inicial //Hacer que regrese al inicio
			//e1++ //Obtener el siguiente jugador
			//if (e1 == (n / 2)) {e1 = 0} //Validar que se mantenga con los IDs del equipo 1
			e1 = NextPlayer(e1, players, n)
			players[e1].continuar = true //El siguiente jugador puede salir

			if (players[id].pos == players[id].pos_objetivo) {
				fmt.Println("El jugador %s del equipo %s llegó a la meta", id, players[id].equipo)
				players[id].meta = true //Indicar que el jugador llegó a la meta
				equipo2[id] = false //Indicar que para de avanzar
				e2 = NextPlayer(e2, players, n)

				if (e2 == -1) {
					fmt.Println("El equipo 2 ganó")
					match_end = true
				}
				
				/*for {
					e2--

					if (players[e2].meta == false && e2 >= n / 2) {break}

					if (e2 < n / 2) {
						fmt.Println("El equipo 2 ganó")
						match_end = true
						break
					}
				}*/

			}
		}

		else {
			fmt.Println("El jugador %s perdió del equipo %s, regresa al inicio", e2, players[e2].equipo)
			equipo2[e2] = false //El jugador actual ya no avanza
			players[e2].continuar = false //Decir que ya no puede seguir avanzando
			players[e2].pos = players[e2].pos_inicial //Hacer que regrese al inicio
			//e2-- //Obtener el siguiente jugador
			//if (e2 == (n / 2) - 1) {e2 = n -1} //Validar que se mantenga con los IDs del equipo 2
			e2 = NextPlayer(e2, players, n)
			players[e2].continuar = true //El siguiente jugador puede salir
		}

		if (match_end == true) {break} //terminó la partida
	}

}

func main() {
	var (
		n      int
		seedok bool
	)

	if len(os.Args) > 2 {
		if num, err := strconv.Atoi(os.Args[2]); err == nil {
			rand.Seed(int64(num))
			seedok = true
		}
	}
	if !seedok {
		rand.Seed(time.Now().UnixNano())
	}
	if len(os.Args) > 1 {
		if x, err := strconv.Atoi(os.Args[1]); err != nil {
			return
		} else {
			n = x
		}
	} else {
		n = 10
	}

	n_players := 10 //Cantidad de jugadores
	players := make([]Player, n)

	Controller(n_players, players)
}
