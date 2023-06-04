Programacion concurrente y programacion distribuida
Intengrantes:
-Natalia Melissa Maury Castañeda
-Sergio Antonio Nuñez Lazo
-Oscar Daniel Flores Palermo

HulaHoopRelayRace
En el presente trabajo de la TA3, se ha realizado el juego de Hula Hoop Relay Race con el uso de programacion concurrente, canales y semaforos en el lenguaje de
programacion Go. El juego consiste en dos equipos que cruzaran un camino delimitado por hula hulas(Hula Hop) y cuando se encuentren dos jugadores en la misma posicion
jugaran "piedra, papel o tijeras", el ganador de ese desafio continua y el perdedor regresar al inicio donde el siguiente compañero del equipo comienza a jugar. Esto
se repetira hasta que todos los jugadores de un equipo hayan llegado al otro extremo, es decir que crucen todo el camino de hula hulas. 

Para realizar este proyecto se hizo una simulacion del juego en consola, donde el numero de jugadores puede ser ingresado por el usuario pero si no ingresa nada por
defecto el numero de jugadores sera 10. Una vez delimitada la cantidad de jugadores se dividen los jugdaores en dos equipos del mismo tamaño, donde se utilizara 
funciones concurrentes, canales y semaforos para simular la partida. Los canales se usaran para comunicar a los jugadores entre ellos, los semaforos se usaran para 
evitar que el programa termine en starvation y/o deadlock, y las funciones concurrentes se utilizan para simular el avance de los dos equipos al mismo tiempo.

Las funciones principales para que la simulacion sea exitosa son las siguientes: controller, cambiar de jugador, avanzar, duelo, ganar.
-La funcion controller se encarga de realizar la simulacion del juego de inicio a fin: creando a los jugadores, asignarles un equipo, de llamar a las demas funciones
 como avanzar a los jugadores, verificar si sucede un duelo y cuando gana un equipo.
-La funcion cambiar de jugador se encarga de indicar al siguiente jugador del equipo que puede empezar a jugar cuando el jugador actual pierdio o llego a la meta, o 
 llego a la meta.
-La funcion de avanzar le permite a los jugadores que estan en los hula hulas acercarse de manera continua al otro extremo.
-La funcion de duelo se encarga de cuando los dos jugadores se encuentren en la misma tengan que simular jugar "piedra, papel o tijera" para definir quien es el 
 ganador que debe continuar y el perdedor que debe salirse.
-La funcion ganar se encarga de verificar que todos los jugadores de un equipo hayan llegado hasta la meta (al otro extremo) y poder declararse como el equipo vencedor.
