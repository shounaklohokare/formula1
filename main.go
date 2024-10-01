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

type Constructors []Constructor

func main() {

	fmt.Println("Starting server on port 8080")

	var constructors Constructors

	err := parseFile(&constructors)
	if err != nil {
		fmt.Printf("Error occured :- %v", err)
		return
	}

	http.HandleFunc("/", constructors.handler)
	http.HandleFunc("/{constructor_name}", constructors.contructorDetailsPageHandler)

	http.ListenAndServe(":8080", nil)

}

func (constructors Constructors) handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/index.html"))

	tmpl.Execute(w, constructors)

}

func (constructors Constructors) contructorDetailsPageHandler(w http.ResponseWriter, r *http.Request) {

	path := path.Base(r.URL.Path)

	fmt.Fprintln(w, path)

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
