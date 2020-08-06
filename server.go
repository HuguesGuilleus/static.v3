// static.v3
// Copyright (c) 2020, HuguesGuilleus
// BSD 3-Clause License

// Create easily a http Handler for static file.
package static

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	Log = log.New(os.Stderr, "[STATIC ERROR]", log.LstdFlags)
	// Dev disable the minifing and read the file at for each request.
	Dev bool = false
)

type Server struct {
	Body   []byte
	Mime   string
	Minify Minifier
	// used only if Dev is true
	get func() []byte
}

func New(mime string, min Minifier) *Server {
	return &Server{
		Mime:   mime,
		Minify: min,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", s.Mime)
	if Dev && s.get != nil {
		s.Body = s.get()
	}
	w.Write(s.Body)
}

func (s *Server) Bytes(body []byte) *Server {
	s.Body = s.Minify.min(body)
	s.get = nil
	return s
}

func (s *Server) Func(f func() []byte) *Server {
	s.Body = s.Minify.min(f())
	s.get = f
	return s
}

// Server a static content with a min Content-Type header.
//
// The content is by default d. If f is non empty, the function read recurrent
// from f serve it. The reading error are silent.
//
// The served content are minify (expect if Dev is enable) with min. If min
// is nil, the content are not minify.
func (s *Server) File(path string) *Server {
	s.get = func() []byte { return readFileOnce(path, s.Minify) }
	go func() {
		b := s.get()
		if len(b) > 0 {
			s.Body = b
		}
	}()
	return s
}

// Like Server.File() but with a splietd path
func (s *Server) FileJoinPath(path ...string) *Server {
	return s.File(filepath.Join(path...))
}

// readFileOnce read the content of the file or directory, minify each, concat
// all part and return it.
func readFileOnce(f string, m Minifier) []byte {
	data := make([]byte, 0)
	filepath.Walk(f, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			Log.Printf("get info error %q: %v", p, err)
			return nil
		} else if info.IsDir() || strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		d, err := ioutil.ReadFile(p)
		if err != nil {
			Log.Printf("read error %q: %v", p, err)
			return nil
		}
		data = append(data, m.min(d)...)
		return nil
	})
	return data
}

type Minifier func([]byte) []byte

// Minify even m if nil
func (m Minifier) min(in []byte) []byte {
	if Dev || m == nil || len(in) == 0 {
		return in
	}
	return m(in)
}
