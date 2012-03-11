package usersample

import (
    "fmt"
    "http"
)


func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/required", required)
    http.HandleFunc("/admin", admin)
}

func handler(w http.ResponseWriter, r *http.Request) {
    var htmlText = `
    <!DOCTYPE html>
    <html>
    <head><title>メインページ</title></head>
    <body>
    <h1>メインページ</h1>
    <section>
    <h2>リンク</h2>
    <p><a href="/required">認証ページ</a></p>
    <p><a href="/admin">管理者ページ</a></p>
    </section>
    </body><html>
    `

    fmt.Fprintf(w, "%s", htmlText)
}

func required(w http.ResponseWriter, r *http.Request) {
    var htmlText = `
    <!DOCTYPE html>
    <html>
    <head><title>認証ページ</title></head>
    <body>
    <h1>認証ページ</h1>
    <section> 
    <h2>リンク</h2>
    <p><a href="/">戻る</a></p>
    <p><a href="/admin">管理者ページ</a></p>
    </section>
    </body></html>
    `

    fmt.Fprintf(w, "%s", htmlText)
}

func admin(w http.ResponseWriter, r *http.Request) {
    var htmlText = `
    <!DOCTYPE html>
    <html>
    <head><title>管理者ページ</title></head>
    <body>
    <h1>管理者ページ</h1>
    <section>
    <h2>リンク</h2>
    <p><a href="/">戻る</a></p>
    <p><a href="required">認証ページ</a></p>
    </section>
    </body></html>
    `
    fmt.Fprintf(w, "%s", htmlText)

}
 
    