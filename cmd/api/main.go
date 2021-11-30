package main

import (
 	"flag"
	"fmt"
	 "log"
	"net/http"
)

const version = "1.0.0"

type config struct {
	 port int
	 env string

}

func main()  {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&cfg.env, "env","development", "Application environment (development | production)")
	flag.Parse()

	fmt.Println("Running")

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "status")
	})
	err := http.ListenAndServe(fmt.Sprintf(":%d",cfg.port),nil)
	if err != nil {
		log.Println(err)

	}
}