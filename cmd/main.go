package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	addr string
	

}

func main(){
	cfg,err:=loadConfig()
	if err!=nil{
		log.Fatal("Issues loading config")
	}
	
	mux:=http.NewServeMux()
	log.Printf("PricklyMango is being served on port %v",cfg.addr)
	http.ListenAndServe(cfg.addr,mux)



}

func loadConfig()(*config,error){
	cfg:=&config{addr: os.Getenv("PORT_NUMBER")}
	return cfg,nil



}
