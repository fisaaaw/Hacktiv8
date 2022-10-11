package restapi

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HitGetWeather() (out []byte, err error) {
	resp, err := http.Get("http://localhost:9090/weather")
	if err != nil {
		return
	}
	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

func HitSetWeather(req []byte) (out []byte, err error) {
	resp, err := http.Post("http://localhost:9090/weather", "application/json", bytes.NewBuffer(req))
	if err != nil {
		return
	}
	out, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
