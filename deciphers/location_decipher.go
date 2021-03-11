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
	/// las distancias siempre llegan en el mismo orden definido en el mapa "satellitePositions" -> kenobi - skywalker - sato
	xKen := float32(satellitePositions["kenobi"][0])
	yKen := float32(satellitePositions["kenobi"][1])

	xSky := float32(satellitePositions["skywalker"][0])
	ySky := float32(satellitePositions["skywalker"][1])

	xSato := float32(satellitePositions["sato"][0])
	ySato := float32(satellitePositions["sato"][1])

	/// Este es un problema de triliteracion, utilizado por los gps para obtener nuestra ubicacion
	///
	/// (−2xken+2xsky)x + (−2yken+2ysky)y = distken^2 − distsky^2 − xken^2 + xsky^2 − yken^2 + ysky^2.
	/// (−2xsky+2xsato)x + (−2ysky+2ysato)y = distsky^2 − distsato^2 − xsky^2 + xsato^2 − ysky^2 + ysato^2.

	/// Ax + By =C
	/// Dx + Ey = F

	/// x = (CE − FB) / (EA−BD)
	/// y = (CD −AF) / (BD−AE)

	a := -2*xKen + 2*xSky
	b := -2*yKen + 2*ySky
	d := -2*xSky + 2*xSato
	e := -2*ySky + 2*ySato
	c := power2(distances[0]) - power2(distances[1]) - power2(xKen) + power2(xSky) - power2(yKen) + power2(ySky)
	f := power2(distances[1]) - power2(distances[2]) - power2(xSky) + power2(xSato) - power2(ySky) + power2(ySato)

	x = ((c * e) - (f * b)) / ((e * a) - (b * d))
	y = ((c * d) - (a * f)) / ((b * d) - (a * e))
	return x, y
}

func power2(num float32) float32 {
	/// la funcion pow de math requiere castear a float 64, esto no lo necesitamos
	return num * num
}
