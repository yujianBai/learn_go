package main

import (
	"fmt"
	"net/http"
)

const SERVER_PORT = 10088
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "hello"

func rootHadler(w http.ResponseWriter, req * http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}

func main(){
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHadler)
	http.ListenAndServeTLS(fmt.Sprint(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
}