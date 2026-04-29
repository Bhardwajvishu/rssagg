package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
)

func main(){
	fmt.Println("Hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("Port is not found")
	}

	router:= chi.NewRouter()

	fmt.Println("Port:", portString)
}