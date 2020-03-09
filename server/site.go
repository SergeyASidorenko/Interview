package main

import "net/http"

func main() {
	http.HandleFunc("/", IndexHandle)
	http.ListenAndServe(":80", nil)
}
func IndexHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HELLO"))
}
