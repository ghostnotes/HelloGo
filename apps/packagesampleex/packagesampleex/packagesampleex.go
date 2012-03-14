package helloworld

import (
    "http"
)

func init() {
    http.HandleFunc("/", handler)
}