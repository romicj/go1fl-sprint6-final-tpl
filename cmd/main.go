package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	mylog := log.New(os.Stderr, "my:", log.LstdFlags)
	serv := server.CreateServer(mylog)

	mylog.Fatal(serv.ListenAndServe())
}
