package jsonsample

import (
    "http"
    "fmt"
    "json"
)

type Person struct {
	Name string
    Intro string
}

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    var person = Person{ Name:"aaa", Intro:"bbb"}
//    fmt.Fprintf(w, person.Name + person.Intro)

    var data []byte
    data, _ = json.Marshal(person)

    fmt.Fprintf(w, string(data))    
}