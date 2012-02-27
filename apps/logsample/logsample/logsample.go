package logsample

import (
    "appengine"
    "fmt"
    "http"
)


func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    c.Debugf("Debug log")

    c.Infof("Info log")
    c.Warningf("Warning log")
    c.Errorf("Error log")
    c.Criticalf("Critical log")

    fmt.Fprintf(w, "Log test")
}