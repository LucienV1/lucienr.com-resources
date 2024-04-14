package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var root *string = flag.String("r", "/path/to/root", "The root directory to serve files from. Default is /path/to/root")

func randomcss(w http.ResponseWriter, r *http.Request) {
	uq := r.URL.Query().Get("n")
	var n int
	if uq == "" {
		n = rand.Intn(30)
	} else {
		var err error
		n, err = strconv.Atoi(uq)
		if err != nil {
			http.Error(w, "Invalid parameter.", 400)
			return
		}
	}
	cssFile := fmt.Sprintf("%s/styles.css.%d", *root, n)
	file, err := os.Open(cssFile)
	if err != nil {
		http.Error(w, "File not found.", 404)
		return
	}
	defer file.Close()

	http.ServeContent(w, r, "styles.css", time.Unix(0, 0), file)
}

func main() {
	port := flag.String("p", "3332", "The port to listen on. Default is 3332")
	flag.Parse()
	http.HandleFunc("/styles.css", randomcss)
	http.Handle("/", http.FileServer(http.Dir(*root)))
	http.ListenAndServe("127.0.0.1:"+*port, nil)
}
