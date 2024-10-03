package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Starting server on port 8080")

	var constructors Constructors

	var constructorsFullInfo ConstructorsFullInfo

	err := parseFiles(&constructors, &constructorsFullInfo)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", constructors.handler)
	http.HandleFunc("/{constructor_name}", constructorsFullInfo.contructorDetailsPageHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
