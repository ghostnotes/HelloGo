package writesample

import (
    "appengine"
    "appengine/datastore"
    "fmt"
    "http"
    "time"
)

type Guest struct {
   Name string
    Date datastore.Time
}

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/write", write)
}

const inputForm = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>名前の登録</title>
</head>
<body>
<form method="POST" action="write">
<label>お名前<input type="text" name="name" /></label>
<input type="submit">
</form>
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s", inputForm)
}

func write(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusNotFound)
        w.Header().Set("Context-Type", "text/plain; charset=utf-8")
        fmt.Fprintf(w, "Not Found")
        return
    }

    c := appengine.NewContext(r)

    // DataStoreへの書き込み
    var g Guest
    g.Name = r.FormValue("name")
    g.Date = datastore.SecondsToTime(time.Seconds())

    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest", nil), &g); err != nil {
        http.Error(w, "Internal Server Error : " + err.String(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/", http.StatusFound)
}