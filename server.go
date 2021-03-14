package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var minLat  int = -90
var maxLat  int = 90
var minLong  int = -180
var maxLong  int = 180

type Response struct {
	Lat    float64     `json:"lat"`
	Long  float64 `json:"long"`
}

func Start(w http.ResponseWriter, r *http.Request){
	rand.Seed(time.Now().Unix())

	rLat :=  float64(minLat + rand.Intn(maxLat - minLat)) + rand.Float64()
	rLong  := float64(minLong + rand.Intn(maxLong - minLong)) + rand.Float64()
	fmt.Println(rLat, rLong)
	//w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Content-Type", "application/json")
	//response := Response {
	//	Lat: rLat,
	//	Long: rLong,
	//}
	//
	//json.NewEncoder(w).Encode(response)
	//http.ServeFile(w,r,"http://127.0.0.1:4000/start")
	s := fmt.Sprintf("%f", rLat)
	s2 := fmt.Sprintf("%f", rLong)
	fmt.Fprintf(w, s + " " + s2)

}


func setup(){
	http.HandleFunc("/start", Start)
}

func main(){
	setup()
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)
}


