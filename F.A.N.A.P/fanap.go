package main

import (
	"fmt"
	"log"
	"net/http"

	"F.A.N.A.P/router"
)

func main() {
	fmt.Println("Rectangle Overlap")

	r := router.Router()
	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")

}
