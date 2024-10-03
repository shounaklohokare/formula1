package main

import (
	"net/http"
	"path"
	"text/template"
)

func (constructors Constructors) handler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, constructors)

}

func (constructorsFullInfo ConstructorsFullInfo) contructorDetailsPageHandler(w http.ResponseWriter, r *http.Request) {

	path := path.Base(r.URL.Path)

	c := constructorsFullInfo.findConstructor(path)

	tmpl := template.Must(template.ParseFiles("templates/constructor_details.html"))

	tmpl.Execute(w, c)

}

func (constructorsFullInfo ConstructorsFullInfo) findConstructor(constructor_name string) ConstructorFullInfo {

	for _, c := range constructorsFullInfo {
		if c.TeamName == constructor_name {
			return c
		}
	}

	return ConstructorFullInfo{}
}
