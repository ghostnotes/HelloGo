package usersample2

import (
    "appengine"
    "appengine/user"
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/loginfederated", loginfederated)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
//    panic('test')
}


func handler(w http.ResponseWriter, r *http.Request) {
    c :=  appengine.NewContext(r)
    if !printUserInfo(c, w, r) {
        // ログインURLにリダイレクト
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, url, http.StatusFound)
    }
}

func loginfederated(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    if !printUserInfo(c, w, r) {
        // ログインURL(OpenID)にリダイレクト
        url, err := user.LoginURLFederated(c, r.URL.String(), "me.yahoo.co.jp")
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, url, http.StatusFound)
    }
}

func printUserInfo(c appengine.Context, w http.ResponseWriter, r *http.Request)(result bool) {
    result = false
    u := user.Current(c)
    if u != nil {
        // ログイン済み
        var htmlHeader string = `
            <!DOCTYPE html>
            <html>
            <head><title>ログイン済</title></head>
        `
        fmt.Fprintf(w, "%s", htmlHeader)
        fmt.Fprintf(w, "<body>")

        if user.IsAdmin(c) {
            fmt.Fprintf(w, "<p>今、管理者ユーザとしてログインしています。</p>")
        }else{
            fmt.Fprintf(w, "<p>今、認証済みです。</p>")
        }

        fmt.Fprintf(w, "<p>Email: %s</p>", u.Email)
        fmt.Fprintf(w, "<p>AuthDomain: %s</p>", u.AuthDomain)
        fmt.Fprintf(w, "<p>Id: %s</p>", u.Id)
        fmt.Fprintf(w, "<p>Identity: %s</p>", u.FederatedIdentity)
        fmt.Fprintf(w, "<p>Provider: %s</p>", u.FederatedProvider)

        url, err := user.LogoutURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.String(), http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "<a href=\"%s\">ログアウト</a>", url)
        fmt.Fprintf(w, "</body></html>")
        return true    
     }

     return
}