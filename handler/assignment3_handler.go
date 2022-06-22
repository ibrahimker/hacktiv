package handler

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/ibrahimker/latihan-register/entity"
)

var webData entity.WebData

const htmlPath = "static/web.html"
const jsonPath = "static/weather.json"

func Assignment3Handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")
	// read from json file and write to webData
	file, _ := ioutil.ReadFile(jsonPath)
	json.Unmarshal(file, &webData)
	templates, _ := template.ParseFiles(htmlPath)
	context := entity.WebData{
		Status: entity.Status{
			Water:          webData.Status.Water,
			Wind:           webData.Status.Wind,
			StatusCompiled: webData.Status.StatusCompiled,
		},
	}
	templates.Execute(w, context)
}

func GenerateToJson() {
	for {
		// fill web data
		webData.Status.StatusCompiled = "aman"
		webData.Status.Water = rand.Intn(100)
		webData.Status.Wind = rand.Intn(100)
		if (webData.Status.Water >= 6 && webData.Status.Water <= 8) || (webData.Status.Wind >= 7 && webData.Status.Wind <= 15) {
			webData.Status.StatusCompiled = "siaga"
		} else if webData.Status.Water > 8 || webData.Status.Wind > 15 {
			webData.Status.StatusCompiled = "bahaya"
		}

		// write to json file
		jsonString, _ := json.Marshal(&webData)
		ioutil.WriteFile(jsonPath, jsonString, os.ModePerm)

		// sleep for 15 seconds
		time.Sleep(15 * time.Second)
	}
}
