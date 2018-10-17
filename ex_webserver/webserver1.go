package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type mydata struct {
	Campo1 string
	Campo2 int
}

func main() {
	http.HandleFunc("/", myhandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	var data mydata
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Errorf("could not decode request: %v", err)
		http.Error(w, "could not decode request", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Ricevuto: %s e %d\n", data.Campo1, data.Campo2)
}
