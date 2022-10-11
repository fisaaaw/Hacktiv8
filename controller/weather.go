package controller

import (
	api "ass3/rest_api"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type weather struct {
	Status detail `json:"status"`
}
type detail struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func GetWeather(ctx *gin.Context) {
	resp, err := api.HitGetWeather()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	var payload weather
	fmt.Println(string(resp))
	err = json.Unmarshal(resp, &payload)
	fmt.Println(payload)

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"result": payload,
	})

}

func SetWeather(ctx *gin.Context) {
	var req weather
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	request, _ := json.Marshal(req)

	resp, err := api.HitSetWeather(request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	var payload weather
	fmt.Println(string(resp))
	err = json.Unmarshal(resp, &payload)
	fmt.Println(payload)

	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"result": payload,
	})

}
