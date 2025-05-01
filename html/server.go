package html

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/AlexeyNilov/space_sim/space"
)

func StartServer(addr *string, side int, grid string, space *space.Space) {
	// Create a new HTTP server
	server := http.Server{
		Addr: *addr,
	}

	// Create a new hub
	hub := NewHub()
	go hub.Run()

	// Serve the static style.css file
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Handle WebSocket connections
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, w, r)
	})

	// Handle the main page
	http.Handle("/", &TemplateHandler{Filename: "template.html", Side: side, Grid: grid})

	// Start a goroutine to update the grid periodically
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			// Update the space (you can add your space update logic here)

			side, grid := GenerateGrid(space)

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
