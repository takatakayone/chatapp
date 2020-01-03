package app

import (
	"fmt"
	"net/http"
)

func urlMappings() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "simple server")
	})

	http.HandleFunc("/ws", serveWs)
}
