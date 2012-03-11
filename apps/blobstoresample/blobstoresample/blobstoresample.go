package blobstoresample

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
    http.HandleFunc("/uploadURL", uploadURL)
}

func uploadURL(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        c.Debugf("GET request")        
    }else{
        c.Debugf("other request method.")
    }
}

var htmlTemplate = template.Must(template.New("html").Parse(htmlText))

const htmlText = `
<!DOCTYPE html>
<html><head><title>Blobstore API</title></head>
<body><form action="{{.}}" method="POST" enctype="multipart/form-data">
<label>更新するファイル: <input type="file" name="file"></label>
<input type="submit" value="送信">
</form></body></html>
`

func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
	c.Debugf("00")
    uploadURL, err := blobstore.UploadURL(c, "/upload", nil)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
		c.Errorf("01")
        return
    }
    err = htmlTemplate.Execute(w, uploadURL)
    if err != nil {
        http.Error(w, err.String(), http.StatusInternalServerError)
        c.Errorf("02")
    }
}

func upload(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
	c.Debugf("10")
    blobs, _, err := blobstore.ParseUpload(r)
    if err != nil {
        c.Errorf("parse upload failed.")
        http.Error(w, err.String(), http.StatusInternalServerError)
        return
    }else{
		c.Debugf("parse upload successful!")
    }

    file := blobs["file"]
    if len(file) == 0 {
        c.Errorf("no file uploaded")
        http.Redirect(w, r, "/", http.StatusFound)
        return
    }else{
		c.Debugf("find file uploaded.")
    }

    http.Redirect(w, r, "/store?blobKey=" + string(file[0].BlobKey), http.StatusFound)
}

func store(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    c.Debugf("20")

    blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}