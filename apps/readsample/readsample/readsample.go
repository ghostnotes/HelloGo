package readsample

import (
    "appengine"
    "appengine/datastore"
    "fmt"
    "http"
    "template"
    "time"
)

func init() {
    http.HandleFunc("/", handler)
}

type Guest struct {
    Name string
    Date datastore.Time
}

type Guest_View struct {
    Name string
    Date string
}

const guestTemplateHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>登録者リスト</title>
</head>
<body>
<table border="1">
<tr><th>名前</th><th>登録日時</th></tr>
{{range .}}
    <tr>
    {{if .Name}}
      <td>{{.Name|html}}</td>
    {{else}}
      <td>名無し</td>
    {{end}}
    {{if .Date}}
      <td>{{.Date|html}}</td>
    {{else}}
      <td> - </td>
    {{end}}
    </tr>
{{end}}
</table>
</body>
</html>
`

var guestTemplate = template.Must(template.New("guest").Parse(guestTemplateHTML))

func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    q := datastore.NewQuery("Guest").Order("Date")
    count, err := q.Count(c)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }

    guests := make([]Guest, 0, count)
    if _, err := q.GetAll(c, &guests); err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }
    guest_views := make([]Guest_View, count)
    for pos, guest := range guests {
        guest_views[pos].Name = fmt.Sprintf("%s", guest.Name)
        localTime := time.SecondsToLocalTime(int64(guest.Date) / 100000)
        guest_views[pos].Date = fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", localTime.Year, localTime.Month, localTime.Day, localTime.Hour, localTime.Minute, localTime.Second)
    }

    if err := guestTemplate.Execute(w, guest_views); err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
    }
}