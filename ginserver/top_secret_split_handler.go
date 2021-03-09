package ginserver

import (
	"example.com/juampiq6/example_api/cookieManager"
	"example.com/juampiq6/example_api/models"
	"github.com/gin-gonic/gin"
)

func postTopSecretSplitHandler(c *gin.Context) {
	cookieuuid := cookieManager.GetCookieFromContext(c)
	satelliteName := c.Param("name")

	var satellitesInfo = models.SateliteInfo{
		Name: satelliteName,
	}
	// bindea el json de la request con el modelo, si falla devuelve 400
	err := c.BindJSON(&satellitesInfo)
	if err == nil {
		cookieManager.AddSatelliteInfo(&cookieuuid, &satellitesInfo)
		c.JSON(200, "Se guardó la información del satelite!")
	}
}

func getTopSecretSplitHandler(c *gin.Context) {
	cookieuuid := cookieManager.GetCookieFromContext(c)
	if cookieManager.CheckIfHasThreeValues(&cookieuuid) {
		listSatellite := cookieManager.GetSatelliteInfoList(&cookieuuid)
		response := proccessDeciphers(&listSatellite)
		c.JSON(200, response)

	} else {
		c.JSON(409, "No existe suficiente información para calcular la fuente y el contenido del mensaje")
	}

}
