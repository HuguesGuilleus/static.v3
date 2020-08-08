// static.v2
// Copyright (c) 2020, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"errors"
	"io"
	"path/filepath"
	"text/template"
)

var templateNoInit = errors.New("The template is not init.")

// Can be breaken
type Template struct {
	T   *template.Template
	get func() []byte
}

func NewTemplate() *Template {
	return &Template{}
}

func (t *Template) Execute(w io.Writer, data interface{}) error {
	if t == nil || t.T == nil {
		return templateNoInit
	}
	if Dev && t.get != nil {
		t.parse(t.get())
	}
	return t.T.Execute(w, data)
}

func (t *Template) Bytes(in []byte) *Template {
	t.get = nil
	if Dev {
		t.parse(in)
	} else {
		t.parse(HtmlMinify(in))
	}
	return t
}

func (t *Template) Func(f func() []byte) *Template {
	t.get = f
	t.parse(HtmlMinify(f()))
	return t
}

func (t *Template) File(f string) *Template {
	t.get = func() []byte { return readFileOnce(f, HtmlMinify) }
	t.parse(t.get())
	return t
}
func (t *Template) FileJoinPath(path ...string) *Template {
	return t.File(filepath.Join(path...))
}

func (t *Template) parse(in []byte) {
	var err error
	t.T, err = template.New("").Parse(string(in))
	if err != nil {
		t.T = nil
		Log.Printf("template parse error: %v", err)
	}
}
