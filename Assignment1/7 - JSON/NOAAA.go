package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	//	kNOAAURLscales string = "https://services.swpc.noaa.gov/products/noaa-scales.json"
	kNOAAURLsolarWind string = "https://services.swpc.noaa.gov/products/summary/solar-wind-speed.json"
	kNOAAURLmagField  string = "https://services.swpc.noaa.gov/products/summary/solar-wind-mag-field.json"
	kNOAAURLflux10cm  string = "https://services.swpc.noaa.gov/products/summary/10cm-flux.json"
)

type SolarWind struct {
	WindSpeed string
	TimeStamp string
}

type MagField struct {
	Bt        string
	Bz        string
	TimeStamp string
}

type Flux struct {
	Flux      string
	TimeStamp string
}

func main() {
	myClient := http.Client{Timeout: 10 * time.Second}
	// Read data from the sources
	solarWind := SolarWind{}
	err := getJson(myClient, kNOAAURLsolarWind, &solarWind)
	if err != nil {
		log.Fatalln("Error during writing to output file", err)
		return
	}
	fmt.Println("Solar wind : ", solarWind.WindSpeed)

	magField := MagField{}
	err = getJson(myClient, kNOAAURLmagField, &magField)
	if err != nil {
		log.Fatalln("Error during writing to output file", err)
		return
	}
	fmt.Println("Magfield: Bt ", magField.Bt, " Bz: ", magField.Bz)

	flux := Flux{}
	err = getJson(myClient, kNOAAURLflux10cm, &flux)
	if err != nil {
		log.Fatalln("Error during writing to output file", err)
		return
	}
	fmt.Println("Flux 10 cm: ", flux.Flux)

}

func getJson(myC http.Client, url string, target interface{}) error {
	r, err := myC.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
