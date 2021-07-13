package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serverHandler)
	log.Fatalln(http.ListenAndServe(":5000", nil))
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<!DOCTYPE html>
	<html lang="en">
	<head>
		<h1> This is a simple server </h1>
	</head>
	<body>
	</body>
	</html>`)
}
