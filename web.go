// static.v2
// Copyright (c) 2020, Hugues GUILLEUS. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package static

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/xml"
	"log"
	"regexp"
)

// A binding of File for Css file.
//
//	http.HandleFunc("/style.css", static.Css(nil, "front/style/"))
func Css() *Server {
	return &Server{
		Mime:   "text/css; charset=utf-8",
		Minify: CssMinify,
	}
}
func CssMinify(in []byte) []byte {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	out, err := m.Bytes("text/css", in)
	if err != nil {
		log.Println("[STATIC MINIFY] CSS error:", err)
	}
	return out
}

// A binding of File for Html file.
//
//	http.HandleFunc("/", static.Html(nil, "front/index.html"))
func Html() *Server {
	return &Server{
		Mime:   "text/html",
		Minify: HtmlMinify,
	}
}
func HtmlMinify(in []byte) []byte {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepDocumentTags: true,
		KeepEndTags:      true,
	})
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("image/svg+xml", xml.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	out, err := m.Bytes("text/html", in)
	if err != nil {
		log.Println("[STATIC MINIFY] HTML error:", err)
	}
	return out
}

// A binding of File for Js file.
//
//	http.HandleFunc("/app.js", static.Js(nil, "front/app.js"))
func Js() *Server {
	return &Server{
		Mime:   "application/javascript",
		Minify: JsMinify,
	}
}
func JsMinify(in []byte) []byte {
	m := minify.New()
	m.AddFunc("application/javascript", js.Minify)
	out, err := m.Bytes("application/javascript", in)
	if err != nil {
		log.Println("[STATIC MINIFY] JS error:", err)
	}
	return out
}

// A binding of File for SVG image.
//
//	http.HandleFunc("/icon.svg", static.Svg(nil, "front/icon.svg"))
func SVG() *Server {
	return &Server{
		Mime:   "image/svg+xml",
		Minify: SvgMinify,
	}
}
func SvgMinify(in []byte) []byte {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("image/svg+xml", xml.Minify)
	out, err := m.Bytes("image/svg+xml", in)
	if err != nil {
		log.Println("[STATIC MINIFY] SVG error:", err)
	}
	return out
}

// A binding of File for JPEG image.
func Jpeg() *Server {
	return &Server{Mime: "image/jpeg"}
}

// A binding of File for PNG image.
func Png() *Server {
	return &Server{Mime: "image/jpeg"}
}

// A binding of File for WebP image.
func WebP() *Server {
	return &Server{Mime: "image/jpeg"}
}
