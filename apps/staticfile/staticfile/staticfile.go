package staticfile

import (
    "http"
    "template"
)

func init() {
    http.HandleFunc("/staticfile", handler)
}

const templateHTML = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>レスポンス</title>
</head>

<body>
    入力した名前は{{.|html}}です。
</body>
</html>
`

func handler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    htmltemplate := template.Must(template.New("html").Parse(templateHTML))
    err := htmltemplate.Execute(w, name)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
    }
}