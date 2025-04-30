package main

import (
	"github.com/AlexeyNilov/space_sim/html"
	"github.com/AlexeyNilov/space_sim/space"
	"flag"
	"log"
	"net/http"
)

func main() {
	// Parse flags
	var addr = flag.String("addr", "127.0.0.1:8080", "The address:port combination to listen on for HTTP requests.")
	flag.Parse()

	// Create a new HTTP server
	server := http.Server{
		Addr: *addr,
	}

	const spaceSize = 16 // Total number of points
	space := space.CreateSpace(spaceSize)
	side, grid := html.GenerateGrid(space)

	// Serve the static style.css file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	http.Handle("/", &html.TemplateHandler{Filename: "template.html", Side: side, Grid: grid})

	// Start the server
	log.Println("Starting web server on", *addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
