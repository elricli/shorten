package shorten

import (
	"context"
	"errors"
	"log"
	"net/url"

	"github.com/drrrMikado/shorten/internal/bloomfilter"
	"github.com/drrrMikado/shorten/internal/fastrand"
	"github.com/drrrMikado/shorten/internal/validator"
	"github.com/go-redis/redis/v8"
)

var (
	redisHashTableKey = "shorten_hash_table_key"
)

// Insert url.
func Insert(ctx context.Context, rawurl string, redisClient *redis.Client, bf *bloomfilter.BloomFilter) (key string, err error) {
	if !validator.IsURL(rawurl) {
		return "", errors.New(rawurl + " is not valid url")
	}
	for {
		key = fastrand.String(10)
		if !bf.MightContain([]byte(key)) {
			break
		}
		log.Println("key might contain, key:", key)
	}
	if key == "" {
		return "", errors.New("key is empty")
	}
	if err = redisClient.HSet(ctx, redisHashTableKey, key, rawurl).Err(); err != nil {
		return
	}
	go bf.Insert([]byte(key))
	return
}

// Get url by key.
func Get(ctx context.Context, key string, redisClient *redis.Client, bf *bloomfilter.BloomFilter) (v string, err error) {
	if !bf.MightContain([]byte(key)) {
		return "", errors.New("key not exist")
	}
	v, err = redisClient.HGet(ctx, redisHashTableKey, key).Result()
	if err != nil {
		return
	}
	var u *url.URL
	u, err = url.Parse(v)
	if err != nil {
		return "", err
	} else if u.Scheme == "" {
		u.Scheme = "http"
		v = u.String()
	}
	return
}

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
	var ulrBox= document.getElementById("short_url");
	if (ulrBox) {
		ulrBox.focus();
		ulrBox.select();
	}
}
{{end}}
</script>
</html>
`
