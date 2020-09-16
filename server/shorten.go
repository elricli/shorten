package server

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

const indexTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta http-equiv="Content-type" content="text/html;charset=UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta name="keywords" content="URL shortener,shorten,URL,link,smaller,shorter,hyperlink,shorten web address">
<title>URL Shorten</title>
</head>
<body style="margin: 0;font-family: Roboto, Arial, sans-serif;font-size: x-large;">
<div id="container" style="text-align: center;width: 60%;margin: 0 auto;">
<div style="margin-top: 50px;">
<p>Please specify a URL to shorten.</p>
{{if .ErrMsg}}<p style="color: red;">{{.ErrMsg}}</p>{{end}}
<form action="/shorten" method="post">
<input type="url" name="url" maxlength="500" size="30" required placeholder="https://example.com" style="font-size: x-large;"><br>
<input type="submit" value="shorten!" style="margin: 20px;font-size: x-large;">
</form>
{{if .URL}}
Your shortened URL is:<br/>
<input type="text" id="short_url" value="" onclick="select_text();" readonly="readonly" style="font-size: large;">
<div style="font-size: small">Your shortened URL goes to: {{.OriURL}}</div>
{{end}}
</div>
</div>
</body>
<script>
{{if .URL}}
let domain = window.location.origin + "/"+{{.URL}};
document.getElementById("short_url").setAttribute("value", domain);
function select_text() {
	var urlbox = document.getElementById("short_url");
	if (urlbox) {
		urlbox.focus();
		urlbox.select();
	}
}
{{end}}
</script>
</html>
`

// T is template struct.
type T struct {
	URL    string
	ErrMsg string
	OriURL string
}

// Shorten http handler.
func Shorten(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	url := r.Form.Get("url")
	log.Println(r.Form)
	key, err := svc.Shorten(url)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errcode": 1,
			"errmsg":  err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"errcode": 0,
		"errmsg":  "",
		"data":    key,
	})
	return
}

// Shorten2 http handler..
func Shorten2(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(string(debug.Stack()))
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	r.ParseForm()
	url := r.Form.Get("url")
	t := T{}
	key, err := svc.Shorten(url)
	if err != nil {
		t.ErrMsg = err.Error()
	}
	t.URL = key
	t.OriURL = url
	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	tmpl.Execute(w, t)
}

// Redirect http handler.
func Redirect(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	url, err := svc.Redirect(key)
	if err != nil || url == "" {
		w.Write([]byte(`Sorry, the link you accessed doesn't exist on our service. Please keep in mind that our shortened links are case sensitive and may contain letters and numbers.`))
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return
}

// Index handler.
func Index(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}()
	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	tmpl.Execute(w, T{})
}
