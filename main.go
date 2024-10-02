package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"text/template"
)

type Constructor struct {
	Name string `json:"team_name"`
	Base string `json:"base"`
}

type ConstructorDetails struct {
	Base               string `json:"base"`
	TeamPrincipal      string `json:"team_principal"`
	TechnicalChief     string `json:"technical_chief"`
	Chasis             string `json:"chasis"`
	EngineSupplier     string `json:"engine_supplier"`
	WorldChampionships string `json:"world_championships"`
	Wins               string `json:"wins"`
	PolePositions      string `json:"pole_positions"`
	FastestLaps        string `json:"fastest_laps"`
	Drivers            []struct {
		Name   string `json:"name"`
		Number string `json:"number"`
	} `json:"drivers"`
}

type Constructors []Constructor

type ConstructorDetailsWrapper struct {
	TeamName           string             `json:"team_name"`
	ConstructorDetails ConstructorDetails `json:"details"`
}

type ConstructorDetailsWrapperArray []ConstructorDetailsWrapper

func main() {

	fmt.Println("Starting server on port 8080")

	var constructors Constructors

	err := parseFile(&constructors)
	if err != nil {
		fmt.Printf("Error occured :- %v", err)
		return
	}

	var cwa ConstructorDetailsWrapperArray

	parseFile2(&cwa)

	http.HandleFunc("/", constructors.handler)
	http.HandleFunc("/{constructor_name}", cwa.contructorDetailsPageHandler)

	http.ListenAndServe(":8080", nil)

}

func (constructors Constructors) handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/index.html"))

	tmpl.Execute(w, constructors)

}

func (cwa ConstructorDetailsWrapperArray) contructorDetailsPageHandler(w http.ResponseWriter, r *http.Request) {

	path := path.Base(r.URL.Path)

	c := cwa.findConstructor(path)

	fmt.Println(c)

	fmt.Fprintln(w, c)

}

func (cwa ConstructorDetailsWrapperArray) findConstructor(constructor_name string) ConstructorDetailsWrapper {

	for _, c := range cwa {
		if c.TeamName == constructor_name {
			return c
		}
	}

	return ConstructorDetailsWrapper{}
}

func parseFile(contructors *Constructors) error {

	f, err := os.Open("f1_constructors_basic_info.json")
	if err != nil {
		return err
	}

	err = json.NewDecoder(f).Decode(contructors)
	if err != nil {
		return err
	}

	return nil
}

func parseFile2(cwa *ConstructorDetailsWrapperArray) error {

	f, err := os.Open("f1_constructors_details.json")
	if err != nil {
		return err
	}

	err = json.NewDecoder(f).Decode(cwa)
	if err != nil {
		return err
	}

	fmt.Println(cwa)

	return nil
}
