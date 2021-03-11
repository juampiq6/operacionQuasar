package deciphers

import "strings"

func GetMessage(messages ...[]string) (msg string) {
	// inicializamos el map
	messagesMap := make(map[string]int)
	// iteramos por los argumentos recibidos
	for _, messageReceived := range messages {
		// iteramos por las palabras que recibio cada satelite
		for i, word := range messageReceived {
			// si la palabra es un string vacio, la omitimos
			if word != "" {
				// si la palabra no existe en el map, o si existe y su posicion en el array es mas grande que la anterior ocurrencia
				if index, exists := messagesMap[word]; !exists || (exists && index > i) {
					// lo agregamos al map, asignandole esa posicion
					messagesMap[word] = i
				}
			}

		}
	}
	return getMessageFromMap(&messagesMap)
}

func getMessageFromMap(messageMap *map[string]int) string {
	// determinamos la cantidad de palabras
	mapLen := len(*messageMap)
	// creamos el slice con esa cantidad de palabras
	messageSlice := make([]string, mapLen)
	// ordenamos los elementos del mapa dentro del slice
	for k, i := range *messageMap {
		if i >= mapLen {
			messageSlice = append(messageSlice, k)
		} else {
			messageSlice[i] = k
		}
	}
	// retornamos las palabras del mensaje unidas por un espacio
	return strings.Join(messageSlice, " ")
}
