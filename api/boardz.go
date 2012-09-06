package main 

import (
	"flag"
	"fmt"
	"github.com/bmizerany/pat"
	"log"
	"net/http"
)




func main() {
	var addr string
	var port int
	var mongoip string 

	flag.StringVar(&addr, "addr", "127.0.0.1", "Address")
	flag.IntVar(&port, "port", 8080, "Port") 
	flag.StringVar(&mongoip, "mongoip", "127.0.0.1", "MongoDB Ip")
	flag.Parse()
	
	m := pat.New()

	// GET
	m.Get("/boards/", http.HandlerFunc(ListBoards))
	m.Get("/boards/:id", http.HandlerFunc(ShowBoard))
	m.Get("/postits/", http.HandlerFunc(ListPostits))
	m.Get("/postits/:id", http.HandlerFunc(ShowPostit))

	// POST
	m.Post("/boards/", http.HandlerFunc(CreateBoard))
	m.Post("/postits/", http.HandlerFunc(CreatePostit))
	
	http.Handle("/", m)

	// Start the HTTP server
	log.Printf("Listening on http://%s/\n", addr)
	foo := fmt.Sprintf("%s:%d", addr, port)	
	err := http.ListenAndServe(foo, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}	
}

