# Programación concurrente y programación distribuida

Integrantes:

-Natalia Melissa Maury Castañeda

-Sergio Antonio Nuñez Lazo

-Oscar Daniel Flores Palermo

## HulaHoopRelayRace
Para la Tarea Académica (TA3), se realizó el juego de Hula Hoop Relay Race utilizando programación concurrente, canales y semáforos en el lenguaje de programación Go. El juego consiste en dos equipos que cruzarán un camino delimitado por hula hulas (Hula Hoop) y cuando dos jugadores se encuentran en la misma posición jugarán "piedra, papel o tijeras", el ganador del desafío continua y el perdedor regresar al inicio donde el siguiente compañero del equipo comienza a jugar. Esto se repetirá hasta que todos los jugadores de un equipo hayan llegado al otro extremo, es decir que crucen todo el camino de hula hulas. 

Para realizar este proyecto se hizo una simulación del juego en consola, donde el número de jugadores puede ser ingresado por el usuario, pero si no ingresa nada por
defecto el número de jugadores será 10. Una vez delimitada la cantidad de jugadores se dividen los jugadores en dos equipos del mismo tamaño, donde se utilizará 
funciones concurrentes, canales y semáforos para simular la partida. Los canales se usarán para comunicar a los jugadores entre ellos, los semáforos se usarán para 
evitar que el programa termine en starvation y/o deadlock, y las funciones concurrentes se utilizan para simular el avance de los dos equipos al mismo tiempo.

Las funciones principales para que la simulación sea exitosa son las siguientes: controller, cambiar de jugador, avanzar, duelo, ganar.
- La función controller se encarga de realizar la simulación del juego de inicio a fin: creando a los jugadores, asignarles un equipo, de llamar a las demás funciones
 como avanzar a los jugadores, verificar si sucede un duelo y cuando gana un equipo.
- La función cambiar de jugador se encarga de indicar al siguiente jugador del equipo que puede empezar a jugar cuando el jugador actual perdió o llego a la meta, o 
 llego a la meta.
- La función de avanzar le permite a los jugadores que están en los hula hulas acercarse de manera continua al otro extremo.
- La función de duelo se encarga de cuando los dos jugadores se encuentren en la misma tengan que simular jugar "piedra, papel o tijera" para definir quién es el 
 ganador que debe continuar y el perdedor que debe salirse.
- La función ganar se encarga de verificar que todos los jugadores de un equipo hayan llegado hasta la meta (al otro extremo) y poder declararse como el equipo vencedor.
