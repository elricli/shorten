package static

import "embed"

//go:embed css/shorten.min.css
//go:embed js/shorten.min.js
//go:embed img/*
//go:embed html/*.html
var FS embed.FS
