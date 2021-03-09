package cookieManager

import (
	"log"

	"example.com/juampiq6/example_api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var cookiesMap = map[string][]models.SateliteInfo{}

func GetSatelliteInfoList(cookieuuid *string) []models.SateliteInfo {
	return cookiesMap[*cookieuuid]
}

func AddSatelliteInfo(cookieuuid *string, info *models.SateliteInfo) {
	infoForCookie, exists := cookiesMap[*cookieuuid]
	if !exists {
		cookiesMap[*cookieuuid] = []models.SateliteInfo{
			*info,
		}
	} else {
		if CheckIfHasThreeValues(cookieuuid) {
			for _, v := range infoForCookie {
				if info.Name == v.Name {
					infoForCookie = append(infoForCookie, *info)
				}
			}
		} else {
			cookiesMap[*cookieuuid] = append(cookiesMap[*cookieuuid], *info)
		}
	}
}

func CheckIfHasThreeValues(cookieuuid *string) bool {
	return len(cookiesMap[*cookieuuid]) == 3
}

func generateCookieUuid() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatalln("Error generating UUID", err)
	}
	return uuid.String()
}

var cookieKeyName = "uuid"

// Obtiene el valor si existe la cookie sino larga error
func GetCookieFromContext(c *gin.Context) string {
	value, err := c.Cookie(cookieKeyName)
	if err != nil {
		value = generateCookieUuid()
		c.SetCookie(cookieKeyName, value, 86400, "/topsecret_split", "", false, true)
	}
	return value
}

// TODO: funcion para purgar registros q nunca mas se van a usar?
