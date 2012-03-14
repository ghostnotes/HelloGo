package structsample

import (
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    var st struct {
        no int
        name string
    }

    st.no = 1
    st.name = "構造体1"

    fmt.Fprintf(w, "No = [%d]<br>", st.no)
    fmt.Fprintf(w, "Name = [%s]", st.name)
}