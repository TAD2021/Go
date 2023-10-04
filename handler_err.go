package main

import (
	"net/http"
)

func handlerErr(w http.ResponseWriter, r *http.Request){
	responseWithJSON(w, 400, "Sommething went wrong")
}