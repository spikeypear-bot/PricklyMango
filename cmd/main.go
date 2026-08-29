package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spikeypear-bot/PricklyMango/internal"
)

type config struct {
	addr string
	configPath string	

}

func main(){
	cfg,err:=loadConfig()
	if err!=nil{
		log.Fatal("Issues loading config")
	}

	testHandler:=internal.NewTestHandle()
	
	mux:=http.NewServeMux()
	mux.Handle("/test",testHandler)
	log.Printf("PricklyMango is being served on port %v",cfg.addr)
	http.ListenAndServe(cfg.addr,mux)



}

func loadConfig()(*config,error){
	cfg:=&config{addr: os.Getenv("PORT_NUMBER"),configPath: os.Getenv("CONFIG_PATH")}
	return cfg,nil



}

