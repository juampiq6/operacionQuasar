package ginserver

import (
	"example.com/juampiq6/example_api/deciphers"
	"example.com/juampiq6/example_api/models"
	"github.com/gin-gonic/gin"
)

func postTopSecretHandler(c *gin.Context) {
	var satellitesInfoArray models.TopSecretRequest
	// bindea el json de la request con el modelo, si falla devuelve 400
	err := c.BindJSON(&satellitesInfoArray)
	if err == nil {
		if len(satellitesInfoArray.Satellites) < 3 {
			c.JSON(404, "Son necesarios por lo menos 3 satelites para determinar la ubicación de la nave")
		} else {

			response := proccessDeciphers(&satellitesInfoArray.Satellites)
			c.JSON(200, response)
		}
	}

}

func proccessDeciphers(satellites *[]models.SateliteInfo) models.TopSecretResponse {
	var messagesList [][]string
	var distancesList = make([]float32, 3)
	for _, v := range *satellites {
		messagesList = append(messagesList, v.Message)
		distancesList[deciphers.GetSatelliteOrder(&v.Name)] = v.Distance
	}
	obtainedMessage := deciphers.GetMessage(messagesList...)
	obtainedX, obtainedY := deciphers.GetLocation(distancesList...)
	response := models.TopSecretResponse{
		Message: obtainedMessage,
		Position: models.Position{
			X: obtainedX,
			Y: obtainedY,
		},
	}
	return response
}
