package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGetChirpsParallel(b *testing.B) {
	// Your mux with real or stubbed handlers:
	port := "8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/api/chirps", func(w http.ResponseWriter, r *http.Request) {
		// Simulate some real processing:
		w.Write([]byte(`["chirp1", "chirp2", "chirp3"]`))
	})

	// Spin up a test server (in-memory)
	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	// Reset the timer so Go measures only the actual request loop
	b.ResetTimer()

	// Use b.RunParallel to spawn multiple goroutines
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// Each iteration does one GET request
			resp, err := http.Get(fmt.Sprintf("http://localhost:%s/api/chirps", port))
			if err != nil {
				// b.Error("Could not send GET request: %v", err)
				fmt.Println("Could not send GET request:", err)
			}
			resp.Body.Close()
		}
	})
}
