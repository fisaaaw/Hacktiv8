package controller

import (
	"encoding/json"
	"io/ioutil"
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

	resp, err := ioutil.ReadFile("./data.json")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	var payload weather
	err = json.Unmarshal(resp, &payload)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK,
		payload,
	)

}
func SetWeather(ctx *gin.Context) {
	var requestBody weather

	err := ctx.BindJSON(&requestBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	file, _ := json.MarshalIndent(requestBody, "", " ")
	_ = ioutil.WriteFile("data.json", file, 0644)

	ctx.AbortWithStatusJSON(http.StatusOK,
		requestBody,
	)

}
