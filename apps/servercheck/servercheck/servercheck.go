package servercheck

import (
    "appengine"
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    isDev := appengine.IsDevAppServer()

    if isDev {
        //trueの時は開発サーバ
        fmt.Fprintf(w, "開発サーバで実行中\n")
    } else {
        fmt.Fprintf(w, "本番環境で実行中\n")
    }
}