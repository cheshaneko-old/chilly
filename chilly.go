package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"os"
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"github.com/cheshaneko/chilly/gpiofarm"
	"github.com/cheshaneko/chilly/mockfarm"
	"encoding/json"
	//"time"
)

type Room struct {
	Temperature uint
	Humidity uint
}

type IFarm interface {
	Open() error
	Close() error
	OnLight() error
	OffLight() error
	OnMotor() error
	OffMotor() error
	WaterGreen() error
	WaterViolet() error
	WaterOrange() error
	WaterBlue() error
	TemperatureAndHumidityRoom() (uint, uint, error)
}

var (
	myfarm IFarm
)

func onTask() {
	myfarm.OnLight()
}

func offTask() {
	myfarm.OffLight()
}

func requestOn(w http.ResponseWriter, r *http.Request) {
	myfarm.OnLight()
}

func requestOff(w http.ResponseWriter, r *http.Request) {
	myfarm.OffLight()
}

func requestOnMotor(w http.ResponseWriter, r *http.Request) {
	myfarm.OnMotor()
}

func requestOffMotor(w http.ResponseWriter, r *http.Request) {
	myfarm.OffMotor()
}

func requestWaterGreen(w http.ResponseWriter, r *http.Request) {
	myfarm.WaterGreen()
}

func requestWaterViolet(w http.ResponseWriter, r *http.Request) {
	myfarm.WaterViolet()
}

func requestWaterOrange(w http.ResponseWriter, r *http.Request) {
	myfarm.WaterOrange()
}

func requestWaterBlue(w http.ResponseWriter, r *http.Request) {
	myfarm.WaterBlue()
}

func requestRoom(w http.ResponseWriter, r *http.Request) {
	t, h, err := myfarm.TemperatureAndHumidityRoom()
	if err != nil {
		fmt.Println(err)
        http.Error(w, "Not found", http.StatusNotFound)
        return
	}
	room := Room{t, h}
	jData, _ := json.Marshal(room)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
    log.Println(r.URL)
    if r.URL.Path != "/" {
        http.Error(w, "Not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }
    http.ServeFile(w, r, "frontend/home.html")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", serveHome)
	myRouter.HandleFunc("/OnLight", requestOn).Methods("POST")
	myRouter.HandleFunc("/OffLight", requestOff).Methods("POST")
	myRouter.HandleFunc("/OnMotor", requestOnMotor).Methods("POST")
	myRouter.HandleFunc("/OffMotor", requestOffMotor).Methods("POST")
	myRouter.HandleFunc("/WaterGreen", requestWaterGreen).Methods("POST")
	myRouter.HandleFunc("/WaterViolet", requestWaterViolet).Methods("POST")
	myRouter.HandleFunc("/WaterOrange", requestWaterOrange).Methods("POST")
	myRouter.HandleFunc("/WaterBlue", requestWaterBlue).Methods("POST")
	myRouter.HandleFunc("/room", requestRoom).Methods("get")
    myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	if true {
		myfarm = new(gpiofarm.GpioFarm)
	} else {
		myfarm = new(mockfarm.MockFarm)
	}

	// Open and map memory to access gpio, check for errors
	if err := myfarm.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer myfarm.Close()

	gocron.Every(1).Day().At("07:00").Do(onTask)
	gocron.Every(1).Day().At("19:00").Do(offTask)

	gocron.Start()

	handleRequests()
}
