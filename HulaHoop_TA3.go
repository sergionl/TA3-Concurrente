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
}

func match(Player _player1, Player _player2) {
	
}

func pipesort(in <-chan int, out, sorted chan<- int) {
	min := <-in
	for next := range in {
		if next < min {
			out <- min
			min = next
		} else {
			out <- next
		}
	}
	sorted <- min
	close(out) //Cerrar el canal
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

	ch := make([]chan int, n+1)
	duel_winner := make([]chan bool) //Ganador del enfrentamiento
	players_pos := make([]chan int, 2) //Almacenan las posiciones de 1 jugador de cada equipo
	//players_pos[0] = jugador equipo 1, players_pos[1] = jugador equipo 2
	sorted := make(chan int) 
	a := make([]int, n)
	ch[0] = make(chan int)
	
	fmt.Println("Jugadores y sus equipos: ")
	
	//Crear los jugadores y asignarlos a su equipo correspondiente
	for i := 0; i < n; i++ {
		players[i].continuar = true //setear continuar a true
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

	go func Jugar(Player _player) {
		//Avanzar
		if (_player.equipo == 1) {
			_player.pos += 1
			players_pos[0]
		}

		else{
			_player.pos -=1
			players_pos[1]
		}

		//Verificar si hay enfrentamiento o no:
		if (players_pos[0] == players_pos[1]){

		}

		Close(players_pos) //Cerrar el canal
	}

	//Proceso para generar números aleatorios
	go func() {
		for i := 0; i < n; i++ {
			num := rand.Intn(100)
			fmt.Println(num)
			ch[0] <- num
		}
		close(ch[0]) //Cerrar el canal
	}()
	for i := 0; i < n; i++ {
		a[i] = <-sorted //Recibir el arreglo ordenad/sorted
	}
	fmt.Println(a) //Mostrar/Imprimir el arreglo ordenado
}
