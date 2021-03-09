package deciphers

var satellitePositions = map[string][3]int{
	// en la tercera posicion guardamos el orden que deberia tener, ya q un map no tiene nativamente orden
	"kenobi":    {-500, -200, 0},
	"skywalker": {100, -100, 1},
	"sato":      {500, 100, 2},
}

func GetSatelliteOrder(name *string) int {
	var order int
	for key, value := range satellitePositions {
		if *name == key {
			order = value[2]
			break
		}
	}
	return order
}

func GetLocation(distances ...float32) (x, y float32) {
	/// la formula de Euclides nos ayuda a calcular la distancia entre dos puntos dadas sus coordenadas
	/// d(a,b)^2 = (Xa - Xb)^2 + (Ya - Yb)^2
	/// Si despejamos primero X y luego Y de la formula, podremos reemplazar el valor de X en la formula donde despajamos Y,
	/// y de esta manera poder depender de una sola variable
	/// las formulas finales han sido simplificadas para

	return
}
