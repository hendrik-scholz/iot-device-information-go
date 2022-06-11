package main

import (
	"net/http"

	"github.com/hendrik-scholz/iot-device-information-go/info"

	"github.com/gin-gonic/gin"
)

const port string = ":8080"

func main() {
	r := gin.Default()
	r.GET("/authorization", getAuthorization)
	r.GET("/geoposition", getGeoposition)
	r.GET("/identification", getIdentification)
	r.GET("/uuid", getUUID)
	r.Run(port)
}

func getAuthorization(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, info.GetAuthorization())
}

func getGeoposition(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, info.GetGeoPosition())
}

func getIdentification(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, info.GetIdentification())
}

func getUUID(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, info.GetUUID())
}