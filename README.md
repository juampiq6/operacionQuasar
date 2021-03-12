# Operación Quasar
Deployed in **HEROKU**

**https://operacion-quasar-jpq.herokuapp.com/**

## Servicios

POST  **/topsecret**
Example request body:

	{
		"satellites": [
			{ 
				"name": "kenobi"
				"distance": 100
				"message": ["este", "", "", "mensaje", ""]
			},
			{
				"name": "skywalker",
				"distance": 115.5,
				"message": ["este","", "un", "", ""]
			},
			{
				"name": "sato",
				"distance": 142.7,
				"message": ["","es","", "", "secreto"]
			}]
	}
	
>	If the request body is incorrect, a 400 response code will be sent.

Example response body:

	{
		"position": {
			"x": -487.2859,
			"y": 1557.0142
		},
		"message": "este es un mensaje secreto"
	}
	
	200 OK

----

GET  **/topsecret_split/**
>	This method uses cookies to retrieve the values already sent. If 3 values exist, the response can be successful. Otherwise, a 409 response code will be sent, as there is not enough information to calculate the result.

Example response body:
 
	
	{
		"position": {
			"x": -487.2859,
			"y": 1557.0142
		},
		"message": "este es un mensaje secreto"
	}
	
	200 OK


---

POST  **/topsecret_split/{satellite_name}**
>	This method uses cookies to store the values sent

Example request -> /topsecret_split/{satellite_name}
Body:

	{
		"distance": 400,
		"message": ["este", "", "", "mensaje", ""]
	}
	
>	If the request body is incorrect, a 400 response code will be sent.

Example response body:
 
	"Se guardó la información del satelite!"
	
	200 OK
