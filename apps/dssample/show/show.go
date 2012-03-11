package show

import (
    "appengine"
    "appengine/datastore"
    "fmt"
    "http"
    "template"
    "time"
)

func init() {
    http.HandleFunc("/", show)
    http.HandleFunc("/input", input)
}

type Guest struct {
    Name string
    Date datastore.Time
}

type Guest_View struct {
    Id int64
    Name string
    Date string
}

const showHTML = `
<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>データストアの更新・削除</title>
<script>
function insert() {
   document.frm.action="/input";
   submit();
}

function update(txt, id) {
    document.frm.action="/update";
    document.frm.id.value = id;
    document.frm.updname.value = txt.value;
    submit();
}

function del(txt, id) {
    document.frm.action="/delete";
    document.frm.id.value = id;
    document.frm.updname.value = txt.value;
    submit();
}

function submit() {
    document.frm.method = "POST";
    document.frm.submit();
}
</script>
</head>

<body>
<form name="frm">
<section>
<h1>新規登録</h1>
<label>名前: </label><input type="text" name="name" /><input type="button" value="登録" onClick="insert()" />
</section>
<section>
<h1>登録済みデータ</h1>
<table border="1">
<table border="1">
<tr><th>更新</th><th>削除</th><th>名前</th><th>更新時間</th>
{{range .}}
    <tr>
        <td><input type="button" value="更新" onClick="update(updname{{.Id|html}}, {{.Id|html}}" /></td>
        <td><input type="button" value="削除" onClick="del(updname{{.Id|html}}, {{.Id|html}})" /><td>
        <td>{{.Date|html}}</td>
    </tr>
{{end}}
</table>
</section>
<input type="hidden" name="id" />
<input type="hidden" name="updname" />
</form>
</body>
</html>
`

var showHTMLTemplate = template.Must(template.New("show").Parse(showHTML))

func show(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    q := datastore.NewQuery("Guest").Order("Date")
    count, err := q.Count(c)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }

    guests := make([]Guest, 0, count)
    guest_views :=([]Guest_View, count)
    if keys, err := q.GetAll(c, &guests); err != nil {
        http.Error(w, errString(), http.StatusInternalServerError)
        return
    } else {
        for pos, guest := range guests {
            guest_views[pos].Id = keys[pos].IntID()
            guest_views[pos].Name = guest.Name
            localTime := time.SecondsToLocalTime(int64(guest.Date) / 1000000)
            guest_views[pos].Date = fmt.Sprintf("%04d/%02d/%02d %02d:%02d:%02d", localTime.Year, localTime.Month, localTime.Day, localTime.Hour, localTime.Minute, localTime.Second)
        }
    }

    if err := showHTMLTemplate.Execute(w, guest_views); err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
    }
}

func input(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    var g Guest
    g.Name = r.FormValue("name")
    g.Date = datastore.SecondsToTime(time.Seconds())
    if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Guest", nil), &g); err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/", http.StatusFound)
}