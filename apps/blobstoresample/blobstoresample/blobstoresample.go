package blogstoresample

import (
    "appengine"
    "appengine/blobstore"
    "http"
    "template"
)

func init() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/upload", upload)
    http.HandleFunc("/store", store)
}

var htmlTemplate = template.Must(template.New("html").Parse(htmlText))

const htmlText = `
<!DOCTYPE html>
<html><head><title>Blobstore API</title></head>
<body><form action="{{.}}" method="POST" enctype="multipart/form-data">
<label>更新するファイル:<input tpe="file" name="file"></label>
<input type="submit" value="送信">
</form></body></html>
`

func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    uploadURL, err := blobstore.UploadURL(c, "/upload", nil)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }

    err = htmlTemplate.Execute(w, uploadURL)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
    }
}

func upload(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    blobs, _, err := blobstore.ParseUpload(r)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError) 
        return 
   }

    file := blobs["file"]
    if len(file) == 0 {
        c.Errorf("no file uploaded")
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }

    http.Redirect(w, r, "/store?blobKey=" + string(file[0].BlobKey), http.StatusFound)
}

func store(w http.ResponseWriter, r *http.Request) {
    blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}