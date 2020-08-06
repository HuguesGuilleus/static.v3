// static.v2
// Copyright (c) 2020, HuguesGuilleus
// BSD 3-Clause License

package static

import (
	"io"
	"path/filepath"
	"text/template"
)

// Can be breaken
type Template struct {
	T   *template.Template
	get func() []byte
}

func NewTemplate() *Template {
	return &Template{}
}

func (t *Template) Execute(w io.Writer, data interface{}) error {
	if Dev && t.get != nil {
		t.parse(t.get())
	}
	return t.T.Execute(w, data)
}

func (t *Template) Bytes(in []byte) *Template {
	t.get = nil
	t.parse(in)
	return t
}

func (t *Template) Func(f func() []byte) *Template {
	t.get = f
	t.parse(f())
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
	t.T, err = template.New("").Parse(string(t.get()))
	if err != nil {
		Log.Printf("template parse error: %v", err)
	}
}
