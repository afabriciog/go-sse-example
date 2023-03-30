package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		// Set the response headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Create the channel to send data to the client
		messageChan := make(chan string)

		// Start a goroutine to send data to the client
		go func() {
			for {
				// Send the current time to the client every second
				messageChan <- fmt.Sprintf("data: %v\n\n", time.Now().Format("2006-01-02 15:04:05"))
				time.Sleep(time.Second)
			}
		}()

		// Send the data to the client as it comes in
		for {
			select {
			case message := <-messageChan:
				_, err := fmt.Fprint(w, message)
				if err != nil {
					fmt.Printf("Server error: %v\n", err)
					return
				}
				w.(http.Flusher).Flush()
			case <-r.Context().Done():
				fmt.Println("Client disconnected")
				return
			}
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		return
	}
}
