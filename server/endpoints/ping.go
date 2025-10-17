package endpoints

import "net/http"

func EndpointGetPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
