package appengine

import (
    "appengine"
    "fmt"
    "http"
)


func init() {
    http.HandleFunc("/", handler)
}


func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintf(w, "<!DOCTYPE html>")
    fmt.Fprintf(w, "<html><head><title>appengineパッケージ</title></head><body>")

    if appengine.IsDevAppServer() {
        // 開発環境
        fmt.Fprintf(w, "開発環境で動作しています。。<br>")
    } else {
        fmt.Fprintf(w, "本番環境で動作しています。<br>")
    }

    c := appengine.NewContext(r)

    fmt.Fprintf(w, "AppID(): %s<br>", appengine.AppID(c))
    fmt.Fprintf(w, "DefaultVersionHostname(): %s<br>", appengine.DefaultVersionHostname(c))

    fmt.Fprintf(w, "VersionID(): %s<br>", appengine.VersionID(c))
    fmt.Fprintf(w, "</body></html>")
}