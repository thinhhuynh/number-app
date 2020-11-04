package web

import (
	"encoding/json"
	"log"
	"net/http"
	"number-app/util"
)

func (a *App) FindHighestPrimeNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	number := r.URL.Query().Get("number")
	log.Println(".....API find-highest-prime is staring...", number)
	ok, bigNumber := util.IsANumber(number)
	
	if (!ok) {
		sendErr(w, http.StatusInternalServerError, "INVALID_INPUT_NUMBER")
		return
	}

	highestPrime := util.FindHighestPrime(bigNumber)

	err := json.NewEncoder(w).Encode(highestPrime)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
	log.Println(".....API find-highest-prime ended...")
}


