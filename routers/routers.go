package routers

import (
	c "blueDigitalKwegonian/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router is main API router
func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/api/ck", c.ConvertKwengonianToDecimal).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
