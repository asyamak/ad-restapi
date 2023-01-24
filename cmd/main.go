package main

import (
	"ad-api/config"
	"ad-api/internal/app"

	"log"
)

func main(){
	cnf,err := config.New("./config/config.json")
	if err != nil {
		log.Printf("main: error receive config: %v",err)
		return
	}
	
	if err = app.New(cnf).Start(); err != nil{
		log.Printf("main: error initialize application: %v",err)
		return
	}

	
}