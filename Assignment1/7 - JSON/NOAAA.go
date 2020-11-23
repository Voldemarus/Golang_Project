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
	fmt.Println("\n* Source data, taken from JSON sources")
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
	fmt.Println("Magnetic field: Bt ", magField.Bt, " Bz: ", magField.Bz)

	flux := Flux{}
	err = getJson(myClient, kNOAAURLflux10cm, &flux)
	if err != nil {
		log.Fatalln("Error during writing to output file", err)
		return
	}
	fmt.Println("Flux 10 cm: ", flux.Flux, "\n")

	// To set up timestamp properly we should take all three timestmps,
	// and select the latest one.
	windTime := decodeTimeString(solarWind.TimeStamp)
	magTime := decodeTimeString(magField.TimeStamp)
	fluxTime := decodeTimeString(flux.TimeStamp)

	lastTime := windTime
	if magTime.After(lastTime) {
		lastTime = magTime
	}
	if fluxTime.After(lastTime) {
		lastTime = fluxTime
	}
	// Generate timestamp string in RFC3339 format
	outputTimeString := lastTime.Format(time.RFC3339)
	fmt.Println("* TimeStamp - ", outputTimeString)

	//
	// Now we should create output JSON, which will agregete all information, gathered so far
	//
	outputMap := map[string]string{"SolarWind": solarWind.WindSpeed, "Bt": magField.Bt,
		"Bz": magField.Bz, "Flux": flux.Flux, "TimeStamp": outputTimeString}

	fmt.Println("\n* Created data structure - ", outputMap, "\n")

	outJson, err := json.Marshal(outputMap)
	if err != nil {
		log.Fatalln("Cannot encode JSON - ", err)
	}

	// Print out generate JSON byte stream as a string representation
	fmt.Println("* Final JSON - ", string(outJson), "\n")

}

func decodeTimeString(timeStr string) time.Time {
	// Example 2020-11-23 07:58:00
	timeValue, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		log.Fatalln("Cannot convert time string - ", err)
	}
	return timeValue

}

func getJson(myC http.Client, url string, target interface{}) error {
	r, err := myC.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
