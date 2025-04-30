package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/AlexeyNilov/space_sim/html"
	"github.com/AlexeyNilov/space_sim/space"
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

	// Create a new hub
	hub := html.NewHub()
	go hub.Run()

	// Serve the static style.css file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Handle WebSocket connections
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		html.ServeWs(hub, w, r)
	})

	// Handle the main page
	http.Handle("/", &html.TemplateHandler{Filename: "template.html", Side: side, Grid: grid})

	// Start a goroutine to update the grid periodically
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		pointID := 0

		for range ticker.C {
			// Update the space (you can add your space update logic here)
			(*space)[pointID].Content = "P"
			if pointID == spaceSize-1 {
				pointID = 0
			} else {
				pointID++
			}
			side, grid := html.GenerateGrid(space)

			// Create the update message
			update := struct {
				Side int    `json:"side"`
				Grid string `json:"grid"`
			}{
				Side: side,
				Grid: grid,
			}

			// Convert to JSON
			message, err := json.Marshal(update)
			if err != nil {
				log.Printf("Error marshaling update: %v", err)
				continue
			}

			// Broadcast the update to all connected clients
			hub.Broadcast(message)
		}
	}()

	// Start the server
	log.Println("Starting web server on", *addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
