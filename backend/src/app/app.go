package app

import "net/http"

func StartApplication() {

	urlMappings()

	http.ListenAndServe(":8080", nil)
}
