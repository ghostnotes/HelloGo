package operator

import (
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    add := 3 + 3
    dif := 7 - 3
//    pro := 2 * 5
//    quo := 3/2
//    rem := 3 % 2

    fmt.Fprintf(w, "add = %d\n", add)
    fmt.Fprintf(w, "dif = %d\n", dif)
}