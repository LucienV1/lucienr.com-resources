package main

import (
        "fmt"
        "math/rand"
        "net/http"
        "os"
        "time"
)

var root string = "/var/www/lucienr.com/"

func randomcss(w http.ResponseWriter, r *http.Request) {
        n := rand.Intn(30)
        cssFile := fmt.Sprintf("%s/styles.css.%d", root, n)
        file, err := os.Open(cssFile)
        if err != nil {
                http.Error(w, "File not found.", 404)
                return
        }
        defer file.Close()

        http.ServeContent(w, r, "styles.css", time.Unix(0, 0), file)
}

func main() {
        http.HandleFunc("/styles.css", randomcss)
        http.Handle("/", http.FileServer(http.Dir(root)))
        http.ListenAndServe("127.0.0.1:3332", nil)
}
